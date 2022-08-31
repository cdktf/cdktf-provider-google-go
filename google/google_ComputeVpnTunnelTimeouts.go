// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ComputeVpnTunnelTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_vpn_tunnel#create ComputeVpnTunnel#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_vpn_tunnel#delete ComputeVpnTunnel#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

