package cloudassetfolderfeed


type CloudAssetFolderFeedFeedOutputConfigPubsubDestination struct {
	// Destination on Cloud Pubsub topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_asset_folder_feed#topic CloudAssetFolderFeed#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
}

