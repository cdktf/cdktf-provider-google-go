// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudassetfolderfeed


type CloudAssetFolderFeedFeedOutputConfigPubsubDestination struct {
	// Destination on Cloud Pubsub topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.29.0/docs/resources/cloud_asset_folder_feed#topic CloudAssetFolderFeed#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
}

