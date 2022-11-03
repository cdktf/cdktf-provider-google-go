package bigqueryanalyticshublisting


type BigqueryAnalyticsHubListingBigqueryDataset struct {
	// Resource name of the dataset source for this listing. e.g. projects/myproject/datasets/123.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_analytics_hub_listing#dataset BigqueryAnalyticsHubListing#dataset}
	Dataset *string `field:"required" json:"dataset" yaml:"dataset"`
}

