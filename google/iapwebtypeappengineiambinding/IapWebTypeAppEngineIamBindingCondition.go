package iapwebtypeappengineiambinding


type IapWebTypeAppEngineIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_web_type_app_engine_iam_binding#expression IapWebTypeAppEngineIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_web_type_app_engine_iam_binding#title IapWebTypeAppEngineIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/iap_web_type_app_engine_iam_binding#description IapWebTypeAppEngineIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

