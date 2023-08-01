package cloudassetorganizationfeed


type CloudAssetOrganizationFeedFeedOutputConfig struct {
	// pubsub_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/cloud_asset_organization_feed#pubsub_destination CloudAssetOrganizationFeed#pubsub_destination}
	PubsubDestination *CloudAssetOrganizationFeedFeedOutputConfigPubsubDestination `field:"required" json:"pubsubDestination" yaml:"pubsubDestination"`
}

