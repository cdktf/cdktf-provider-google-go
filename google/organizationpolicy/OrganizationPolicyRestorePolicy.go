package organizationpolicy


type OrganizationPolicyRestorePolicy struct {
	// May only be set to true. If set, then the default Policy is restored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/organization_policy#default OrganizationPolicy#default}
	Default interface{} `field:"required" json:"default" yaml:"default"`
}

