package computeresourcepolicy


type ComputeResourcePolicyInstanceSchedulePolicyVmStartSchedule struct {
	// Specifies the frequency for the operation, using the unix-cron format.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_resource_policy#schedule ComputeResourcePolicy#schedule}
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
}

