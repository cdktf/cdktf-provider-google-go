package storagetransferjob


type StorageTransferJobScheduleStartTimeOfDay struct {
	// Hours of day in 24 hour format. Should be from 0 to 23.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/storage_transfer_job#hours StorageTransferJob#hours}
	Hours *float64 `field:"required" json:"hours" yaml:"hours"`
	// Minutes of hour of day. Must be from 0 to 59.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/storage_transfer_job#minutes StorageTransferJob#minutes}
	Minutes *float64 `field:"required" json:"minutes" yaml:"minutes"`
	// Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/storage_transfer_job#nanos StorageTransferJob#nanos}
	Nanos *float64 `field:"required" json:"nanos" yaml:"nanos"`
	// Seconds of minutes of the time. Must normally be from 0 to 59.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/storage_transfer_job#seconds StorageTransferJob#seconds}
	Seconds *float64 `field:"required" json:"seconds" yaml:"seconds"`
}

