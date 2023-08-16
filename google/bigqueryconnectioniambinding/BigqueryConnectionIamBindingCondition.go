package bigqueryconnectioniambinding


type BigqueryConnectionIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_connection_iam_binding#expression BigqueryConnectionIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_connection_iam_binding#title BigqueryConnectionIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/bigquery_connection_iam_binding#description BigqueryConnectionIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

