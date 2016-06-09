package localstore

import (
	"fmt"

	"github.com/keegancsmith/que-go"
	"golang.org/x/net/context"
	"gopkg.in/gorp.v1"
	"gopkg.in/inconshreveable/log15.v2"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/dbutil"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/store"
	"sourcegraph.com/sourcegraph/sourcegraph/services/backend/accesscontrol"
)

func init() {
	// From schema.sql in github.com/keegancsmith/que-go
	AppSchema.CreateSQL = append(AppSchema.CreateSQL, `
CREATE TABLE que_jobs
(
  priority    smallint    NOT NULL DEFAULT 100,
  run_at      timestamptz NOT NULL DEFAULT now(),
  job_id      bigserial   NOT NULL,
  job_class   text        NOT NULL,
  args        json        NOT NULL DEFAULT '[]'::json,
  error_count integer     NOT NULL DEFAULT 0,
  last_error  text,
  queue       text        NOT NULL DEFAULT '',

  CONSTRAINT que_jobs_pkey PRIMARY KEY (queue, priority, run_at, job_id)
);`)
	AppSchema.DropSQL = append(AppSchema.DropSQL, "DROP TABLE IF EXISTS que_jobs;")
}

const queueMaxAttempts = 2

type queue struct{}

var _ store.Queue = (*queue)(nil)

func (q *queue) Enqueue(ctx context.Context, j *store.Job) error {
	if err := accesscontrol.VerifyUserHasReadAccess(ctx, "Queue.Enqueue."+j.Type, nil); err != nil {
		return err
	}
	c, err := q.client(ctx)
	if err != nil {
		return err
	}
	return c.Enqueue(q.toQue(j))
}

func (q *queue) LockJob(ctx context.Context) (*store.LockedJob, error) {
	if err := accesscontrol.VerifyUserHasReadAccess(ctx, "Queue.LockJob", nil); err != nil {
		return nil, err
	}
	c, err := q.client(ctx)
	if err != nil {
		return nil, err
	}
	j, err := c.LockJob("")
	if err != nil || j == nil {
		return nil, err
	}

	// We don't want to retry jobs forever, we would rather log and drain
	// the queue.
	errFunc := j.Error
	if j.ErrorCount+1 >= queueMaxAttempts {
		errFunc = func(reason string) error {
			log15.Debug("store.Job.MarkError ignoring error", "type", j.Type, "lastErr", j.LastError.String, "err", err)
			return j.Delete()
		}
	}

	return store.NewLockedJob(q.fromQue(j), j.Delete, errFunc), nil
}

func (q *queue) toQue(j *store.Job) *que.Job {
	return &que.Job{
		Type: j.Type,
		Args: j.Args,
	}
}

func (q *queue) fromQue(j *que.Job) *store.Job {
	return &store.Job{
		Type: j.Type,
		Args: j.Args,
	}
}

func (q *queue) client(ctx context.Context) (*que.Client, error) {
	dbh := dbutil.GetUnderlyingSQLExecutor(appDBH(ctx))
	dbm, ok := dbh.(*gorp.DbMap)
	if !ok {
		return nil, fmt.Errorf("queue could not get underlying *sql.Db from appDBH")
	}
	return que.NewClient(dbm.Db), nil
}
