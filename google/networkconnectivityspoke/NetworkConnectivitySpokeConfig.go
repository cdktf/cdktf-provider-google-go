package networkconnectivityspoke

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkConnectivitySpokeConfig struct {
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
	// Immutable. The URI of the hub that this spoke is attached to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_connectivity_spoke#hub NetworkConnectivitySpoke#hub}
	Hub *string `field:"required" json:"hub" yaml:"hub"`
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_connectivity_spoke#location NetworkConnectivitySpoke#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Immutable. The name of the spoke. Spoke names must be unique.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_connectivity_spoke#name NetworkConnectivitySpoke#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// An optional description of the spoke.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_connectivity_spoke#description NetworkConnectivitySpoke#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_connectivity_spoke#id NetworkConnectivitySpoke#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_connectivity_spoke#labels NetworkConnectivitySpoke#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// linked_interconnect_attachments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_connectivity_spoke#linked_interconnect_attachments NetworkConnectivitySpoke#linked_interconnect_attachments}
	LinkedInterconnectAttachments *NetworkConnectivitySpokeLinkedInterconnectAttachments `field:"optional" json:"linkedInterconnectAttachments" yaml:"linkedInterconnectAttachments"`
	// linked_router_appliance_instances block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_connectivity_spoke#linked_router_appliance_instances NetworkConnectivitySpoke#linked_router_appliance_instances}
	LinkedRouterApplianceInstances *NetworkConnectivitySpokeLinkedRouterApplianceInstances `field:"optional" json:"linkedRouterApplianceInstances" yaml:"linkedRouterApplianceInstances"`
	// linked_vpn_tunnels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_connectivity_spoke#linked_vpn_tunnels NetworkConnectivitySpoke#linked_vpn_tunnels}
	LinkedVpnTunnels *NetworkConnectivitySpokeLinkedVpnTunnels `field:"optional" json:"linkedVpnTunnels" yaml:"linkedVpnTunnels"`
	// The project for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_connectivity_spoke#project NetworkConnectivitySpoke#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/network_connectivity_spoke#timeouts NetworkConnectivitySpoke#timeouts}
	Timeouts *NetworkConnectivitySpokeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

