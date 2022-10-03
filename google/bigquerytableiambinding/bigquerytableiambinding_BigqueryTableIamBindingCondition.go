package bigquerytableiambinding


type BigqueryTableIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_table_iam_binding#expression BigqueryTableIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_table_iam_binding#title BigqueryTableIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_table_iam_binding#description BigqueryTableIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

