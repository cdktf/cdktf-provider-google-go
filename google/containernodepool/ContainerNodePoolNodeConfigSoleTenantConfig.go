// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containernodepool


type ContainerNodePoolNodeConfigSoleTenantConfig struct {
	// node_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_node_pool#node_affinity ContainerNodePool#node_affinity}
	NodeAffinity interface{} `field:"required" json:"nodeAffinity" yaml:"nodeAffinity"`
	// Specifies the minimum number of vCPUs that each sole tenant node must have to use CPU overcommit.
	//
	// If not specified, the CPU overcommit feature is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/container_node_pool#min_node_cpus ContainerNodePool#min_node_cpus}
	MinNodeCpus *float64 `field:"optional" json:"minNodeCpus" yaml:"minNodeCpus"`
}

