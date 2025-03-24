// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computehavpngateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeHaVpnGatewayConfig struct {
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
	// Name of the resource.
	//
	// Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_ha_vpn_gateway#name ComputeHaVpnGateway#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The network this VPN gateway is accepting traffic for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_ha_vpn_gateway#network ComputeHaVpnGateway#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_ha_vpn_gateway#description ComputeHaVpnGateway#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The IP family of the gateway IPs for the HA-VPN gateway interfaces.
	//
	// If not specified, IPV4 will be used. Default value: "IPV4" Possible values: ["IPV4", "IPV6"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_ha_vpn_gateway#gateway_ip_version ComputeHaVpnGateway#gateway_ip_version}
	GatewayIpVersion *string `field:"optional" json:"gatewayIpVersion" yaml:"gatewayIpVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_ha_vpn_gateway#id ComputeHaVpnGateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels for this resource.
	//
	// These can only be added or modified by the setLabels method.
	// Each label key/value pair must comply with RFC1035. Label values may be empty.
	//
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_ha_vpn_gateway#labels ComputeHaVpnGateway#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_ha_vpn_gateway#project ComputeHaVpnGateway#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region this gateway should sit in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_ha_vpn_gateway#region ComputeHaVpnGateway#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The stack type for this VPN gateway to identify the IP protocols that are enabled.
	//
	// If not specified, IPV4_ONLY will be used. Default value: "IPV4_ONLY" Possible values: ["IPV4_ONLY", "IPV4_IPV6", "IPV6_ONLY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_ha_vpn_gateway#stack_type ComputeHaVpnGateway#stack_type}
	StackType *string `field:"optional" json:"stackType" yaml:"stackType"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_ha_vpn_gateway#timeouts ComputeHaVpnGateway#timeouts}
	Timeouts *ComputeHaVpnGatewayTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// vpn_interfaces block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_ha_vpn_gateway#vpn_interfaces ComputeHaVpnGateway#vpn_interfaces}
	VpnInterfaces interface{} `field:"optional" json:"vpnInterfaces" yaml:"vpnInterfaces"`
}

