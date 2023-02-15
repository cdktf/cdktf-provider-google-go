package projectorganizationpolicy


type ProjectOrganizationPolicyListPolicyAllow struct {
	// The policy allows or denies all values.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/project_organization_policy#all ProjectOrganizationPolicy#all}
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// The policy can define specific values that are allowed or denied.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/project_organization_policy#values ProjectOrganizationPolicy#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

