package accesscontextmanagerserviceperimeters


type AccessContextManagerServicePerimetersServicePerimetersStatusIngressPoliciesIngressFromSources struct {
	// An 'AccessLevel' resource name that allow resources within the 'ServicePerimeters' to be accessed from the internet.
	//
	// 'AccessLevels' listed
	// must be in the same policy as this 'ServicePerimeter'. Referencing a nonexistent
	// 'AccessLevel' will cause an error. If no 'AccessLevel' names are listed,
	// resources within the perimeter can only be accessed via Google Cloud calls
	// with request origins within the perimeter.
	// Example 'accessPolicies/MY_POLICY/accessLevels/MY_LEVEL.'
	// If * is specified, then all IngressSources will be allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/access_context_manager_service_perimeters#access_level AccessContextManagerServicePerimeters#access_level}
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
	// A Google Cloud resource that is allowed to ingress the perimeter.
	//
	// Requests from these resources will be allowed to access perimeter data.
	// Currently only projects are allowed. Format 'projects/{project_number}'
	// The project may be in any Google Cloud organization, not just the
	// organization that the perimeter is defined in. '*' is not allowed, the case
	// of allowing all Google Cloud resources only is not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/access_context_manager_service_perimeters#resource AccessContextManagerServicePerimeters#resource}
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
}

