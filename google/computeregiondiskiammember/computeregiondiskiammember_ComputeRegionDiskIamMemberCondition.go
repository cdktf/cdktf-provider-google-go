package computeregiondiskiammember


type ComputeRegionDiskIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_disk_iam_member#expression ComputeRegionDiskIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_disk_iam_member#title ComputeRegionDiskIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_disk_iam_member#description ComputeRegionDiskIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

