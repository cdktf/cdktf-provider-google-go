package computeinstance


type ComputeInstanceScratchDisk struct {
	// The disk interface used for attaching this disk. One of SCSI or NVME.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance#interface ComputeInstance#interface}
	Interface *string `field:"required" json:"interface" yaml:"interface"`
}

