package cloudassetprojectfeed


type CloudAssetProjectFeedFeedOutputConfig struct {
	// pubsub_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_asset_project_feed#pubsub_destination CloudAssetProjectFeed#pubsub_destination}
	PubsubDestination *CloudAssetProjectFeedFeedOutputConfigPubsubDestination `field:"required" json:"pubsubDestination" yaml:"pubsubDestination"`
}

