// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type IapWebIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_web_iam_member#expression IapWebIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_web_iam_member#title IapWebIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_web_iam_member#description IapWebIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

