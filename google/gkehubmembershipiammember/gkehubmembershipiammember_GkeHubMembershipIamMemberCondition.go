package gkehubmembershipiammember


type GkeHubMembershipIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/gke_hub_membership_iam_member#expression GkeHubMembershipIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/gke_hub_membership_iam_member#title GkeHubMembershipIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/gke_hub_membership_iam_member#description GkeHubMembershipIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

