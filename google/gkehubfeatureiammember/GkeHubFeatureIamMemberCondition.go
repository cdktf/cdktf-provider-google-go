package gkehubfeatureiammember


type GkeHubFeatureIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/gke_hub_feature_iam_member#expression GkeHubFeatureIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/gke_hub_feature_iam_member#title GkeHubFeatureIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/gke_hub_feature_iam_member#description GkeHubFeatureIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

