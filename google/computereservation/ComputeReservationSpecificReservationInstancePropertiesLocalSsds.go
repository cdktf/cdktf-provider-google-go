package computereservation


type ComputeReservationSpecificReservationInstancePropertiesLocalSsds struct {
	// The size of the disk in base-2 GB.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_reservation#disk_size_gb ComputeReservation#disk_size_gb}
	DiskSizeGb *float64 `field:"required" json:"diskSizeGb" yaml:"diskSizeGb"`
	// The disk interface to use for attaching this disk. Default value: "SCSI" Possible values: ["SCSI", "NVME"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_reservation#interface ComputeReservation#interface}
	Interface *string `field:"optional" json:"interface" yaml:"interface"`
}

