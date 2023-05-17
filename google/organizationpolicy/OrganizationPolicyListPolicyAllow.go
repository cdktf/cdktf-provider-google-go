package organizationpolicy


type OrganizationPolicyListPolicyAllow struct {
	// The policy allows or denies all values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/organization_policy#all OrganizationPolicy#all}
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// The policy can define specific values that are allowed or denied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/organization_policy#values OrganizationPolicy#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

