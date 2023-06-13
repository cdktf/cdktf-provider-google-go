package bigquerydatapolicydatapolicyiammember


type BigqueryDatapolicyDataPolicyIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_datapolicy_data_policy_iam_member#expression BigqueryDatapolicyDataPolicyIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_datapolicy_data_policy_iam_member#title BigqueryDatapolicyDataPolicyIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/bigquery_datapolicy_data_policy_iam_member#description BigqueryDatapolicyDataPolicyIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

