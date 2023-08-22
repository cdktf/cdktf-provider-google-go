package dataplextaskiammember


type DataplexTaskIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataplex_task_iam_member#expression DataplexTaskIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataplex_task_iam_member#title DataplexTaskIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataplex_task_iam_member#description DataplexTaskIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

