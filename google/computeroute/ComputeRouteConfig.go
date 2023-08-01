package computeroute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRouteConfig struct {
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
	// The destination range of outgoing packets that this route applies to. Only IPv4 is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_route#dest_range ComputeRoute#dest_range}
	DestRange *string `field:"required" json:"destRange" yaml:"destRange"`
	// Name of the resource.
	//
	// Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_route#name ComputeRoute#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The network that this route applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_route#network ComputeRoute#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// An optional description of this resource. Provide this property when you create the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_route#description ComputeRoute#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_route#id ComputeRoute#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// URL to a gateway that should handle matching packets.
	//
	// Currently, you can only specify the internet gateway, using a full or
	// partial valid URL:
	// 'https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway'
	// 'projects/project/global/gateways/default-internet-gateway'
	// 'global/gateways/default-internet-gateway'
	// The string 'default-internet-gateway'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_route#next_hop_gateway ComputeRoute#next_hop_gateway}
	NextHopGateway *string `field:"optional" json:"nextHopGateway" yaml:"nextHopGateway"`
	// The IP address or URL to a forwarding rule of type loadBalancingScheme=INTERNAL that should handle matching packets.
	//
	// With the GA provider you can only specify the forwarding
	// rule as a partial or full URL. For example, the following
	// are all valid values:
	// 10.128.0.56
	// https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
	// regions/region/forwardingRules/forwardingRule
	//
	// When the beta provider, you can also specify the IP address
	// of a forwarding rule from the same VPC or any peered VPC.
	//
	// Note that this can only be used when the destinationRange is
	// a public (non-RFC 1918) IP CIDR range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_route#next_hop_ilb ComputeRoute#next_hop_ilb}
	NextHopIlb *string `field:"optional" json:"nextHopIlb" yaml:"nextHopIlb"`
	// URL to an instance that should handle matching packets.
	//
	// You can specify this as a full or partial URL. For example:
	// 'https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance'
	// 'projects/project/zones/zone/instances/instance'
	// 'zones/zone/instances/instance'
	// Just the instance name, with the zone in 'next_hop_instance_zone'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_route#next_hop_instance ComputeRoute#next_hop_instance}
	NextHopInstance *string `field:"optional" json:"nextHopInstance" yaml:"nextHopInstance"`
	// The zone of the instance specified in next_hop_instance. Omit if next_hop_instance is specified as a URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_route#next_hop_instance_zone ComputeRoute#next_hop_instance_zone}
	NextHopInstanceZone *string `field:"optional" json:"nextHopInstanceZone" yaml:"nextHopInstanceZone"`
	// Network IP address of an instance that should handle matching packets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_route#next_hop_ip ComputeRoute#next_hop_ip}
	NextHopIp *string `field:"optional" json:"nextHopIp" yaml:"nextHopIp"`
	// URL to a VpnTunnel that should handle matching packets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_route#next_hop_vpn_tunnel ComputeRoute#next_hop_vpn_tunnel}
	NextHopVpnTunnel *string `field:"optional" json:"nextHopVpnTunnel" yaml:"nextHopVpnTunnel"`
	// The priority of this route.
	//
	// Priority is used to break ties in cases
	// where there is more than one matching route of equal prefix length.
	//
	// In the case of two routes with equal prefix length, the one with the
	// lowest-numbered priority value wins.
	//
	// Default value is 1000. Valid range is 0 through 65535.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_route#priority ComputeRoute#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_route#project ComputeRoute#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// A list of instance tags to which this route applies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_route#tags ComputeRoute#tags}
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/compute_route#timeouts ComputeRoute#timeouts}
	Timeouts *ComputeRouteTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

