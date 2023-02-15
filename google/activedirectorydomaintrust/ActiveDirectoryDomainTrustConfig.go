package activedirectorydomaintrust

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ActiveDirectoryDomainTrustConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions,  https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/active_directory_domain_trust#domain ActiveDirectoryDomainTrust#domain}
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// The target DNS server IP addresses which can resolve the remote domain involved in the trust.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/active_directory_domain_trust#target_dns_ip_addresses ActiveDirectoryDomainTrust#target_dns_ip_addresses}
	TargetDnsIpAddresses *[]*string `field:"required" json:"targetDnsIpAddresses" yaml:"targetDnsIpAddresses"`
	// The fully qualified target domain name which will be in trust with the current domain.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/active_directory_domain_trust#target_domain_name ActiveDirectoryDomainTrust#target_domain_name}
	TargetDomainName *string `field:"required" json:"targetDomainName" yaml:"targetDomainName"`
	// The trust direction, which decides if the current domain is trusted, trusting, or both. Possible values: ["INBOUND", "OUTBOUND", "BIDIRECTIONAL"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/active_directory_domain_trust#trust_direction ActiveDirectoryDomainTrust#trust_direction}
	TrustDirection *string `field:"required" json:"trustDirection" yaml:"trustDirection"`
	// The trust secret used for the handshake with the target domain. This will not be stored.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/active_directory_domain_trust#trust_handshake_secret ActiveDirectoryDomainTrust#trust_handshake_secret}
	TrustHandshakeSecret *string `field:"required" json:"trustHandshakeSecret" yaml:"trustHandshakeSecret"`
	// The type of trust represented by the trust resource. Possible values: ["FOREST", "EXTERNAL"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/active_directory_domain_trust#trust_type ActiveDirectoryDomainTrust#trust_type}
	TrustType *string `field:"required" json:"trustType" yaml:"trustType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/active_directory_domain_trust#id ActiveDirectoryDomainTrust#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/active_directory_domain_trust#project ActiveDirectoryDomainTrust#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Whether the trusted side has forest/domain wide access or selective access to an approved set of resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/active_directory_domain_trust#selective_authentication ActiveDirectoryDomainTrust#selective_authentication}
	SelectiveAuthentication interface{} `field:"optional" json:"selectiveAuthentication" yaml:"selectiveAuthentication"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/active_directory_domain_trust#timeouts ActiveDirectoryDomainTrust#timeouts}
	Timeouts *ActiveDirectoryDomainTrustTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

