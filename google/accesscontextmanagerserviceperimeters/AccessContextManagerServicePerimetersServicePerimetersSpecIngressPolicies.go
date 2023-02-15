package accesscontextmanagerserviceperimeters


type AccessContextManagerServicePerimetersServicePerimetersSpecIngressPolicies struct {
	// ingress_from block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/access_context_manager_service_perimeters#ingress_from AccessContextManagerServicePerimeters#ingress_from}
	IngressFrom *AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressFrom `field:"optional" json:"ingressFrom" yaml:"ingressFrom"`
	// ingress_to block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/access_context_manager_service_perimeters#ingress_to AccessContextManagerServicePerimeters#ingress_to}
	IngressTo *AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressTo `field:"optional" json:"ingressTo" yaml:"ingressTo"`
}

