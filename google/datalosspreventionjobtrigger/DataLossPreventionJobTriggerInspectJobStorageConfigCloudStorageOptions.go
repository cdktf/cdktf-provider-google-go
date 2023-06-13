package datalosspreventionjobtrigger


type DataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptions struct {
	// file_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_job_trigger#file_set DataLossPreventionJobTrigger#file_set}
	FileSet *DataLossPreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet `field:"required" json:"fileSet" yaml:"fileSet"`
	// Max number of bytes to scan from a file.
	//
	// If a scanned file's size is bigger than this value
	// then the rest of the bytes are omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_job_trigger#bytes_limit_per_file DataLossPreventionJobTrigger#bytes_limit_per_file}
	BytesLimitPerFile *float64 `field:"optional" json:"bytesLimitPerFile" yaml:"bytesLimitPerFile"`
	// Max percentage of bytes to scan from a file.
	//
	// The rest are omitted. The number of bytes scanned is rounded down.
	// Must be between 0 and 100, inclusively. Both 0 and 100 means no limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_job_trigger#bytes_limit_per_file_percent DataLossPreventionJobTrigger#bytes_limit_per_file_percent}
	BytesLimitPerFilePercent *float64 `field:"optional" json:"bytesLimitPerFilePercent" yaml:"bytesLimitPerFilePercent"`
	// Limits the number of files to scan to this percentage of the input FileSet.
	//
	// Number of files scanned is rounded down.
	// Must be between 0 and 100, inclusively. Both 0 and 100 means no limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_job_trigger#files_limit_percent DataLossPreventionJobTrigger#files_limit_percent}
	FilesLimitPercent *float64 `field:"optional" json:"filesLimitPercent" yaml:"filesLimitPercent"`
	// List of file type groups to include in the scan.
	//
	// If empty, all files are scanned and available data
	// format processors are applied. In addition, the binary content of the selected files is always scanned as well.
	// Images are scanned only as binary if the specified region does not support image inspection and no fileTypes were specified. Possible values: ["BINARY_FILE", "TEXT_FILE", "IMAGE", "WORD", "PDF", "AVRO", "CSV", "TSV"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_job_trigger#file_types DataLossPreventionJobTrigger#file_types}
	FileTypes *[]*string `field:"optional" json:"fileTypes" yaml:"fileTypes"`
	// How to sample bytes if not all bytes are scanned.
	//
	// Meaningful only when used in conjunction with bytesLimitPerFile.
	// If not specified, scanning would start from the top. Possible values: ["TOP", "RANDOM_START"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_loss_prevention_job_trigger#sample_method DataLossPreventionJobTrigger#sample_method}
	SampleMethod *string `field:"optional" json:"sampleMethod" yaml:"sampleMethod"`
}

