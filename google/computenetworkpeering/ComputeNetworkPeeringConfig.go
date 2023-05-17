package computenetworkpeering

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeNetworkPeeringConfig struct {
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
	// Name of the peering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_peering#name ComputeNetworkPeering#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The primary network of the peering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_peering#network ComputeNetworkPeering#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// The peer network in the peering. The peer network may belong to a different project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_peering#peer_network ComputeNetworkPeering#peer_network}
	PeerNetwork *string `field:"required" json:"peerNetwork" yaml:"peerNetwork"`
	// Whether to export the custom routes to the peer network. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_peering#export_custom_routes ComputeNetworkPeering#export_custom_routes}
	ExportCustomRoutes interface{} `field:"optional" json:"exportCustomRoutes" yaml:"exportCustomRoutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_peering#export_subnet_routes_with_public_ip ComputeNetworkPeering#export_subnet_routes_with_public_ip}.
	ExportSubnetRoutesWithPublicIp interface{} `field:"optional" json:"exportSubnetRoutesWithPublicIp" yaml:"exportSubnetRoutesWithPublicIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_peering#id ComputeNetworkPeering#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether to export the custom routes from the peer network. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_peering#import_custom_routes ComputeNetworkPeering#import_custom_routes}
	ImportCustomRoutes interface{} `field:"optional" json:"importCustomRoutes" yaml:"importCustomRoutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_peering#import_subnet_routes_with_public_ip ComputeNetworkPeering#import_subnet_routes_with_public_ip}.
	ImportSubnetRoutesWithPublicIp interface{} `field:"optional" json:"importSubnetRoutesWithPublicIp" yaml:"importSubnetRoutesWithPublicIp"`
	// Which IP version(s) of traffic and routes are allowed to be imported or exported between peer networks.
	//
	// The default value is IPV4_ONLY. Possible values: ["IPV4_ONLY", "IPV4_IPV6"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_peering#stack_type ComputeNetworkPeering#stack_type}
	StackType *string `field:"optional" json:"stackType" yaml:"stackType"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/compute_network_peering#timeouts ComputeNetworkPeering#timeouts}
	Timeouts *ComputeNetworkPeeringTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

