package computeregionnetworkfirewallpolicy


type ComputeRegionNetworkFirewallPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_network_firewall_policy#create ComputeRegionNetworkFirewallPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_network_firewall_policy#delete ComputeRegionNetworkFirewallPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_region_network_firewall_policy#update ComputeRegionNetworkFirewallPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
