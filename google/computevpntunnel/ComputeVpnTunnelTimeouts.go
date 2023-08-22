package computevpntunnel


type ComputeVpnTunnelTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_vpn_tunnel#create ComputeVpnTunnel#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_vpn_tunnel#delete ComputeVpnTunnel#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

