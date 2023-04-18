package computeresourcepolicy


type ComputeResourcePolicyInstanceSchedulePolicyVmStartSchedule struct {
	// Specifies the frequency for the operation, using the unix-cron format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.62.0/docs/resources/compute_resource_policy#schedule ComputeResourcePolicy#schedule}
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
}

