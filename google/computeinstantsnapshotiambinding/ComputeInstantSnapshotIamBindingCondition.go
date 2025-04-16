// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstantsnapshotiambinding


type ComputeInstantSnapshotIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/compute_instant_snapshot_iam_binding#expression ComputeInstantSnapshotIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/compute_instant_snapshot_iam_binding#title ComputeInstantSnapshotIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/compute_instant_snapshot_iam_binding#description ComputeInstantSnapshotIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

