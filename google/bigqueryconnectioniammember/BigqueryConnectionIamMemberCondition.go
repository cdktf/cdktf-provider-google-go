package bigqueryconnectioniammember


type BigqueryConnectionIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_connection_iam_member#expression BigqueryConnectionIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_connection_iam_member#title BigqueryConnectionIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_connection_iam_member#description BigqueryConnectionIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

