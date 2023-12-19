// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalnodepool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeonpremBareMetalNodePoolConfig struct {
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
	// The cluster this node pool belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/gkeonprem_bare_metal_node_pool#bare_metal_cluster GkeonpremBareMetalNodePool#bare_metal_cluster}
	BareMetalCluster *string `field:"required" json:"bareMetalCluster" yaml:"bareMetalCluster"`
	// The location of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/gkeonprem_bare_metal_node_pool#location GkeonpremBareMetalNodePool#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The bare metal node pool name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/gkeonprem_bare_metal_node_pool#name GkeonpremBareMetalNodePool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// node_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/gkeonprem_bare_metal_node_pool#node_pool_config GkeonpremBareMetalNodePool#node_pool_config}
	NodePoolConfig *GkeonpremBareMetalNodePoolNodePoolConfig `field:"required" json:"nodePoolConfig" yaml:"nodePoolConfig"`
	// Annotations on the Bare Metal Node Pool.
	//
	// This field has the same restrictions as Kubernetes annotations.
	// The total size of all keys and values combined is limited to 256k.
	// Key can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/).
	// Prefix must be a DNS subdomain.
	// Name must be 63 characters or less, begin and end with alphanumerics,
	// with dashes (-), underscores (_), dots (.), and alphanumerics between.
	//
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/gkeonprem_bare_metal_node_pool#annotations GkeonpremBareMetalNodePool#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// The display name for the Bare Metal Node Pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/gkeonprem_bare_metal_node_pool#display_name GkeonpremBareMetalNodePool#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/gkeonprem_bare_metal_node_pool#id GkeonpremBareMetalNodePool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/gkeonprem_bare_metal_node_pool#project GkeonpremBareMetalNodePool#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/gkeonprem_bare_metal_node_pool#timeouts GkeonpremBareMetalNodePool#timeouts}
	Timeouts *GkeonpremBareMetalNodePoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

