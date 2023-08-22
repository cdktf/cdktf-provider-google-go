package accesscontextmanagerserviceperimeteringresspolicy


type AccessContextManagerServicePerimeterIngressPolicyIngressFrom struct {
	// A list of identities that are allowed access through this ingress policy.
	//
	// Should be in the format of email address. The email address should represent
	// individual user or service account only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/access_context_manager_service_perimeter_ingress_policy#identities AccessContextManagerServicePerimeterIngressPolicy#identities}
	Identities *[]*string `field:"optional" json:"identities" yaml:"identities"`
	// Specifies the type of identities that are allowed access from outside the perimeter.
	//
	// If left unspecified, then members of 'identities' field will be
	// allowed access. Possible values: ["ANY_IDENTITY", "ANY_USER_ACCOUNT", "ANY_SERVICE_ACCOUNT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/access_context_manager_service_perimeter_ingress_policy#identity_type AccessContextManagerServicePerimeterIngressPolicy#identity_type}
	IdentityType *string `field:"optional" json:"identityType" yaml:"identityType"`
	// sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/access_context_manager_service_perimeter_ingress_policy#sources AccessContextManagerServicePerimeterIngressPolicy#sources}
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
}

