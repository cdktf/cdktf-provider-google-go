// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagebucket


type StorageBucketLifecycleRule struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/storage_bucket#action StorageBucket#action}
	Action *StorageBucketLifecycleRuleAction `field:"required" json:"action" yaml:"action"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.39.0/docs/resources/storage_bucket#condition StorageBucket#condition}
	Condition *StorageBucketLifecycleRuleCondition `field:"required" json:"condition" yaml:"condition"`
}

