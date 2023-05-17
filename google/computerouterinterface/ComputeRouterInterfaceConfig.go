package computerouterinterface

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRouterInterfaceConfig struct {
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
	// A unique name for the interface, required by GCE. Changing this forces a new interface to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_router_interface#name ComputeRouterInterface#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the router this interface will be attached to.
	//
	// Changing this forces a new interface to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_router_interface#router ComputeRouterInterface#router}
	Router *string `field:"required" json:"router" yaml:"router"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_router_interface#id ComputeRouterInterface#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name or resource link to the VLAN interconnect for this interface.
	//
	// Changing this forces a new interface to be created. Only one of interconnect_attachment, subnetwork or vpn_tunnel can be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_router_interface#interconnect_attachment ComputeRouterInterface#interconnect_attachment}
	InterconnectAttachment *string `field:"optional" json:"interconnectAttachment" yaml:"interconnectAttachment"`
	// The IP address and range of the interface.
	//
	// The IP range must be in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_router_interface#ip_range ComputeRouterInterface#ip_range}
	IpRange *string `field:"optional" json:"ipRange" yaml:"ipRange"`
	// The regional private internal IP address that is used to establish BGP sessions to a VM instance acting as a third-party Router Appliance.
	//
	// Changing this forces a new interface to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_router_interface#private_ip_address ComputeRouterInterface#private_ip_address}
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// The ID of the project in which this interface's router belongs.
	//
	// If it is not provided, the provider project is used. Changing this forces a new interface to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_router_interface#project ComputeRouterInterface#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The name of the interface that is redundant to this interface.
	//
	// Changing this forces a new interface to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_router_interface#redundant_interface ComputeRouterInterface#redundant_interface}
	RedundantInterface *string `field:"optional" json:"redundantInterface" yaml:"redundantInterface"`
	// The region this interface's router sits in.
	//
	// If not specified, the project region will be used. Changing this forces a new interface to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_router_interface#region ComputeRouterInterface#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The URI of the subnetwork resource that this interface belongs to, which must be in the same region as the Cloud Router.
	//
	// Changing this forces a new interface to be created. Only one of subnetwork, interconnect_attachment or vpn_tunnel can be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_router_interface#subnetwork ComputeRouterInterface#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_router_interface#timeouts ComputeRouterInterface#timeouts}
	Timeouts *ComputeRouterInterfaceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The name or resource link to the VPN tunnel this interface will be linked to.
	//
	// Changing this forces a new interface to be created. Only one of vpn_tunnel, interconnect_attachment or subnetwork can be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_router_interface#vpn_tunnel ComputeRouterInterface#vpn_tunnel}
	VpnTunnel *string `field:"optional" json:"vpnTunnel" yaml:"vpnTunnel"`
}

