package accesscontextmanagerserviceperimeter


type AccessContextManagerServicePerimeterSpecEgressPolicies struct {
	// egress_from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/access_context_manager_service_perimeter#egress_from AccessContextManagerServicePerimeter#egress_from}
	EgressFrom *AccessContextManagerServicePerimeterSpecEgressPoliciesEgressFrom `field:"optional" json:"egressFrom" yaml:"egressFrom"`
	// egress_to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/access_context_manager_service_perimeter#egress_to AccessContextManagerServicePerimeter#egress_to}
	EgressTo *AccessContextManagerServicePerimeterSpecEgressPoliciesEgressTo `field:"optional" json:"egressTo" yaml:"egressTo"`
}

