package iapwebtypeappengineiambinding


type IapWebTypeAppEngineIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/iap_web_type_app_engine_iam_binding#expression IapWebTypeAppEngineIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/iap_web_type_app_engine_iam_binding#title IapWebTypeAppEngineIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/iap_web_type_app_engine_iam_binding#description IapWebTypeAppEngineIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

