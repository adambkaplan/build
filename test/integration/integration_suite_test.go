// Copyright The Shipwright Contributors
//
// SPDX-License-Identifier: Apache-2.0

package integration_test

import (
	"context"
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/shipwright-io/build/test/integration/utils"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

// TODO: clean resources in cluster, e.g. mainly cluster-scope ones
// TODO: clean each resource created per spec
var (
	deleteNSList []string
	tb           *utils.TestBuild
	err          error
	ctx          context.Context
)

var _ = BeforeSuite(func() {
	ctx = context.Background()
})

var _ = BeforeEach(func() {
	tb, err = utils.NewTestBuild()
	if err != nil {
		fmt.Printf("fail to get an instance of TestBuild, error is: %v", err)
	}

	err := tb.CreateNamespace(ctx)
	if err != nil {
		fmt.Printf("fail to create namespace: %v, with error: %v", tb.Namespace, err)
	}

	deleteNSList = append(deleteNSList, tb.Namespace)

	// We store a channel for each Build operator instance we start,
	// so that we can nuke the instance later inside the AfterEach Ginkgo
	// block
	tb.StopBuildOperator, err = tb.StartBuildOperator()
	if err != nil {
		fmt.Println("fail to start the Build powerful operator", err)
	}
})

var _ = AfterEach(func() {
	// Close the channel, meaning we nuke an instance of the Build
	// operator
	if tb.StopBuildOperator != nil {
		close(tb.StopBuildOperator)
	}
})

var _ = AfterSuite(func() {
	// Ensure a proper cleanup of test environments
	Expect(tb.DeleteNamespaces(ctx, deleteNSList)).To(BeNil())
})
