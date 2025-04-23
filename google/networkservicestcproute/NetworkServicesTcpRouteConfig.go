// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkservicestcproute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkServicesTcpRouteConfig struct {
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
	// Name of the TcpRoute resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/network_services_tcp_route#name NetworkServicesTcpRoute#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/network_services_tcp_route#rules NetworkServicesTcpRoute#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// A free-text description of the resource. Max length 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/network_services_tcp_route#description NetworkServicesTcpRoute#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Gateways defines a list of gateways this TcpRoute is attached to, as one of the routing rules to route the requests served by the gateway.
	//
	// Each gateway reference should match the pattern: projects/* /locations/global/gateways/<gateway_name>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/network_services_tcp_route#gateways NetworkServicesTcpRoute#gateways}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	Gateways *[]*string `field:"optional" json:"gateways" yaml:"gateways"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/network_services_tcp_route#id NetworkServicesTcpRoute#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of label tags associated with the TcpRoute resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/network_services_tcp_route#labels NetworkServicesTcpRoute#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Meshes defines a list of meshes this TcpRoute is attached to, as one of the routing rules to route the requests served by the mesh.
	//
	// Each mesh reference should match the pattern: projects/* /locations/global/meshes/<mesh_name>
	// The attached Mesh should be of a type SIDECAR
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/network_services_tcp_route#meshes NetworkServicesTcpRoute#meshes}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	Meshes *[]*string `field:"optional" json:"meshes" yaml:"meshes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/network_services_tcp_route#project NetworkServicesTcpRoute#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/network_services_tcp_route#timeouts NetworkServicesTcpRoute#timeouts}
	Timeouts *NetworkServicesTcpRouteTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

