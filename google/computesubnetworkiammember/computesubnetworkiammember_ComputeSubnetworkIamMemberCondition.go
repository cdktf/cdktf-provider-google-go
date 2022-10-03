package computesubnetworkiammember


type ComputeSubnetworkIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_subnetwork_iam_member#expression ComputeSubnetworkIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_subnetwork_iam_member#title ComputeSubnetworkIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_subnetwork_iam_member#description ComputeSubnetworkIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

