// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionnetworkfirewallpolicywithrules


type ComputeRegionNetworkFirewallPolicyWithRulesRule struct {
	// The Action to perform when the client connection triggers the rule. Can currently be either "allow", "deny", "apply_security_profile_group" or "goto_next".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_region_network_firewall_policy_with_rules#action ComputeRegionNetworkFirewallPolicyWithRules#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_region_network_firewall_policy_with_rules#match ComputeRegionNetworkFirewallPolicyWithRules#match}
	Match *ComputeRegionNetworkFirewallPolicyWithRulesRuleMatch `field:"required" json:"match" yaml:"match"`
	// An integer indicating the priority of a rule in the list.
	//
	// The priority must be a value
	// between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the
	// highest priority and 2147483647 is the lowest priority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_region_network_firewall_policy_with_rules#priority ComputeRegionNetworkFirewallPolicyWithRules#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// A description of the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_region_network_firewall_policy_with_rules#description ComputeRegionNetworkFirewallPolicyWithRules#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The direction in which this rule applies. If unspecified an INGRESS rule is created. Possible values: ["INGRESS", "EGRESS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_region_network_firewall_policy_with_rules#direction ComputeRegionNetworkFirewallPolicyWithRules#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Denotes whether the firewall policy rule is disabled.
	//
	// When set to true,
	// the firewall policy rule is not enforced and traffic behaves as if it did
	// not exist. If this is unspecified, the firewall policy rule will be
	// enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_region_network_firewall_policy_with_rules#disabled ComputeRegionNetworkFirewallPolicyWithRules#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Denotes whether to enable logging for a particular rule.
	//
	// If logging is enabled, logs will be exported to the
	// configured export destination in Stackdriver.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_region_network_firewall_policy_with_rules#enable_logging ComputeRegionNetworkFirewallPolicyWithRules#enable_logging}
	EnableLogging interface{} `field:"optional" json:"enableLogging" yaml:"enableLogging"`
	// An optional name for the rule. This field is not a unique identifier and can be updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_region_network_firewall_policy_with_rules#rule_name ComputeRegionNetworkFirewallPolicyWithRules#rule_name}
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// A fully-qualified URL of a SecurityProfile resource instance. Example: https://networksecurity.googleapis.com/v1/projects/{project}/locations/{location}/securityProfileGroups/my-security-profile-group Must be specified if action is 'apply_security_profile_group'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_region_network_firewall_policy_with_rules#security_profile_group ComputeRegionNetworkFirewallPolicyWithRules#security_profile_group}
	SecurityProfileGroup *string `field:"optional" json:"securityProfileGroup" yaml:"securityProfileGroup"`
	// target_secure_tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_region_network_firewall_policy_with_rules#target_secure_tag ComputeRegionNetworkFirewallPolicyWithRules#target_secure_tag}
	TargetSecureTag interface{} `field:"optional" json:"targetSecureTag" yaml:"targetSecureTag"`
	// A list of service accounts indicating the sets of instances that are applied with this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_region_network_firewall_policy_with_rules#target_service_accounts ComputeRegionNetworkFirewallPolicyWithRules#target_service_accounts}
	TargetServiceAccounts *[]*string `field:"optional" json:"targetServiceAccounts" yaml:"targetServiceAccounts"`
	// Boolean flag indicating if the traffic should be TLS decrypted.
	//
	// It can be set only if action = 'apply_security_profile_group' and cannot be set for other actions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/compute_region_network_firewall_policy_with_rules#tls_inspect ComputeRegionNetworkFirewallPolicyWithRules#tls_inspect}
	TlsInspect interface{} `field:"optional" json:"tlsInspect" yaml:"tlsInspect"`
}

