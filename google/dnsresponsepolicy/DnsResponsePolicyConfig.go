package dnsresponsepolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DnsResponsePolicyConfig struct {
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
	// The user assigned name for this Response Policy, such as 'myresponsepolicy'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_response_policy#response_policy_name DnsResponsePolicy#response_policy_name}
	ResponsePolicyName *string `field:"required" json:"responsePolicyName" yaml:"responsePolicyName"`
	// The description of the response policy, such as 'My new response policy'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_response_policy#description DnsResponsePolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// gke_clusters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_response_policy#gke_clusters DnsResponsePolicy#gke_clusters}
	GkeClusters interface{} `field:"optional" json:"gkeClusters" yaml:"gkeClusters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_response_policy#id DnsResponsePolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// networks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_response_policy#networks DnsResponsePolicy#networks}
	Networks interface{} `field:"optional" json:"networks" yaml:"networks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_response_policy#project DnsResponsePolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dns_response_policy#timeouts DnsResponsePolicy#timeouts}
	Timeouts *DnsResponsePolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

