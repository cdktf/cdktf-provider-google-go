package cloudassetfolderfeed


type CloudAssetFolderFeedFeedOutputConfig struct {
	// pubsub_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloud_asset_folder_feed#pubsub_destination CloudAssetFolderFeed#pubsub_destination}
	PubsubDestination *CloudAssetFolderFeedFeedOutputConfigPubsubDestination `field:"required" json:"pubsubDestination" yaml:"pubsubDestination"`
}

