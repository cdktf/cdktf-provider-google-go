package endpointsserviceconsumersiammember


type EndpointsServiceConsumersIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/endpoints_service_consumers_iam_member#expression EndpointsServiceConsumersIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/endpoints_service_consumers_iam_member#title EndpointsServiceConsumersIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/endpoints_service_consumers_iam_member#description EndpointsServiceConsumersIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

