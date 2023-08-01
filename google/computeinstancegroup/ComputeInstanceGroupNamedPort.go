package computeinstancegroup


type ComputeInstanceGroupNamedPort struct {
	// The name which the port will be mapped to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_instance_group#name ComputeInstanceGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The port number to map the name to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_instance_group#port ComputeInstanceGroup#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
}

