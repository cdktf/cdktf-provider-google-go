package accesscontextmanagerserviceperimeters


type AccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServices struct {
	// The list of APIs usable within the Service Perimeter. Must be empty unless 'enableRestriction' is True.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/access_context_manager_service_perimeters#allowed_services AccessContextManagerServicePerimeters#allowed_services}
	AllowedServices *[]*string `field:"optional" json:"allowedServices" yaml:"allowedServices"`
	// Whether to restrict API calls within the Service Perimeter to the list of APIs specified in 'allowedServices'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/access_context_manager_service_perimeters#enable_restriction AccessContextManagerServicePerimeters#enable_restriction}
	EnableRestriction interface{} `field:"optional" json:"enableRestriction" yaml:"enableRestriction"`
}

