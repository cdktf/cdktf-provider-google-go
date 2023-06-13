package cloudrundomainmapping


type CloudRunDomainMappingSpec struct {
	// The name of the Cloud Run Service that this DomainMapping applies to. The route must exist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/cloud_run_domain_mapping#route_name CloudRunDomainMapping#route_name}
	RouteName *string `field:"required" json:"routeName" yaml:"routeName"`
	// The mode of the certificate. Default value: "AUTOMATIC" Possible values: ["NONE", "AUTOMATIC"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/cloud_run_domain_mapping#certificate_mode CloudRunDomainMapping#certificate_mode}
	CertificateMode *string `field:"optional" json:"certificateMode" yaml:"certificateMode"`
	// If set, the mapping will override any mapping set before this spec was set.
	//
	// It is recommended that the user leaves this empty to receive an error
	// warning about a potential conflict and only set it once the respective UI
	// has given such a warning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/cloud_run_domain_mapping#force_override CloudRunDomainMapping#force_override}
	ForceOverride interface{} `field:"optional" json:"forceOverride" yaml:"forceOverride"`
}

