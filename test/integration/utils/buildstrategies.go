// Copyright The Shipwright Contributors
//
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"context"

	"github.com/shipwright-io/build/pkg/apis/build/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// This class is intended to host all CRUD calls for testing BuildStrategy CRDs resources

// CreateBuildStrategy generates a BuildStrategy on the current test namespace
func (t *TestBuild) CreateBuildStrategy(ctx context.Context, bs *v1alpha1.BuildStrategy) error {
	bsInterface := t.BuildClientSet.BuildV1alpha1().BuildStrategies(t.Namespace)

	_, err := bsInterface.Create(ctx, bs, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}
