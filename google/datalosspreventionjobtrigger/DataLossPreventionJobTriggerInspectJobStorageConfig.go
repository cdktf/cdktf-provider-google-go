// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobStorageConfig struct {
	// big_query_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/data_loss_prevention_job_trigger#big_query_options DataLossPreventionJobTrigger#big_query_options}
	BigQueryOptions *DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions `field:"optional" json:"bigQueryOptions" yaml:"bigQueryOptions"`
	// cloud_storage_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/data_loss_prevention_job_trigger#cloud_storage_options DataLossPreventionJobTrigger#cloud_storage_options}
	CloudStorageOptions *DataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptions `field:"optional" json:"cloudStorageOptions" yaml:"cloudStorageOptions"`
	// datastore_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/data_loss_prevention_job_trigger#datastore_options DataLossPreventionJobTrigger#datastore_options}
	DatastoreOptions *DataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptions `field:"optional" json:"datastoreOptions" yaml:"datastoreOptions"`
	// hybrid_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/data_loss_prevention_job_trigger#hybrid_options DataLossPreventionJobTrigger#hybrid_options}
	HybridOptions *DataLossPreventionJobTriggerInspectJobStorageConfigHybridOptions `field:"optional" json:"hybridOptions" yaml:"hybridOptions"`
	// timespan_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/data_loss_prevention_job_trigger#timespan_config DataLossPreventionJobTrigger#timespan_config}
	TimespanConfig *DataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfig `field:"optional" json:"timespanConfig" yaml:"timespanConfig"`
}

