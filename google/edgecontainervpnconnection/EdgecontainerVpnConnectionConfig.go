// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package edgecontainervpnconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EdgecontainerVpnConnectionConfig struct {
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
	// The canonical Cluster name to connect to. It is in the form of projects/{project}/locations/{location}/clusters/{cluster}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/edgecontainer_vpn_connection#cluster EdgecontainerVpnConnection#cluster}
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// Google Cloud Platform location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/edgecontainer_vpn_connection#location EdgecontainerVpnConnection#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The resource name of VPN connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/edgecontainer_vpn_connection#name EdgecontainerVpnConnection#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Whether this VPN connection has HA enabled on cluster side.
	//
	// If enabled, when creating VPN connection we will attempt to use 2 ANG floating IPs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/edgecontainer_vpn_connection#enable_high_availability EdgecontainerVpnConnection#enable_high_availability}
	EnableHighAvailability interface{} `field:"optional" json:"enableHighAvailability" yaml:"enableHighAvailability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/edgecontainer_vpn_connection#id EdgecontainerVpnConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels associated with this resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/edgecontainer_vpn_connection#labels EdgecontainerVpnConnection#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// NAT gateway IP, or WAN IP address.
	//
	// If a customer has multiple NAT IPs, the customer needs to configure NAT such that only one external IP maps to the GMEC Anthos cluster.
	// This is empty if NAT is not used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/edgecontainer_vpn_connection#nat_gateway_ip EdgecontainerVpnConnection#nat_gateway_ip}
	NatGatewayIp *string `field:"optional" json:"natGatewayIp" yaml:"natGatewayIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/edgecontainer_vpn_connection#project EdgecontainerVpnConnection#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The VPN connection Cloud Router name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/edgecontainer_vpn_connection#router EdgecontainerVpnConnection#router}
	Router *string `field:"optional" json:"router" yaml:"router"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/edgecontainer_vpn_connection#timeouts EdgecontainerVpnConnection#timeouts}
	Timeouts *EdgecontainerVpnConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The network ID of VPC to connect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/edgecontainer_vpn_connection#vpc EdgecontainerVpnConnection#vpc}
	Vpc *string `field:"optional" json:"vpc" yaml:"vpc"`
	// vpc_project block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.11.1/docs/resources/edgecontainer_vpn_connection#vpc_project EdgecontainerVpnConnection#vpc_project}
	VpcProject *EdgecontainerVpnConnectionVpcProject `field:"optional" json:"vpcProject" yaml:"vpcProject"`
}

