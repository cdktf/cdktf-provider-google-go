// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcarepipelinejob


type HealthcarePipelineJobMappingPipelineJob struct {
	// mapping_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/healthcare_pipeline_job#mapping_config HealthcarePipelineJob#mapping_config}
	MappingConfig *HealthcarePipelineJobMappingPipelineJobMappingConfig `field:"required" json:"mappingConfig" yaml:"mappingConfig"`
	// If set, the mapping pipeline will write snapshots to this FHIR store without assigning stable IDs.
	//
	// You must
	// grant your pipeline project's Cloud Healthcare Service
	// Agent serviceaccount healthcare.fhirResources.executeBundle
	// and healthcare.fhirResources.create permissions on the
	// destination store. The destination store must set
	// [disableReferentialIntegrity][FhirStore.disable_referential_integrity]
	// to true. The destination store must use FHIR version R4.
	// Format: project/{projectID}/locations/{locationID}/datasets/{datasetName}/fhirStores/{fhirStoreID}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/healthcare_pipeline_job#fhir_store_destination HealthcarePipelineJob#fhir_store_destination}
	FhirStoreDestination *string `field:"optional" json:"fhirStoreDestination" yaml:"fhirStoreDestination"`
	// fhir_streaming_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/healthcare_pipeline_job#fhir_streaming_source HealthcarePipelineJob#fhir_streaming_source}
	FhirStreamingSource *HealthcarePipelineJobMappingPipelineJobFhirStreamingSource `field:"optional" json:"fhirStreamingSource" yaml:"fhirStreamingSource"`
	// If set to true, a mapping pipeline will send output snapshots to the reconciliation pipeline in its dataset.
	//
	// A reconciliation
	// pipeline must exist in this dataset before a mapping pipeline
	// with a reconciliation destination can be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/healthcare_pipeline_job#reconciliation_destination HealthcarePipelineJob#reconciliation_destination}
	ReconciliationDestination interface{} `field:"optional" json:"reconciliationDestination" yaml:"reconciliationDestination"`
}

