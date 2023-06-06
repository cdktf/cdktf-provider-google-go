package computeregionnetworkfirewallpolicyrule


type ComputeRegionNetworkFirewallPolicyRuleTargetSecureTags struct {
	// Name of the secure tag, created with TagManager's TagValue API. @pattern tagValues/[0-9]+.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.68.0/docs/resources/compute_region_network_firewall_policy_rule#name ComputeRegionNetworkFirewallPolicyRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

