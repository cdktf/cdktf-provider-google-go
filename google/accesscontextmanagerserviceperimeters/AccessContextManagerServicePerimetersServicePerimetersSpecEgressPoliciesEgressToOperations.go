package accesscontextmanagerserviceperimeters


type AccessContextManagerServicePerimetersServicePerimetersSpecEgressPoliciesEgressToOperations struct {
	// method_selectors block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/access_context_manager_service_perimeters#method_selectors AccessContextManagerServicePerimeters#method_selectors}
	MethodSelectors interface{} `field:"optional" json:"methodSelectors" yaml:"methodSelectors"`
	// The name of the API whose methods or permissions the 'IngressPolicy' or  'EgressPolicy' want to allow.
	//
	// A single 'ApiOperation' with serviceName
	// field set to '*' will allow all methods AND permissions for all services.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/access_context_manager_service_perimeters#service_name AccessContextManagerServicePerimeters#service_name}
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}
