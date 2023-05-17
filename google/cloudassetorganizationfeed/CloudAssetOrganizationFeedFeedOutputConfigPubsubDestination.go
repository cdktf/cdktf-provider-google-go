package cloudassetorganizationfeed


type CloudAssetOrganizationFeedFeedOutputConfigPubsubDestination struct {
	// Destination on Cloud Pubsub topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/cloud_asset_organization_feed#topic CloudAssetOrganizationFeed#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
}

