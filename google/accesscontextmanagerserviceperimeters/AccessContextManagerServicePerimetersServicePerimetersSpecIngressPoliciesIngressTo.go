package accesscontextmanagerserviceperimeters


type AccessContextManagerServicePerimetersServicePerimetersSpecIngressPoliciesIngressTo struct {
	// operations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/access_context_manager_service_perimeters#operations AccessContextManagerServicePerimeters#operations}
	Operations interface{} `field:"optional" json:"operations" yaml:"operations"`
	// A list of resources, currently only projects in the form 'projects/<projectnumber>', protected by this 'ServicePerimeter' that are allowed to be accessed by sources defined in the corresponding 'IngressFrom'.
	//
	// A request matches if it contains
	// a resource in this list. If '*' is specified for resources,
	// then this 'IngressTo' rule will authorize access to all
	// resources inside the perimeter, provided that the request
	// also matches the 'operations' field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/access_context_manager_service_perimeters#resources AccessContextManagerServicePerimeters#resources}
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
}

