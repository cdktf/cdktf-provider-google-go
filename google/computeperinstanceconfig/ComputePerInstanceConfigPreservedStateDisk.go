package computeperinstanceconfig


type ComputePerInstanceConfigPreservedStateDisk struct {
	// A unique device name that is reflected into the /dev/ tree of a Linux operating system running within the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_per_instance_config#device_name ComputePerInstanceConfig#device_name}
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// The URI of an existing persistent disk to attach under the specified device-name in the format 'projects/project-id/zones/zone/disks/disk-name'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_per_instance_config#source ComputePerInstanceConfig#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// A value that prescribes what should happen to the stateful disk when the VM instance is deleted.
	//
	// The available options are 'NEVER' and 'ON_PERMANENT_INSTANCE_DELETION'.
	// 'NEVER' - detach the disk when the VM is deleted, but do not delete the disk.
	// 'ON_PERMANENT_INSTANCE_DELETION' will delete the stateful disk when the VM is permanently
	// deleted from the instance group. Default value: "NEVER" Possible values: ["NEVER", "ON_PERMANENT_INSTANCE_DELETION"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_per_instance_config#delete_rule ComputePerInstanceConfig#delete_rule}
	DeleteRule *string `field:"optional" json:"deleteRule" yaml:"deleteRule"`
	// The mode of the disk. Default value: "READ_WRITE" Possible values: ["READ_ONLY", "READ_WRITE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_per_instance_config#mode ComputePerInstanceConfig#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

