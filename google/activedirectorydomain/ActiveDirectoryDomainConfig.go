package activedirectorydomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ActiveDirectoryDomainConfig struct {
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
	// The fully qualified domain name. e.g. mydomain.myorganization.com, with the restrictions, https://cloud.google.com/managed-microsoft-ad/reference/rest/v1/projects.locations.global.domains.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/active_directory_domain#domain_name ActiveDirectoryDomain#domain_name}
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Locations where domain needs to be provisioned.
	//
	// [regions][compute/docs/regions-zones/]
	// e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/active_directory_domain#locations ActiveDirectoryDomain#locations}
	Locations *[]*string `field:"required" json:"locations" yaml:"locations"`
	// The CIDR range of internal addresses that are reserved for this domain.
	//
	// Reserved networks must be /24 or larger.
	// Ranges must be unique and non-overlapping with existing subnets in authorizedNetworks
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/active_directory_domain#reserved_ip_range ActiveDirectoryDomain#reserved_ip_range}
	ReservedIpRange *string `field:"required" json:"reservedIpRange" yaml:"reservedIpRange"`
	// The name of delegated administrator account used to perform Active Directory operations. If not specified, setupadmin will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/active_directory_domain#admin ActiveDirectoryDomain#admin}
	Admin *string `field:"optional" json:"admin" yaml:"admin"`
	// The full names of the Google Compute Engine networks the domain instance is connected to.
	//
	// The domain is only available on networks listed in authorizedNetworks.
	// If CIDR subnets overlap between networks, domain creation will fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/active_directory_domain#authorized_networks ActiveDirectoryDomain#authorized_networks}
	AuthorizedNetworks *[]*string `field:"optional" json:"authorizedNetworks" yaml:"authorizedNetworks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/active_directory_domain#id ActiveDirectoryDomain#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource labels that can contain user-provided metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/active_directory_domain#labels ActiveDirectoryDomain#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/active_directory_domain#project ActiveDirectoryDomain#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/active_directory_domain#timeouts ActiveDirectoryDomain#timeouts}
	Timeouts *ActiveDirectoryDomainTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

