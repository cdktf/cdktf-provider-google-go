package bigqueryanalyticshublistingiammember


type BigqueryAnalyticsHubListingIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_analytics_hub_listing_iam_member#expression BigqueryAnalyticsHubListingIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_analytics_hub_listing_iam_member#title BigqueryAnalyticsHubListingIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_analytics_hub_listing_iam_member#description BigqueryAnalyticsHubListingIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

