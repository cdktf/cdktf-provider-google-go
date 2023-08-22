package dnsmanagedzone

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DnsManagedZoneConfig struct {
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
	// The DNS name of this managed zone, for instance "example.com.".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_managed_zone#dns_name DnsManagedZone#dns_name}
	DnsName *string `field:"required" json:"dnsName" yaml:"dnsName"`
	// User assigned name for this resource. Must be unique within the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_managed_zone#name DnsManagedZone#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// cloud_logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_managed_zone#cloud_logging_config DnsManagedZone#cloud_logging_config}
	CloudLoggingConfig *DnsManagedZoneCloudLoggingConfig `field:"optional" json:"cloudLoggingConfig" yaml:"cloudLoggingConfig"`
	// A textual description field. Defaults to 'Managed by Terraform'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_managed_zone#description DnsManagedZone#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// dnssec_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_managed_zone#dnssec_config DnsManagedZone#dnssec_config}
	DnssecConfig *DnsManagedZoneDnssecConfig `field:"optional" json:"dnssecConfig" yaml:"dnssecConfig"`
	// Set this true to delete all records in the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_managed_zone#force_destroy DnsManagedZone#force_destroy}
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// forwarding_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_managed_zone#forwarding_config DnsManagedZone#forwarding_config}
	ForwardingConfig *DnsManagedZoneForwardingConfig `field:"optional" json:"forwardingConfig" yaml:"forwardingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_managed_zone#id DnsManagedZone#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A set of key/value label pairs to assign to this ManagedZone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_managed_zone#labels DnsManagedZone#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// peering_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_managed_zone#peering_config DnsManagedZone#peering_config}
	PeeringConfig *DnsManagedZonePeeringConfig `field:"optional" json:"peeringConfig" yaml:"peeringConfig"`
	// private_visibility_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_managed_zone#private_visibility_config DnsManagedZone#private_visibility_config}
	PrivateVisibilityConfig *DnsManagedZonePrivateVisibilityConfig `field:"optional" json:"privateVisibilityConfig" yaml:"privateVisibilityConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_managed_zone#project DnsManagedZone#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_managed_zone#timeouts DnsManagedZone#timeouts}
	Timeouts *DnsManagedZoneTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The zone's visibility: public zones are exposed to the Internet, while private zones are visible only to Virtual Private Cloud resources.
	//
	// Default value: "public" Possible values: ["private", "public"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dns_managed_zone#visibility DnsManagedZone#visibility}
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

