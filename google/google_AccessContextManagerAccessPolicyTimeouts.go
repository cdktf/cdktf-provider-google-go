// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type AccessContextManagerAccessPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/access_context_manager_access_policy#create AccessContextManagerAccessPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/access_context_manager_access_policy#delete AccessContextManagerAccessPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/access_context_manager_access_policy#update AccessContextManagerAccessPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

