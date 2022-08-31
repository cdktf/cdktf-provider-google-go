// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ComputeRegionInstanceGroupManagerStatefulDisk struct {
	// The device name of the disk to be attached.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_instance_group_manager#device_name ComputeRegionInstanceGroupManager#device_name}
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// A value that prescribes what should happen to the stateful disk when the VM instance is deleted.
	//
	// The available options are NEVER and ON_PERMANENT_INSTANCE_DELETION. NEVER - detach the disk when the VM is deleted, but do not delete the disk. ON_PERMANENT_INSTANCE_DELETION will delete the stateful disk when the VM is permanently deleted from the instance group. The default is NEVER.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_instance_group_manager#delete_rule ComputeRegionInstanceGroupManager#delete_rule}
	DeleteRule *string `field:"optional" json:"deleteRule" yaml:"deleteRule"`
}

