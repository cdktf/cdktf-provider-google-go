package computevpngateway


type ComputeVpnGatewayTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_vpn_gateway#create ComputeVpnGateway#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_vpn_gateway#delete ComputeVpnGateway#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

