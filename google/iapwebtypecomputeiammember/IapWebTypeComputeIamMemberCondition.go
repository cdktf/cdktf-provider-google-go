package iapwebtypecomputeiammember


type IapWebTypeComputeIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/iap_web_type_compute_iam_member#expression IapWebTypeComputeIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/iap_web_type_compute_iam_member#title IapWebTypeComputeIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/iap_web_type_compute_iam_member#description IapWebTypeComputeIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

