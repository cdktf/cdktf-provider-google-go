// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudassetfolderfeed

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudAssetFolderFeedConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The project whose identity will be used when sending messages to the destination pubsub topic.
	//
	// It also specifies the project for API
	// enablement check, quota, and billing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/cloud_asset_folder_feed#billing_project CloudAssetFolderFeed#billing_project}
	BillingProject *string `field:"required" json:"billingProject" yaml:"billingProject"`
	// This is the client-assigned asset feed identifier and it needs to be unique under a specific parent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/cloud_asset_folder_feed#feed_id CloudAssetFolderFeed#feed_id}
	FeedId *string `field:"required" json:"feedId" yaml:"feedId"`
	// feed_output_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/cloud_asset_folder_feed#feed_output_config CloudAssetFolderFeed#feed_output_config}
	FeedOutputConfig *CloudAssetFolderFeedFeedOutputConfig `field:"required" json:"feedOutputConfig" yaml:"feedOutputConfig"`
	// The folder this feed should be created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/cloud_asset_folder_feed#folder CloudAssetFolderFeed#folder}
	Folder *string `field:"required" json:"folder" yaml:"folder"`
	// A list of the full names of the assets to receive updates.
	//
	// You must specify either or both of
	// assetNames and assetTypes. Only asset updates matching specified assetNames and assetTypes are
	// exported to the feed. For example: //compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1.
	// See https://cloud.google.com/apis/design/resourceNames#fullResourceName for more info.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/cloud_asset_folder_feed#asset_names CloudAssetFolderFeed#asset_names}
	AssetNames *[]*string `field:"optional" json:"assetNames" yaml:"assetNames"`
	// A list of types of the assets to receive updates.
	//
	// You must specify either or both of assetNames
	// and assetTypes. Only asset updates matching specified assetNames and assetTypes are exported to
	// the feed. For example: "compute.googleapis.com/Disk"
	// See https://cloud.google.com/asset-inventory/docs/supported-asset-types for a list of all
	// supported asset types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/cloud_asset_folder_feed#asset_types CloudAssetFolderFeed#asset_types}
	AssetTypes *[]*string `field:"optional" json:"assetTypes" yaml:"assetTypes"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/cloud_asset_folder_feed#condition CloudAssetFolderFeed#condition}
	Condition *CloudAssetFolderFeedCondition `field:"optional" json:"condition" yaml:"condition"`
	// Asset content type.
	//
	// If not specified, no content but the asset name and type will be returned. Possible values: ["CONTENT_TYPE_UNSPECIFIED", "RESOURCE", "IAM_POLICY", "ORG_POLICY", "OS_INVENTORY", "ACCESS_POLICY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/cloud_asset_folder_feed#content_type CloudAssetFolderFeed#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/cloud_asset_folder_feed#id CloudAssetFolderFeed#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.13.0/docs/resources/cloud_asset_folder_feed#timeouts CloudAssetFolderFeed#timeouts}
	Timeouts *CloudAssetFolderFeedTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

