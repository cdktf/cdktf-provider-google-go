package iapwebregionbackendserviceiammember


type IapWebRegionBackendServiceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/iap_web_region_backend_service_iam_member#expression IapWebRegionBackendServiceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/iap_web_region_backend_service_iam_member#title IapWebRegionBackendServiceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/iap_web_region_backend_service_iam_member#description IapWebRegionBackendServiceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

