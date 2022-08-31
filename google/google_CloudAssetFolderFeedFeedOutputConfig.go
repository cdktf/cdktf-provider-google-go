// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type CloudAssetFolderFeedFeedOutputConfig struct {
	// pubsub_destination block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_asset_folder_feed#pubsub_destination CloudAssetFolderFeed#pubsub_destination}
	PubsubDestination *CloudAssetFolderFeedFeedOutputConfigPubsubDestination `field:"required" json:"pubsubDestination" yaml:"pubsubDestination"`
}

