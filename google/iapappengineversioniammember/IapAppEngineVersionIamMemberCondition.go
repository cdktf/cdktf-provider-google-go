package iapappengineversioniammember


type IapAppEngineVersionIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.66.0/docs/resources/iap_app_engine_version_iam_member#expression IapAppEngineVersionIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.66.0/docs/resources/iap_app_engine_version_iam_member#title IapAppEngineVersionIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.66.0/docs/resources/iap_app_engine_version_iam_member#description IapAppEngineVersionIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

