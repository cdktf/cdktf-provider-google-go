package healthcarefhirstore


type HealthcareFhirStoreStreamConfigs struct {
	// bigquery_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/healthcare_fhir_store#bigquery_destination HealthcareFhirStore#bigquery_destination}
	BigqueryDestination *HealthcareFhirStoreStreamConfigsBigqueryDestination `field:"required" json:"bigqueryDestination" yaml:"bigqueryDestination"`
	// Supply a FHIR resource type (such as "Patient" or "Observation").
	//
	// See
	// https://www.hl7.org/fhir/valueset-resource-types.html for a list of all FHIR resource types. The server treats
	// an empty list as an intent to stream all the supported resource types in this FHIR store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/healthcare_fhir_store#resource_types HealthcareFhirStore#resource_types}
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
}

