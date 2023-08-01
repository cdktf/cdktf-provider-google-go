package computeresourcepolicy


type ComputeResourcePolicyGroupPlacementPolicy struct {
	// The number of availability domains instances will be spread across.
	//
	// If two instances are in different
	// availability domain, they will not be put in the same low latency network
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_resource_policy#availability_domain_count ComputeResourcePolicy#availability_domain_count}
	AvailabilityDomainCount *float64 `field:"optional" json:"availabilityDomainCount" yaml:"availabilityDomainCount"`
	// Collocation specifies whether to place VMs inside the same availability domain on the same low-latency network.
	//
	// Specify 'COLLOCATED' to enable collocation. Can only be specified with 'vm_count'. If compute instances are created
	// with a COLLOCATED policy, then exactly 'vm_count' instances must be created at the same time with the resource policy
	// attached. Possible values: ["COLLOCATED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_resource_policy#collocation ComputeResourcePolicy#collocation}
	Collocation *string `field:"optional" json:"collocation" yaml:"collocation"`
	// Number of VMs in this placement group.
	//
	// Google does not recommend that you use this field
	// unless you use a compact policy and you want your policy to work only if it contains this
	// exact number of VMs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_resource_policy#vm_count ComputeResourcePolicy#vm_count}
	VmCount *float64 `field:"optional" json:"vmCount" yaml:"vmCount"`
}

