package accesscontextmanagerserviceperimeter


type AccessContextManagerServicePerimeterStatusEgressPolicies struct {
	// egress_from block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/access_context_manager_service_perimeter#egress_from AccessContextManagerServicePerimeter#egress_from}
	EgressFrom *AccessContextManagerServicePerimeterStatusEgressPoliciesEgressFrom `field:"optional" json:"egressFrom" yaml:"egressFrom"`
	// egress_to block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/access_context_manager_service_perimeter#egress_to AccessContextManagerServicePerimeter#egress_to}
	EgressTo *AccessContextManagerServicePerimeterStatusEgressPoliciesEgressTo `field:"optional" json:"egressTo" yaml:"egressTo"`
}

