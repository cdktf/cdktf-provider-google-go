package storagetransferjob


type StorageTransferJobTransferSpecObjectConditions struct {
	// exclude_prefixes must follow the requirements described for include_prefixes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_transfer_job#exclude_prefixes StorageTransferJob#exclude_prefixes}
	ExcludePrefixes *[]*string `field:"optional" json:"excludePrefixes" yaml:"excludePrefixes"`
	// If include_refixes is specified, objects that satisfy the object conditions must have names that start with one of the include_prefixes and that do not start with any of the exclude_prefixes.
	//
	// If include_prefixes is not specified, all objects except those that have names starting with one of the exclude_prefixes must satisfy the object conditions.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_transfer_job#include_prefixes StorageTransferJob#include_prefixes}
	IncludePrefixes *[]*string `field:"optional" json:"includePrefixes" yaml:"includePrefixes"`
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_transfer_job#max_time_elapsed_since_last_modification StorageTransferJob#max_time_elapsed_since_last_modification}
	MaxTimeElapsedSinceLastModification *string `field:"optional" json:"maxTimeElapsedSinceLastModification" yaml:"maxTimeElapsedSinceLastModification"`
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/storage_transfer_job#min_time_elapsed_since_last_modification StorageTransferJob#min_time_elapsed_since_last_modification}
	MinTimeElapsedSinceLastModification *string `field:"optional" json:"minTimeElapsedSinceLastModification" yaml:"minTimeElapsedSinceLastModification"`
}

