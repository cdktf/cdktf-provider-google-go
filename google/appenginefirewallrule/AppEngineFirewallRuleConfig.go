package appenginefirewallrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppEngineFirewallRuleConfig struct {
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
	// The action to take if this rule matches. Possible values: ["UNSPECIFIED_ACTION", "ALLOW", "DENY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/app_engine_firewall_rule#action AppEngineFirewallRule#action}
	Action *string `field:"required" json:"action" yaml:"action"`
	// IP address or range, defined using CIDR notation, of requests that this rule applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/app_engine_firewall_rule#source_range AppEngineFirewallRule#source_range}
	SourceRange *string `field:"required" json:"sourceRange" yaml:"sourceRange"`
	// An optional string description of this rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/app_engine_firewall_rule#description AppEngineFirewallRule#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/app_engine_firewall_rule#id AppEngineFirewallRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A positive integer that defines the order of rule evaluation. Rules with the lowest priority are evaluated first.
	//
	// A default rule at priority Int32.MaxValue matches all IPv4 and
	// IPv6 traffic when no previous rule matches. Only the action of
	// this rule can be modified by the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/app_engine_firewall_rule#priority AppEngineFirewallRule#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/app_engine_firewall_rule#project AppEngineFirewallRule#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/app_engine_firewall_rule#timeouts AppEngineFirewallRule#timeouts}
	Timeouts *AppEngineFirewallRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

