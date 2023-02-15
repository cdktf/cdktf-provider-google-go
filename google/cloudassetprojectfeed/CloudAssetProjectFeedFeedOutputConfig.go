package cloudassetprojectfeed


type CloudAssetProjectFeedFeedOutputConfig struct {
	// pubsub_destination block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_asset_project_feed#pubsub_destination CloudAssetProjectFeed#pubsub_destination}
	PubsubDestination *CloudAssetProjectFeedFeedOutputConfigPubsubDestination `field:"required" json:"pubsubDestination" yaml:"pubsubDestination"`
}

