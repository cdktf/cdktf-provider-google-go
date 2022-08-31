// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type CloudfunctionsFunctionIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions_function_iam_binding#expression CloudfunctionsFunctionIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions_function_iam_binding#title CloudfunctionsFunctionIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions_function_iam_binding#description CloudfunctionsFunctionIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

