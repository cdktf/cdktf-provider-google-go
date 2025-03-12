// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstantsnapshotiammember


type ComputeInstantSnapshotIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/compute_instant_snapshot_iam_member#expression ComputeInstantSnapshotIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/compute_instant_snapshot_iam_member#title ComputeInstantSnapshotIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.25.0/docs/resources/compute_instant_snapshot_iam_member#description ComputeInstantSnapshotIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

