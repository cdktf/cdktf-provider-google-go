package healthcarefhirstore


type HealthcareFhirStoreStreamConfigsBigqueryDestination struct {
	// BigQuery URI to a dataset, up to 2000 characters long, in the format bq://projectId.bqDatasetId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/healthcare_fhir_store#dataset_uri HealthcareFhirStore#dataset_uri}
	DatasetUri *string `field:"required" json:"datasetUri" yaml:"datasetUri"`
	// schema_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/healthcare_fhir_store#schema_config HealthcareFhirStore#schema_config}
	SchemaConfig *HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfig `field:"required" json:"schemaConfig" yaml:"schemaConfig"`
}

