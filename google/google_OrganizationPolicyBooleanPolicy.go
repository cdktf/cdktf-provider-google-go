// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type OrganizationPolicyBooleanPolicy struct {
	// If true, then the Policy is enforced. If false, then any configuration is acceptable.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/organization_policy#enforced OrganizationPolicy#enforced}
	Enforced interface{} `field:"required" json:"enforced" yaml:"enforced"`
}

