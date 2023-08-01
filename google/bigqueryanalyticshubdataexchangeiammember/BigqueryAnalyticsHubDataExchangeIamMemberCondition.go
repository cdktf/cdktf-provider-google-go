package bigqueryanalyticshubdataexchangeiammember


type BigqueryAnalyticsHubDataExchangeIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/bigquery_analytics_hub_data_exchange_iam_member#expression BigqueryAnalyticsHubDataExchangeIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/bigquery_analytics_hub_data_exchange_iam_member#title BigqueryAnalyticsHubDataExchangeIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/bigquery_analytics_hub_data_exchange_iam_member#description BigqueryAnalyticsHubDataExchangeIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

