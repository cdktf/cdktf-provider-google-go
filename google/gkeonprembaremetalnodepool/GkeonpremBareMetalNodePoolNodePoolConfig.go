// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalnodepool


type GkeonpremBareMetalNodePoolNodePoolConfig struct {
	// node_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/gkeonprem_bare_metal_node_pool#node_configs GkeonpremBareMetalNodePool#node_configs}
	NodeConfigs interface{} `field:"required" json:"nodeConfigs" yaml:"nodeConfigs"`
	// The map of Kubernetes labels (key/value pairs) to be applied to each node.
	//
	// These will added in addition to any default label(s)
	// that Kubernetes may apply to the node. In case of conflict in
	// label keys, the applied set may differ depending on the Kubernetes
	// version -- it's best to assume the behavior is undefined and
	// conflicts should be avoided. For more information, including usage
	// and the valid values, see:
	//   - http://kubernetes.io/v1.1/docs/user-guide/labels.html
	// An object containing a list of "key": value pairs.
	// For example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/gkeonprem_bare_metal_node_pool#labels GkeonpremBareMetalNodePool#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Specifies the nodes operating system (default: LINUX).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/gkeonprem_bare_metal_node_pool#operating_system GkeonpremBareMetalNodePool#operating_system}
	OperatingSystem *string `field:"optional" json:"operatingSystem" yaml:"operatingSystem"`
	// taints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/gkeonprem_bare_metal_node_pool#taints GkeonpremBareMetalNodePool#taints}
	Taints interface{} `field:"optional" json:"taints" yaml:"taints"`
}

