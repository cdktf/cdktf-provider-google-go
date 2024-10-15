// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package healthcarepipelinejob


type HealthcarePipelineJobReconciliationPipelineJob struct {
	// Specifies the top level directory of the matching configs used in all mapping pipelines, which extract properties for resources to be matched on.
	//
	// Example: gs://{bucket-id}/{path/to/matching/configs}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/healthcare_pipeline_job#matching_uri_prefix HealthcarePipelineJob#matching_uri_prefix}
	MatchingUriPrefix *string `field:"required" json:"matchingUriPrefix" yaml:"matchingUriPrefix"`
	// merge_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/healthcare_pipeline_job#merge_config HealthcarePipelineJob#merge_config}
	MergeConfig *HealthcarePipelineJobReconciliationPipelineJobMergeConfig `field:"required" json:"mergeConfig" yaml:"mergeConfig"`
	// The harmonized FHIR store to write harmonized FHIR resources to, in the format of: project/{projectID}/locations/{locationID}/datasets/{datasetName}/fhirStores/{id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/healthcare_pipeline_job#fhir_store_destination HealthcarePipelineJob#fhir_store_destination}
	FhirStoreDestination *string `field:"optional" json:"fhirStoreDestination" yaml:"fhirStoreDestination"`
}

