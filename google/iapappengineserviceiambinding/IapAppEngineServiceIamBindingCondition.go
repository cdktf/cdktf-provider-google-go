package iapappengineserviceiambinding


type IapAppEngineServiceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/iap_app_engine_service_iam_binding#expression IapAppEngineServiceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/iap_app_engine_service_iam_binding#title IapAppEngineServiceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/iap_app_engine_service_iam_binding#description IapAppEngineServiceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

