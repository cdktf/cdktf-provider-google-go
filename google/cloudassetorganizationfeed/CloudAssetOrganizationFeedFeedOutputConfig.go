package cloudassetorganizationfeed


type CloudAssetOrganizationFeedFeedOutputConfig struct {
	// pubsub_destination block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloud_asset_organization_feed#pubsub_destination CloudAssetOrganizationFeed#pubsub_destination}
	PubsubDestination *CloudAssetOrganizationFeedFeedOutputConfigPubsubDestination `field:"required" json:"pubsubDestination" yaml:"pubsubDestination"`
}

