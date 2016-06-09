// @flow

import autotest from "sourcegraph/util/autotest";

import React from "react";

import TreeSearch from "sourcegraph/tree/TreeSearch";
import TreeStore from "sourcegraph/tree/TreeStore";
import * as TreeActions from "sourcegraph/tree/TreeActions";

import testdataFiles from "sourcegraph/tree/testdata/TreeSearch-files.json";
import testdataNotFound from "sourcegraph/tree/testdata/TreeSearch-notFound.json";

describe("TreeSearch", () => {
	it("should render files", () => {
		TreeStore.directDispatch(new TreeActions.FileListFetched("repo", "c", {Files: ["p1/p2/f3", "p1/f2"]}));
		autotest(testdataFiles, `${__dirname}/testdata/TreeSearch-files.json`,
			<TreeSearch repo="repo" rev="rev" commitID="c" path="p1/p2" prefetch={true} overlay={true} initialDefs={true} location={{query: {q: ""}}} />,
			{router: {}, status: {}, user: {}},
		);
	});

	it("should display 404 for not found directory", () => {
		TreeStore.directDispatch(new TreeActions.FileListFetched("repo", "c", {Files: ["p1/p2/f3", "p1/f2"]}));
		autotest(testdataNotFound, `${__dirname}/testdata/TreeSearch-notFound.json`,
			<TreeSearch repo="repo" rev="rev" commitID="c" path="p1/notfound" prefetch={true} overlay={true} initialDefs={true} location={{query: {q: ""}}} />,
			{router: {}, status: {}, user: {}},
		);
	});
});
