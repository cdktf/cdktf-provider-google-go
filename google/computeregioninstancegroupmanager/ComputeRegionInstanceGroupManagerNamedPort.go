package computeregioninstancegroupmanager


type ComputeRegionInstanceGroupManagerNamedPort struct {
	// The name of the port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_instance_group_manager#name ComputeRegionInstanceGroupManager#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The port number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/compute_region_instance_group_manager#port ComputeRegionInstanceGroupManager#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

