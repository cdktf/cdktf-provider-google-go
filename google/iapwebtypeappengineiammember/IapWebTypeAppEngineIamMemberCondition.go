package iapwebtypeappengineiammember


type IapWebTypeAppEngineIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iap_web_type_app_engine_iam_member#expression IapWebTypeAppEngineIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iap_web_type_app_engine_iam_member#title IapWebTypeAppEngineIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/iap_web_type_app_engine_iam_member#description IapWebTypeAppEngineIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

