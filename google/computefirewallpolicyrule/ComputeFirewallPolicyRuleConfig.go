// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computefirewallpolicyrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeFirewallPolicyRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The Action to perform when the client connection triggers the rule. Valid actions are "allow", "deny", "goto_next" and "apply_security_profile_group".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/compute_firewall_policy_rule#action ComputeFirewallPolicyRule#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// The direction in which this rule applies. Possible values: INGRESS, EGRESS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/compute_firewall_policy_rule#direction ComputeFirewallPolicyRule#direction}
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// The firewall policy of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/compute_firewall_policy_rule#firewall_policy ComputeFirewallPolicyRule#firewall_policy}
	FirewallPolicy *string `field:"required" json:"firewallPolicy" yaml:"firewallPolicy"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/compute_firewall_policy_rule#match ComputeFirewallPolicyRule#match}
	Match *ComputeFirewallPolicyRuleMatch `field:"required" json:"match" yaml:"match"`
	// An integer indicating the priority of a rule in the list.
	//
	// The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/compute_firewall_policy_rule#priority ComputeFirewallPolicyRule#priority}
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// An optional description for this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/compute_firewall_policy_rule#description ComputeFirewallPolicyRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Denotes whether the firewall policy rule is disabled.
	//
	// When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/compute_firewall_policy_rule#disabled ComputeFirewallPolicyRule#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Denotes whether to enable logging for a particular rule.
	//
	// If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "goto_next" rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/compute_firewall_policy_rule#enable_logging ComputeFirewallPolicyRule#enable_logging}
	EnableLogging interface{} `field:"optional" json:"enableLogging" yaml:"enableLogging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/compute_firewall_policy_rule#id ComputeFirewallPolicyRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A fully-qualified URL of a SecurityProfileGroup resource.
	//
	// Example: https://networksecurity.googleapis.com/v1/organizations/{organizationId}/locations/global/securityProfileGroups/my-security-profile-group. It must be specified if action = 'apply_security_profile_group' and cannot be specified for other actions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/compute_firewall_policy_rule#security_profile_group ComputeFirewallPolicyRule#security_profile_group}
	SecurityProfileGroup *string `field:"optional" json:"securityProfileGroup" yaml:"securityProfileGroup"`
	// A list of network resource URLs to which this rule applies.
	//
	// This field allows you to control which network's VMs get this rule. If this field is left blank, all VMs within the organization will receive the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/compute_firewall_policy_rule#target_resources ComputeFirewallPolicyRule#target_resources}
	TargetResources *[]*string `field:"optional" json:"targetResources" yaml:"targetResources"`
	// A list of service accounts indicating the sets of instances that are applied with this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/compute_firewall_policy_rule#target_service_accounts ComputeFirewallPolicyRule#target_service_accounts}
	TargetServiceAccounts *[]*string `field:"optional" json:"targetServiceAccounts" yaml:"targetServiceAccounts"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/compute_firewall_policy_rule#timeouts ComputeFirewallPolicyRule#timeouts}
	Timeouts *ComputeFirewallPolicyRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Boolean flag indicating if the traffic should be TLS decrypted.
	//
	// It can be set only if action = 'apply_security_profile_group' and cannot be set for other actions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.6.0/docs/resources/compute_firewall_policy_rule#tls_inspect ComputeFirewallPolicyRule#tls_inspect}
	TlsInspect interface{} `field:"optional" json:"tlsInspect" yaml:"tlsInspect"`
}

