package endpointsserviceiammember


type EndpointsServiceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/endpoints_service_iam_member#expression EndpointsServiceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/endpoints_service_iam_member#title EndpointsServiceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/endpoints_service_iam_member#description EndpointsServiceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

