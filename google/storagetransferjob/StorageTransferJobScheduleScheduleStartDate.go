package storagetransferjob


type StorageTransferJobScheduleScheduleStartDate struct {
	// Day of month. Must be from 1 to 31 and valid for the year and month.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/storage_transfer_job#day StorageTransferJob#day}
	Day *float64 `field:"required" json:"day" yaml:"day"`
	// Month of year. Must be from 1 to 12.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/storage_transfer_job#month StorageTransferJob#month}
	Month *float64 `field:"required" json:"month" yaml:"month"`
	// Year of date. Must be from 1 to 9999.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/storage_transfer_job#year StorageTransferJob#year}
	Year *float64 `field:"required" json:"year" yaml:"year"`
}

