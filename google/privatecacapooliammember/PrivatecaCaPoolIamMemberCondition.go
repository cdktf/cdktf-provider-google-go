package privatecacapooliammember


type PrivatecaCaPoolIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_ca_pool_iam_member#expression PrivatecaCaPoolIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_ca_pool_iam_member#title PrivatecaCaPoolIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/privateca_ca_pool_iam_member#description PrivatecaCaPoolIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

