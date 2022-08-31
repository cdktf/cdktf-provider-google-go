// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ComputeImageIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_image_iam_binding#expression ComputeImageIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_image_iam_binding#title ComputeImageIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_image_iam_binding#description ComputeImageIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

