package bigquerytableiambinding


type BigqueryTableIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_table_iam_binding#expression BigqueryTableIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_table_iam_binding#title BigqueryTableIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/bigquery_table_iam_binding#description BigqueryTableIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

