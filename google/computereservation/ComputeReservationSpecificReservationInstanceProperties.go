package computereservation


type ComputeReservationSpecificReservationInstanceProperties struct {
	// The name of the machine type to reserve.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_reservation#machine_type ComputeReservation#machine_type}
	MachineType *string `field:"required" json:"machineType" yaml:"machineType"`
	// guest_accelerators block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_reservation#guest_accelerators ComputeReservation#guest_accelerators}
	GuestAccelerators interface{} `field:"optional" json:"guestAccelerators" yaml:"guestAccelerators"`
	// local_ssds block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_reservation#local_ssds ComputeReservation#local_ssds}
	LocalSsds interface{} `field:"optional" json:"localSsds" yaml:"localSsds"`
	// The minimum CPU platform for the reservation.
	//
	// For example,
	// '"Intel Skylake"'. See
	// the CPU platform availability reference](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform#availablezones)
	// for information on available CPU platforms.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_reservation#min_cpu_platform ComputeReservation#min_cpu_platform}
	MinCpuPlatform *string `field:"optional" json:"minCpuPlatform" yaml:"minCpuPlatform"`
}

