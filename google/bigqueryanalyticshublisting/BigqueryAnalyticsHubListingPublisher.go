package bigqueryanalyticshublisting


type BigqueryAnalyticsHubListingPublisher struct {
	// Name of the listing publisher.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_analytics_hub_listing#name BigqueryAnalyticsHubListing#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Email or URL of the listing publisher.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_analytics_hub_listing#primary_contact BigqueryAnalyticsHubListing#primary_contact}
	PrimaryContact *string `field:"optional" json:"primaryContact" yaml:"primaryContact"`
}
