// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcarepipelinejob


type HealthcarePipelineJobMappingPipelineJobFhirStreamingSource struct {
	// The path to the FHIR store in the format projects/{projectId}/locations/{locationId}/datasets/{datasetId}/fhirStores/{fhirStoreId}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/healthcare_pipeline_job#fhir_store HealthcarePipelineJob#fhir_store}
	FhirStore *string `field:"required" json:"fhirStore" yaml:"fhirStore"`
	// Describes the streaming FHIR data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.33.0/docs/resources/healthcare_pipeline_job#description HealthcarePipelineJob#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

