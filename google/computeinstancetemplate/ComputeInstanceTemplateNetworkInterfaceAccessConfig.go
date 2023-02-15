package computeinstancetemplate


type ComputeInstanceTemplateNetworkInterfaceAccessConfig struct {
	// The IP address that will be 1:1 mapped to the instance's network ip.
	//
	// If not given, one will be generated.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance_template#nat_ip ComputeInstanceTemplate#nat_ip}
	NatIp *string `field:"optional" json:"natIp" yaml:"natIp"`
	// The networking tier used for configuring this instance template.
	//
	// This field can take the following values: PREMIUM, STANDARD, FIXED_STANDARD. If this field is not specified, it is assumed to be PREMIUM.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance_template#network_tier ComputeInstanceTemplate#network_tier}
	NetworkTier *string `field:"optional" json:"networkTier" yaml:"networkTier"`
}

