package computeinstancegroupmanager


type ComputeInstanceGroupManagerVersion struct {
	// The full URL to an instance template from which all new instances of this version will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_instance_group_manager#instance_template ComputeInstanceGroupManager#instance_template}
	InstanceTemplate *string `field:"required" json:"instanceTemplate" yaml:"instanceTemplate"`
	// Version name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_instance_group_manager#name ComputeInstanceGroupManager#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// target_size block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_instance_group_manager#target_size ComputeInstanceGroupManager#target_size}
	TargetSize *ComputeInstanceGroupManagerVersionTargetSize `field:"optional" json:"targetSize" yaml:"targetSize"`
}

