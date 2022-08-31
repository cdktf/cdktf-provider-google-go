// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type BinaryAuthorizationAttestorIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/binary_authorization_attestor_iam_binding#expression BinaryAuthorizationAttestorIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/binary_authorization_attestor_iam_binding#title BinaryAuthorizationAttestorIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/binary_authorization_attestor_iam_binding#description BinaryAuthorizationAttestorIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

