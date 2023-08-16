package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfig struct {
	// timestamp_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_job_trigger#timestamp_field DataLossPreventionJobTrigger#timestamp_field}
	TimestampField *DataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfigTimestampField `field:"required" json:"timestampField" yaml:"timestampField"`
	// When the job is started by a JobTrigger we will automatically figure out a valid startTime to avoid scanning files that have not been modified since the last time the JobTrigger executed.
	//
	// This will
	// be based on the time of the execution of the last run of the JobTrigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_job_trigger#enable_auto_population_of_timespan_config DataLossPreventionJobTrigger#enable_auto_population_of_timespan_config}
	EnableAutoPopulationOfTimespanConfig interface{} `field:"optional" json:"enableAutoPopulationOfTimespanConfig" yaml:"enableAutoPopulationOfTimespanConfig"`
	// Exclude files or rows newer than this value. If set to zero, no upper time limit is applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_job_trigger#end_time DataLossPreventionJobTrigger#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// Exclude files or rows older than this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/data_loss_prevention_job_trigger#start_time DataLossPreventionJobTrigger#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

