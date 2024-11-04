// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package e2etest

import (
	"context"
	"testing"

	"github.com/hashicorp/go-version"

	"github.com/martian-cloud/terraform-exec/tfexec"
)

func TestInit(t *testing.T) {
	runTest(t, "basic", func(t *testing.T, tfv *version.Version, tf *tfexec.Terraform) {
		err := tf.Init(context.Background())
		if err != nil {
			t.Fatalf("error running Init in test directory: %s", err)
		}
	})
}
