// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudassetprojectfeed


type CloudAssetProjectFeedFeedOutputConfigPubsubDestination struct {
	// Destination on Cloud Pubsub topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.16.0/docs/resources/cloud_asset_project_feed#topic CloudAssetProjectFeed#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
}

