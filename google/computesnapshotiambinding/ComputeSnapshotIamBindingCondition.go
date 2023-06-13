package computesnapshotiambinding


type ComputeSnapshotIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_snapshot_iam_binding#expression ComputeSnapshotIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_snapshot_iam_binding#title ComputeSnapshotIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_snapshot_iam_binding#description ComputeSnapshotIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

