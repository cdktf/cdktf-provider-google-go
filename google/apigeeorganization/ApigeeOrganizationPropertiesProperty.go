package apigeeorganization


type ApigeeOrganizationPropertiesProperty struct {
	// Name of the property.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/apigee_organization#name ApigeeOrganization#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Value of the property.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/apigee_organization#value ApigeeOrganization#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

