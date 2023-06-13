package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobActionsDeidentify struct {
	// User settable Cloud Storage bucket and folders to store de-identified files.
	//
	// This field must be set for cloud storage deidentification.
	//
	// The output Cloud Storage bucket must be different from the input bucket.
	//
	// De-identified files will overwrite files in the output path.
	//
	// Form of: gs://bucket/folder/ or gs://bucket
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_job_trigger#cloud_storage_output DataLossPreventionJobTrigger#cloud_storage_output}
	CloudStorageOutput *string `field:"required" json:"cloudStorageOutput" yaml:"cloudStorageOutput"`
	// List of user-specified file type groups to transform. If specified, only the files with these filetypes will be transformed.
	//
	// If empty, all supported files will be transformed. Supported types may be automatically added over time.
	//
	// If a file type is set in this field that isn't supported by the Deidentify action then the job will fail and will not be successfully created/started. Possible values: ["IMAGE", "TEXT_FILE", "CSV", "TSV"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_job_trigger#file_types_to_transform DataLossPreventionJobTrigger#file_types_to_transform}
	FileTypesToTransform *[]*string `field:"optional" json:"fileTypesToTransform" yaml:"fileTypesToTransform"`
	// transformation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_job_trigger#transformation_config DataLossPreventionJobTrigger#transformation_config}
	TransformationConfig *DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationConfig `field:"optional" json:"transformationConfig" yaml:"transformationConfig"`
	// transformation_details_storage_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_job_trigger#transformation_details_storage_config DataLossPreventionJobTrigger#transformation_details_storage_config}
	TransformationDetailsStorageConfig *DataLossPreventionJobTriggerInspectJobActionsDeidentifyTransformationDetailsStorageConfig `field:"optional" json:"transformationDetailsStorageConfig" yaml:"transformationDetailsStorageConfig"`
}

