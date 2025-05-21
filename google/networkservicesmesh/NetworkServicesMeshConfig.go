// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkservicesmesh

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkServicesMeshConfig struct {
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
	// Short name of the Mesh resource to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/network_services_mesh#name NetworkServicesMesh#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A free-text description of the resource. Max length 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/network_services_mesh#description NetworkServicesMesh#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/network_services_mesh#id NetworkServicesMesh#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional.
	//
	// If set to a valid TCP port (1-65535), instructs the SIDECAR proxy to listen on the
	// specified port of localhost (127.0.0.1) address. The SIDECAR proxy will expect all traffic to
	// be redirected to this port regardless of its actual ip:port destination. If unset, a port
	// '15001' is used as the interception port. This will is applicable only for sidecar proxy
	// deployments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/network_services_mesh#interception_port NetworkServicesMesh#interception_port}
	InterceptionPort *float64 `field:"optional" json:"interceptionPort" yaml:"interceptionPort"`
	// Set of label tags associated with the Mesh resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/network_services_mesh#labels NetworkServicesMesh#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Location (region) of the Mesh resource to be created.
	//
	// Only the value 'global' is currently allowed; defaults to 'global' if omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/network_services_mesh#location NetworkServicesMesh#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/network_services_mesh#project NetworkServicesMesh#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.0/docs/resources/network_services_mesh#timeouts NetworkServicesMesh#timeouts}
	Timeouts *NetworkServicesMeshTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

