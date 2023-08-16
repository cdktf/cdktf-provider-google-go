package dnspolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DnsPolicyConfig struct {
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
	// User assigned name for this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_policy#name DnsPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// alternative_name_server_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_policy#alternative_name_server_config DnsPolicy#alternative_name_server_config}
	AlternativeNameServerConfig *DnsPolicyAlternativeNameServerConfig `field:"optional" json:"alternativeNameServerConfig" yaml:"alternativeNameServerConfig"`
	// A textual description field. Defaults to 'Managed by Terraform'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_policy#description DnsPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Allows networks bound to this policy to receive DNS queries sent by VMs or applications over VPN connections.
	//
	// When enabled, a
	// virtual IP address will be allocated from each of the sub-networks
	// that are bound to this policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_policy#enable_inbound_forwarding DnsPolicy#enable_inbound_forwarding}
	EnableInboundForwarding interface{} `field:"optional" json:"enableInboundForwarding" yaml:"enableInboundForwarding"`
	// Controls whether logging is enabled for the networks bound to this policy. Defaults to no logging if not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_policy#enable_logging DnsPolicy#enable_logging}
	EnableLogging interface{} `field:"optional" json:"enableLogging" yaml:"enableLogging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_policy#id DnsPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_policy#networks DnsPolicy#networks}
	Networks interface{} `field:"optional" json:"networks" yaml:"networks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_policy#project DnsPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_policy#timeouts DnsPolicy#timeouts}
	Timeouts *DnsPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

