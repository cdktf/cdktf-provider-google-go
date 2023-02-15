package cloudassetprojectfeed


type CloudAssetProjectFeedFeedOutputConfigPubsubDestination struct {
	// Destination on Cloud Pubsub topic.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_asset_project_feed#topic CloudAssetProjectFeed#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
}

