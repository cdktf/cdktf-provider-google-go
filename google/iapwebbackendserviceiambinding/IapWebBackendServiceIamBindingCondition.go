package iapwebbackendserviceiambinding


type IapWebBackendServiceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/iap_web_backend_service_iam_binding#expression IapWebBackendServiceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/iap_web_backend_service_iam_binding#title IapWebBackendServiceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/iap_web_backend_service_iam_binding#description IapWebBackendServiceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

