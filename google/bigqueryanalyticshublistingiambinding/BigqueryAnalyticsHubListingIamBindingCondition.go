package bigqueryanalyticshublistingiambinding


type BigqueryAnalyticsHubListingIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigquery_analytics_hub_listing_iam_binding#expression BigqueryAnalyticsHubListingIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigquery_analytics_hub_listing_iam_binding#title BigqueryAnalyticsHubListingIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigquery_analytics_hub_listing_iam_binding#description BigqueryAnalyticsHubListingIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

