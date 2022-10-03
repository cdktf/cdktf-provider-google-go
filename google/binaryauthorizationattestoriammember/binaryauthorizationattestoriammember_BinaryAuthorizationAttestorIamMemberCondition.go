package binaryauthorizationattestoriammember


type BinaryAuthorizationAttestorIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/binary_authorization_attestor_iam_member#expression BinaryAuthorizationAttestorIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/binary_authorization_attestor_iam_member#title BinaryAuthorizationAttestorIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/binary_authorization_attestor_iam_member#description BinaryAuthorizationAttestorIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

