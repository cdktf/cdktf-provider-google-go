package cloudbuildv2connectioniammember


type Cloudbuildv2ConnectionIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuildv2_connection_iam_member#expression Cloudbuildv2ConnectionIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuildv2_connection_iam_member#title Cloudbuildv2ConnectionIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/cloudbuildv2_connection_iam_member#description Cloudbuildv2ConnectionIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

