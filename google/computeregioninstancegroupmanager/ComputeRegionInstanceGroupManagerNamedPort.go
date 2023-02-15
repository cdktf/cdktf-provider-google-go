package computeregioninstancegroupmanager


type ComputeRegionInstanceGroupManagerNamedPort struct {
	// The name of the port.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_instance_group_manager#name ComputeRegionInstanceGroupManager#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The port number.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_instance_group_manager#port ComputeRegionInstanceGroupManager#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

