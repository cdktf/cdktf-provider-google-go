package computenetworkfirewallpolicyrule


type ComputeNetworkFirewallPolicyRuleTargetSecureTags struct {
	// Name of the secure tag, created with TagManager's TagValue API. @pattern tagValues/[0-9]+.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/compute_network_firewall_policy_rule#name ComputeNetworkFirewallPolicyRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

