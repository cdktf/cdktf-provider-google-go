package computesnapshotiammember


type ComputeSnapshotIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_snapshot_iam_member#expression ComputeSnapshotIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_snapshot_iam_member#title ComputeSnapshotIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_snapshot_iam_member#description ComputeSnapshotIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

