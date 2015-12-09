package statsutil

import (
	"os/exec"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"golang.org/x/net/context"
	"gopkg.in/inconshreveable/log15.v2"

	"src.sourcegraph.com/sourcegraph/fed"
	"src.sourcegraph.com/sourcegraph/go-sourcegraph/sourcegraph"
)

var numReposGauge = prometheus.NewGauge(prometheus.GaugeOpts{
	Namespace: "src",
	Subsystem: "usage_stats",
	Name:      "repos_total",
	Help:      "Total repos on the local Sourcegraph instance.",
})

var buildLabels = []string{"build_type"}
var numBuildsGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "src",
	Subsystem: "usage_stats",
	Name:      "builds_total",
	Help:      "Total builds on the local Sourcegraph instance.",
}, buildLabels)

var numUsersGauge = prometheus.NewGauge(prometheus.GaugeOpts{
	Namespace: "src",
	Subsystem: "usage_stats",
	Name:      "users_total",
	Help:      "Total users on the local Sourcegraph instance.",
})

var execDurationGauge = prometheus.NewGauge(prometheus.GaugeOpts{
	Namespace: "src",
	Subsystem: "process",
	Name:      "exec_duration_seconds",
	Help:      "Duration of executing the echo command.",
})

func init() {
	prometheus.MustRegister(numReposGauge)
	prometheus.MustRegister(numBuildsGauge)
	prometheus.MustRegister(numUsersGauge)
	prometheus.MustRegister(execDurationGauge)

	// test the latency of exec, which may increase under certain memory conditions
	go func() {
		for {
			time.Sleep(10 * time.Second)

			s := time.Now()
			if err := exec.Command("echo").Run(); err != nil {
				log15.Warn("exec measurement failed", "error", err)
				continue
			}
			execDurationGauge.Set(time.Now().Sub(s).Seconds())
		}
	}()
}

// ComputeUsageStats takes a daily snapshot of the basic statistics of all
// local repos.
func ComputeUsageStats(ctx context.Context, interval time.Duration) {
	cl := sourcegraph.NewClientFromContext(ctx)
	if cl == nil {
		log15.Warn("ComputeUsageStats: could not construct client, usage stats will not be computed")
		return
	}
	for {
		updateNumRepos(cl, ctx)
		updateNumBuilds(cl, ctx)
		updateNumUsers(cl, ctx)

		time.Sleep(interval)
	}
}

func updateNumRepos(cl *sourcegraph.Client, ctx context.Context) {
	reposList, err := cl.Repos.List(ctx, &sourcegraph.RepoListOptions{
		ListOptions: sourcegraph.ListOptions{PerPage: 10000},
	})
	if err != nil {
		log15.Warn("ComputeUsageStats: could not compute number of repos", "error", err)
		return
	}
	numReposGauge.Set(float64(len(reposList.Repos)))

	if fed.Config.IsRoot {
		// don't compute committer stats on the mothership.
		return
	}
}

func updateNumBuilds(cl *sourcegraph.Client, ctx context.Context) {
	numBuilds, err := ComputeBuildStats(cl, ctx)
	if err != nil {
		log15.Warn("ComputeUsageStats: could not compute number of builds", "error", err)
		return
	}
	for buildType, buildCount := range numBuilds {
		numBuildsGauge.WithLabelValues(buildType).Set(float64(buildCount))
	}
}

func updateNumUsers(cl *sourcegraph.Client, ctx context.Context) {
	usersList, err := cl.Users.List(ctx, &sourcegraph.UsersListOptions{
		ListOptions: sourcegraph.ListOptions{PerPage: 10000},
	})
	if err != nil {
		log15.Warn("ComputeUsageStats: could not compute number of users", "error", err)
		return
	}
	numUsersGauge.Set(float64(len(usersList.Users)))
}
