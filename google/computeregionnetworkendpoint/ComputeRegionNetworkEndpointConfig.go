// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionnetworkendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionNetworkEndpointConfig struct {
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
	// Port number of network endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/compute_region_network_endpoint#port ComputeRegionNetworkEndpoint#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The network endpoint group this endpoint is part of.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/compute_region_network_endpoint#region_network_endpoint_group ComputeRegionNetworkEndpoint#region_network_endpoint_group}
	RegionNetworkEndpointGroup *string `field:"required" json:"regionNetworkEndpointGroup" yaml:"regionNetworkEndpointGroup"`
	// Fully qualified domain name of network endpoint.
	//
	// This can only be specified when network_endpoint_type of the NEG is INTERNET_FQDN_PORT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/compute_region_network_endpoint#fqdn ComputeRegionNetworkEndpoint#fqdn}
	Fqdn *string `field:"optional" json:"fqdn" yaml:"fqdn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/compute_region_network_endpoint#id ComputeRegionNetworkEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// IPv4 address external endpoint.
	//
	// This can only be specified when network_endpoint_type of the NEG is INTERNET_IP_PORT.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/compute_region_network_endpoint#ip_address ComputeRegionNetworkEndpoint#ip_address}
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/compute_region_network_endpoint#project ComputeRegionNetworkEndpoint#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Region where the containing network endpoint group is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/compute_region_network_endpoint#region ComputeRegionNetworkEndpoint#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/compute_region_network_endpoint#timeouts ComputeRegionNetworkEndpoint#timeouts}
	Timeouts *ComputeRegionNetworkEndpointTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

