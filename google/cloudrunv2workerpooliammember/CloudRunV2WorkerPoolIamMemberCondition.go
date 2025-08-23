// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudrunv2workerpooliammember


type CloudRunV2WorkerPoolIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/cloud_run_v2_worker_pool_iam_member#expression CloudRunV2WorkerPoolIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/cloud_run_v2_worker_pool_iam_member#title CloudRunV2WorkerPoolIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/cloud_run_v2_worker_pool_iam_member#description CloudRunV2WorkerPoolIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

