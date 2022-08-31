// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type OrganizationPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/organization_policy#create OrganizationPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/organization_policy#delete OrganizationPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/organization_policy#read OrganizationPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/organization_policy#update OrganizationPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

