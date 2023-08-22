package dnsresponsepolicyrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DnsResponsePolicyRuleConfig struct {
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
	// The DNS name (wildcard or exact) to apply this rule to. Must be unique within the Response Policy Rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_response_policy_rule#dns_name DnsResponsePolicyRule#dns_name}
	DnsName *string `field:"required" json:"dnsName" yaml:"dnsName"`
	// Identifies the response policy addressed by this request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_response_policy_rule#response_policy DnsResponsePolicyRule#response_policy}
	ResponsePolicy *string `field:"required" json:"responsePolicy" yaml:"responsePolicy"`
	// An identifier for this rule. Must be unique with the ResponsePolicy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_response_policy_rule#rule_name DnsResponsePolicyRule#rule_name}
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_response_policy_rule#id DnsResponsePolicyRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// local_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_response_policy_rule#local_data DnsResponsePolicyRule#local_data}
	LocalData *DnsResponsePolicyRuleLocalData `field:"optional" json:"localData" yaml:"localData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_response_policy_rule#project DnsResponsePolicyRule#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_response_policy_rule#timeouts DnsResponsePolicyRule#timeouts}
	Timeouts *DnsResponsePolicyRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

