package computeregiondiskiambinding


type ComputeRegionDiskIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_disk_iam_binding#expression ComputeRegionDiskIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_disk_iam_binding#title ComputeRegionDiskIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_disk_iam_binding#description ComputeRegionDiskIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

