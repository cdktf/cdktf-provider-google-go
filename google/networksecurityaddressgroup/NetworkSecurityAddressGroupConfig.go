package networksecurityaddressgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkSecurityAddressGroupConfig struct {
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
	// Capacity of the Address Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/network_security_address_group#capacity NetworkSecurityAddressGroup#capacity}
	Capacity *float64 `field:"required" json:"capacity" yaml:"capacity"`
	// The location of the gateway security policy. The default value is 'global'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/network_security_address_group#location NetworkSecurityAddressGroup#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Name of the AddressGroup resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/network_security_address_group#name NetworkSecurityAddressGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of the Address Group. Possible values are "IPV4" or "IPV6". Possible values: ["IPV4", "IPV6"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/network_security_address_group#type NetworkSecurityAddressGroup#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Free-text description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/network_security_address_group#description NetworkSecurityAddressGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/network_security_address_group#id NetworkSecurityAddressGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// List of items.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/network_security_address_group#items NetworkSecurityAddressGroup#items}
	Items *[]*string `field:"optional" json:"items" yaml:"items"`
	// Set of label tags associated with the AddressGroup resource.
	//
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/network_security_address_group#labels NetworkSecurityAddressGroup#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of the parent this address group belongs to. Format: organizations/{organization_id} or projects/{project_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/network_security_address_group#parent NetworkSecurityAddressGroup#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/network_security_address_group#timeouts NetworkSecurityAddressGroup#timeouts}
	Timeouts *NetworkSecurityAddressGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

