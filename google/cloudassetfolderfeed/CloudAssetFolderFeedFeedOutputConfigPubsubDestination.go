package cloudassetfolderfeed


type CloudAssetFolderFeedFeedOutputConfigPubsubDestination struct {
	// Destination on Cloud Pubsub topic.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_asset_folder_feed#topic CloudAssetFolderFeed#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
}

