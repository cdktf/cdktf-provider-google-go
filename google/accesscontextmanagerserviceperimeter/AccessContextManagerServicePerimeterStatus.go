package accesscontextmanagerserviceperimeter


type AccessContextManagerServicePerimeterStatus struct {
	// A list of AccessLevel resource names that allow resources within the ServicePerimeter to be accessed from the internet.
	//
	// AccessLevels listed must be in the same policy as this
	// ServicePerimeter. Referencing a nonexistent AccessLevel is a
	// syntax error. If no AccessLevel names are listed, resources within
	// the perimeter can only be accessed via GCP calls with request
	// origins within the perimeter. For Service Perimeter Bridge, must
	// be empty.
	//
	// Format: accessPolicies/{policy_id}/accessLevels/{access_level_name}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/access_context_manager_service_perimeter#access_levels AccessContextManagerServicePerimeter#access_levels}
	AccessLevels *[]*string `field:"optional" json:"accessLevels" yaml:"accessLevels"`
	// egress_policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/access_context_manager_service_perimeter#egress_policies AccessContextManagerServicePerimeter#egress_policies}
	EgressPolicies interface{} `field:"optional" json:"egressPolicies" yaml:"egressPolicies"`
	// ingress_policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/access_context_manager_service_perimeter#ingress_policies AccessContextManagerServicePerimeter#ingress_policies}
	IngressPolicies interface{} `field:"optional" json:"ingressPolicies" yaml:"ingressPolicies"`
	// A list of GCP resources that are inside of the service perimeter. Currently only projects are allowed. Format: projects/{project_number}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/access_context_manager_service_perimeter#resources AccessContextManagerServicePerimeter#resources}
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// GCP services that are subject to the Service Perimeter restrictions.
	//
	// Must contain a list of services. For example, if
	// 'storage.googleapis.com' is specified, access to the storage
	// buckets inside the perimeter must meet the perimeter's access
	// restrictions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/access_context_manager_service_perimeter#restricted_services AccessContextManagerServicePerimeter#restricted_services}
	RestrictedServices *[]*string `field:"optional" json:"restrictedServices" yaml:"restrictedServices"`
	// vpc_accessible_services block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/access_context_manager_service_perimeter#vpc_accessible_services AccessContextManagerServicePerimeter#vpc_accessible_services}
	VpcAccessibleServices *AccessContextManagerServicePerimeterStatusVpcAccessibleServices `field:"optional" json:"vpcAccessibleServices" yaml:"vpcAccessibleServices"`
}

