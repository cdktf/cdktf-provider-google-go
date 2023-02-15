package iapwebtypeappengineiammember


type IapWebTypeAppEngineIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_web_type_app_engine_iam_member#expression IapWebTypeAppEngineIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_web_type_app_engine_iam_member#title IapWebTypeAppEngineIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_web_type_app_engine_iam_member#description IapWebTypeAppEngineIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

