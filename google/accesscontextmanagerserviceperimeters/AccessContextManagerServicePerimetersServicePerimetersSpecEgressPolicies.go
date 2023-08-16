package accesscontextmanagerserviceperimeters


type AccessContextManagerServicePerimetersServicePerimetersSpecEgressPolicies struct {
	// egress_from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/access_context_manager_service_perimeters#egress_from AccessContextManagerServicePerimeters#egress_from}
	EgressFrom *AccessContextManagerServicePerimetersServicePerimetersSpecEgressPoliciesEgressFrom `field:"optional" json:"egressFrom" yaml:"egressFrom"`
	// egress_to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/access_context_manager_service_perimeters#egress_to AccessContextManagerServicePerimeters#egress_to}
	EgressTo *AccessContextManagerServicePerimetersServicePerimetersSpecEgressPoliciesEgressTo `field:"optional" json:"egressTo" yaml:"egressTo"`
}

