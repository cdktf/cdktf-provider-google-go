package bigqueryanalyticshublisting


type BigqueryAnalyticsHubListingDataProvider struct {
	// Name of the data provider.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_analytics_hub_listing#name BigqueryAnalyticsHubListing#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Email or URL of the data provider.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_analytics_hub_listing#primary_contact BigqueryAnalyticsHubListing#primary_contact}
	PrimaryContact *string `field:"optional" json:"primaryContact" yaml:"primaryContact"`
}

