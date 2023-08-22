package apigeeorganization


type ApigeeOrganizationPropertiesProperty struct {
	// Name of the property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/apigee_organization#name ApigeeOrganization#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Value of the property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/apigee_organization#value ApigeeOrganization#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

