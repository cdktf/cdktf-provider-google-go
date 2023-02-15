package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobStorageConfig struct {
	// big_query_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_job_trigger#big_query_options DataLossPreventionJobTrigger#big_query_options}
	BigQueryOptions *DataLossPreventionJobTriggerInspectJobStorageConfigBigQueryOptions `field:"optional" json:"bigQueryOptions" yaml:"bigQueryOptions"`
	// cloud_storage_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_job_trigger#cloud_storage_options DataLossPreventionJobTrigger#cloud_storage_options}
	CloudStorageOptions *DataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptions `field:"optional" json:"cloudStorageOptions" yaml:"cloudStorageOptions"`
	// datastore_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_job_trigger#datastore_options DataLossPreventionJobTrigger#datastore_options}
	DatastoreOptions *DataLossPreventionJobTriggerInspectJobStorageConfigDatastoreOptions `field:"optional" json:"datastoreOptions" yaml:"datastoreOptions"`
	// timespan_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_loss_prevention_job_trigger#timespan_config DataLossPreventionJobTrigger#timespan_config}
	TimespanConfig *DataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfig `field:"optional" json:"timespanConfig" yaml:"timespanConfig"`
}

