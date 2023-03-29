package bigquerydatapolicydatapolicyiambinding


type BigqueryDatapolicyDataPolicyIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_datapolicy_data_policy_iam_binding#expression BigqueryDatapolicyDataPolicyIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_datapolicy_data_policy_iam_binding#title BigqueryDatapolicyDataPolicyIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_datapolicy_data_policy_iam_binding#description BigqueryDatapolicyDataPolicyIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

