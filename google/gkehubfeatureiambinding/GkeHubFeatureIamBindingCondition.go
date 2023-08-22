package gkehubfeatureiambinding


type GkeHubFeatureIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/gke_hub_feature_iam_binding#expression GkeHubFeatureIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/gke_hub_feature_iam_binding#title GkeHubFeatureIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/gke_hub_feature_iam_binding#description GkeHubFeatureIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

