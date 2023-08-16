package healthcarefhirstore


type HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfig struct {
	// The depth for all recursive structures in the output analytics schema.
	//
	// For example, concept in the CodeSystem
	// resource is a recursive structure; when the depth is 2, the CodeSystem table will have a column called
	// concept.concept but not concept.concept.concept. If not specified or set to 0, the server will use the default
	// value 2. The maximum depth allowed is 5.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/healthcare_fhir_store#recursive_structure_depth HealthcareFhirStore#recursive_structure_depth}
	RecursiveStructureDepth *float64 `field:"required" json:"recursiveStructureDepth" yaml:"recursiveStructureDepth"`
	// last_updated_partition_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/healthcare_fhir_store#last_updated_partition_config HealthcareFhirStore#last_updated_partition_config}
	LastUpdatedPartitionConfig *HealthcareFhirStoreStreamConfigsBigqueryDestinationSchemaConfigLastUpdatedPartitionConfig `field:"optional" json:"lastUpdatedPartitionConfig" yaml:"lastUpdatedPartitionConfig"`
	// Specifies the output schema type.
	//
	// ANALYTICS: Analytics schema defined by the FHIR community.
	// See https://github.com/FHIR/sql-on-fhir/blob/master/sql-on-fhir.md.
	// ANALYTICS_V2: Analytics V2, similar to schema defined by the FHIR community, with added support for extensions with one or more occurrences and contained resources in stringified JSON.
	// LOSSLESS: A data-driven schema generated from the fields present in the FHIR data being exported, with no additional simplification. Default value: "ANALYTICS" Possible values: ["ANALYTICS", "ANALYTICS_V2", "LOSSLESS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/healthcare_fhir_store#schema_type HealthcareFhirStore#schema_type}
	SchemaType *string `field:"optional" json:"schemaType" yaml:"schemaType"`
}

