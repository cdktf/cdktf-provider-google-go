package computeimageiammember


type ComputeImageIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_image_iam_member#expression ComputeImageIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_image_iam_member#title ComputeImageIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_image_iam_member#description ComputeImageIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

