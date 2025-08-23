// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmwareenginenetworkpeering

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VmwareengineNetworkPeeringConfig struct {
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
	// The ID of the Network Peering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/vmwareengine_network_peering#name VmwareengineNetworkPeering#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The relative resource name of the network to peer with a standard VMware Engine network.
	//
	// The provided network can be a consumer VPC network or another standard VMware Engine network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/vmwareengine_network_peering#peer_network VmwareengineNetworkPeering#peer_network}
	PeerNetwork *string `field:"required" json:"peerNetwork" yaml:"peerNetwork"`
	// The type of the network to peer with the VMware Engine network.
	//
	// Possible values: ["STANDARD", "VMWARE_ENGINE_NETWORK", "PRIVATE_SERVICES_ACCESS", "NETAPP_CLOUD_VOLUMES", "THIRD_PARTY_SERVICE", "DELL_POWERSCALE", "GOOGLE_CLOUD_NETAPP_VOLUMES"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/vmwareengine_network_peering#peer_network_type VmwareengineNetworkPeering#peer_network_type}
	PeerNetworkType *string `field:"required" json:"peerNetworkType" yaml:"peerNetworkType"`
	// The relative resource name of the VMware Engine network.
	//
	// Specify the name in the following form:
	// projects/{project}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId} where {project}
	// can either be a project number or a project ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/vmwareengine_network_peering#vmware_engine_network VmwareengineNetworkPeering#vmware_engine_network}
	VmwareEngineNetwork *string `field:"required" json:"vmwareEngineNetwork" yaml:"vmwareEngineNetwork"`
	// User-provided description for this network peering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/vmwareengine_network_peering#description VmwareengineNetworkPeering#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// True if custom routes are exported to the peered network; false otherwise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/vmwareengine_network_peering#export_custom_routes VmwareengineNetworkPeering#export_custom_routes}
	ExportCustomRoutes interface{} `field:"optional" json:"exportCustomRoutes" yaml:"exportCustomRoutes"`
	// True if all subnet routes with a public IP address range are exported; false otherwise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/vmwareengine_network_peering#export_custom_routes_with_public_ip VmwareengineNetworkPeering#export_custom_routes_with_public_ip}
	ExportCustomRoutesWithPublicIp interface{} `field:"optional" json:"exportCustomRoutesWithPublicIp" yaml:"exportCustomRoutesWithPublicIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/vmwareengine_network_peering#id VmwareengineNetworkPeering#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// True if custom routes are imported from the peered network; false otherwise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/vmwareengine_network_peering#import_custom_routes VmwareengineNetworkPeering#import_custom_routes}
	ImportCustomRoutes interface{} `field:"optional" json:"importCustomRoutes" yaml:"importCustomRoutes"`
	// True if custom routes are imported from the peered network; false otherwise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/vmwareengine_network_peering#import_custom_routes_with_public_ip VmwareengineNetworkPeering#import_custom_routes_with_public_ip}
	ImportCustomRoutesWithPublicIp interface{} `field:"optional" json:"importCustomRoutesWithPublicIp" yaml:"importCustomRoutesWithPublicIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/vmwareengine_network_peering#project VmwareengineNetworkPeering#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/vmwareengine_network_peering#timeouts VmwareengineNetworkPeering#timeouts}
	Timeouts *VmwareengineNetworkPeeringTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

