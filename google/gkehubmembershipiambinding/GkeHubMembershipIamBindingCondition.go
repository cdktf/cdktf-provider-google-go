package gkehubmembershipiambinding


type GkeHubMembershipIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/gke_hub_membership_iam_binding#expression GkeHubMembershipIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/gke_hub_membership_iam_binding#title GkeHubMembershipIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/gke_hub_membership_iam_binding#description GkeHubMembershipIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

