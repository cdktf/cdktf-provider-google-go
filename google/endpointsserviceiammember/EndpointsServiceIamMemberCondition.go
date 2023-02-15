package endpointsserviceiammember


type EndpointsServiceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/endpoints_service_iam_member#expression EndpointsServiceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/endpoints_service_iam_member#title EndpointsServiceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/endpoints_service_iam_member#description EndpointsServiceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

