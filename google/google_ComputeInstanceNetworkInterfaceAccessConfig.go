// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ComputeInstanceNetworkInterfaceAccessConfig struct {
	// The IP address that is be 1:1 mapped to the instance's network ip.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance#nat_ip ComputeInstance#nat_ip}
	NatIp *string `field:"optional" json:"natIp" yaml:"natIp"`
	// The networking tier used for configuring this instance. One of PREMIUM or STANDARD.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance#network_tier ComputeInstance#network_tier}
	NetworkTier *string `field:"optional" json:"networkTier" yaml:"networkTier"`
	// The DNS domain name for the public PTR record.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_instance#public_ptr_domain_name ComputeInstance#public_ptr_domain_name}
	PublicPtrDomainName *string `field:"optional" json:"publicPtrDomainName" yaml:"publicPtrDomainName"`
}

