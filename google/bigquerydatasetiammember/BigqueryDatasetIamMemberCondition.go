package bigquerydatasetiammember


type BigqueryDatasetIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_dataset_iam_member#expression BigqueryDatasetIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_dataset_iam_member#title BigqueryDatasetIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_dataset_iam_member#description BigqueryDatasetIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

