package organizationiammember


type OrganizationIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/organization_iam_member#expression OrganizationIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/organization_iam_member#title OrganizationIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/organization_iam_member#description OrganizationIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

