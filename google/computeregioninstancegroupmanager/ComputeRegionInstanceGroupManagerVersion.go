package computeregioninstancegroupmanager


type ComputeRegionInstanceGroupManagerVersion struct {
	// The full URL to an instance template from which all new instances of this version will be created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_instance_group_manager#instance_template ComputeRegionInstanceGroupManager#instance_template}
	InstanceTemplate *string `field:"required" json:"instanceTemplate" yaml:"instanceTemplate"`
	// Version name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_instance_group_manager#name ComputeRegionInstanceGroupManager#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// target_size block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_instance_group_manager#target_size ComputeRegionInstanceGroupManager#target_size}
	TargetSize *ComputeRegionInstanceGroupManagerVersionTargetSize `field:"optional" json:"targetSize" yaml:"targetSize"`
}

