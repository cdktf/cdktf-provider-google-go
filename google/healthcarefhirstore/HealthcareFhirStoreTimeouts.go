package healthcarefhirstore


type HealthcareFhirStoreTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/healthcare_fhir_store#create HealthcareFhirStore#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/healthcare_fhir_store#delete HealthcareFhirStore#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/healthcare_fhir_store#update HealthcareFhirStore#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

