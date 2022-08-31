// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type PrivatecaCertificateTemplateIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate_template_iam_member#expression PrivatecaCertificateTemplateIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate_template_iam_member#title PrivatecaCertificateTemplateIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/privateca_certificate_template_iam_member#description PrivatecaCertificateTemplateIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

