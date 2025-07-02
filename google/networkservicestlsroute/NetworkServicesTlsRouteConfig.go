// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkservicestlsroute

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkServicesTlsRouteConfig struct {
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
	// Name of the TlsRoute resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/network_services_tls_route#name NetworkServicesTlsRoute#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/network_services_tls_route#rules NetworkServicesTlsRoute#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
	// A free-text description of the resource. Max length 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/network_services_tls_route#description NetworkServicesTlsRoute#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Gateways defines a list of gateways this TlsRoute is attached to, as one of the routing rules to route the requests served by the gateway.
	//
	// Each gateway reference should match the pattern: projects/* /locations/global/gateways/<gateway_name>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/network_services_tls_route#gateways NetworkServicesTlsRoute#gateways}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	Gateways *[]*string `field:"optional" json:"gateways" yaml:"gateways"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/network_services_tls_route#id NetworkServicesTlsRoute#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Meshes defines a list of meshes this TlsRoute is attached to, as one of the routing rules to route the requests served by the mesh.
	//
	// Each mesh reference should match the pattern: projects/* /locations/global/meshes/<mesh_name>
	// The attached Mesh should be of a type SIDECAR
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/network_services_tls_route#meshes NetworkServicesTlsRoute#meshes}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	Meshes *[]*string `field:"optional" json:"meshes" yaml:"meshes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/network_services_tls_route#project NetworkServicesTlsRoute#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/network_services_tls_route#timeouts NetworkServicesTlsRoute#timeouts}
	Timeouts *NetworkServicesTlsRouteTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

