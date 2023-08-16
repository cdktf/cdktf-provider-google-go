package accesscontextmanagerserviceperimeter


type AccessContextManagerServicePerimeterStatusIngressPolicies struct {
	// ingress_from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/access_context_manager_service_perimeter#ingress_from AccessContextManagerServicePerimeter#ingress_from}
	IngressFrom *AccessContextManagerServicePerimeterStatusIngressPoliciesIngressFrom `field:"optional" json:"ingressFrom" yaml:"ingressFrom"`
	// ingress_to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/access_context_manager_service_perimeter#ingress_to AccessContextManagerServicePerimeter#ingress_to}
	IngressTo *AccessContextManagerServicePerimeterStatusIngressPoliciesIngressTo `field:"optional" json:"ingressTo" yaml:"ingressTo"`
}

