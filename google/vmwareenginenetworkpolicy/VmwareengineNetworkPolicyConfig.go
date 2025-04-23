// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vmwareenginenetworkpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VmwareengineNetworkPolicyConfig struct {
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
	// IP address range in CIDR notation used to create internet access and external IP access.
	//
	// An RFC 1918 CIDR block, with a "/26" prefix, is required. The range cannot overlap with any
	// prefixes either in the consumer VPC network or in use by the private clouds attached to that VPC network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/vmwareengine_network_policy#edge_services_cidr VmwareengineNetworkPolicy#edge_services_cidr}
	EdgeServicesCidr *string `field:"required" json:"edgeServicesCidr" yaml:"edgeServicesCidr"`
	// The resource name of the location (region) to create the new network policy in.
	//
	// Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
	// For example: projects/my-project/locations/us-central1
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/vmwareengine_network_policy#location VmwareengineNetworkPolicy#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID of the Network Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/vmwareengine_network_policy#name VmwareengineNetworkPolicy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The relative resource name of the VMware Engine network.
	//
	// Specify the name in the following form:
	// projects/{project}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId} where {project}
	// can either be a project number or a project ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/vmwareengine_network_policy#vmware_engine_network VmwareengineNetworkPolicy#vmware_engine_network}
	VmwareEngineNetwork *string `field:"required" json:"vmwareEngineNetwork" yaml:"vmwareEngineNetwork"`
	// User-provided description for this network policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/vmwareengine_network_policy#description VmwareengineNetworkPolicy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// external_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/vmwareengine_network_policy#external_ip VmwareengineNetworkPolicy#external_ip}
	ExternalIp *VmwareengineNetworkPolicyExternalIp `field:"optional" json:"externalIp" yaml:"externalIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/vmwareengine_network_policy#id VmwareengineNetworkPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// internet_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/vmwareengine_network_policy#internet_access VmwareengineNetworkPolicy#internet_access}
	InternetAccess *VmwareengineNetworkPolicyInternetAccess `field:"optional" json:"internetAccess" yaml:"internetAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/vmwareengine_network_policy#project VmwareengineNetworkPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.31.0/docs/resources/vmwareengine_network_policy#timeouts VmwareengineNetworkPolicy#timeouts}
	Timeouts *VmwareengineNetworkPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

