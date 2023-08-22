package bigqueryanalyticshubdataexchangeiambinding


type BigqueryAnalyticsHubDataExchangeIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigquery_analytics_hub_data_exchange_iam_binding#expression BigqueryAnalyticsHubDataExchangeIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigquery_analytics_hub_data_exchange_iam_binding#title BigqueryAnalyticsHubDataExchangeIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/bigquery_analytics_hub_data_exchange_iam_binding#description BigqueryAnalyticsHubDataExchangeIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

