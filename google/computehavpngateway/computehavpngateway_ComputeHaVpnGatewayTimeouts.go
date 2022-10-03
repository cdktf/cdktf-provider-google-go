package computehavpngateway


type ComputeHaVpnGatewayTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_ha_vpn_gateway#create ComputeHaVpnGateway#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_ha_vpn_gateway#delete ComputeHaVpnGateway#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

