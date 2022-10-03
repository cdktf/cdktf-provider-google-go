package computediskiambinding


type ComputeDiskIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_disk_iam_binding#expression ComputeDiskIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_disk_iam_binding#title ComputeDiskIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_disk_iam_binding#description ComputeDiskIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

