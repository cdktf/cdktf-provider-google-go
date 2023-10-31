// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package edgecontainernodepool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EdgecontainerNodePoolConfig struct {
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
	// The name of the target Distributed Cloud Edge Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.4.0/docs/resources/edgecontainer_node_pool#cluster EdgecontainerNodePool#cluster}
	Cluster *string `field:"required" json:"cluster" yaml:"cluster"`
	// The location of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.4.0/docs/resources/edgecontainer_node_pool#location EdgecontainerNodePool#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The resource name of the node pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.4.0/docs/resources/edgecontainer_node_pool#name EdgecontainerNodePool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The number of nodes in the pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.4.0/docs/resources/edgecontainer_node_pool#node_count EdgecontainerNodePool#node_count}
	NodeCount *float64 `field:"required" json:"nodeCount" yaml:"nodeCount"`
	// Name of the Google Distributed Cloud Edge zone where this node pool will be created. For example: 'us-central1-edge-customer-a'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.4.0/docs/resources/edgecontainer_node_pool#node_location EdgecontainerNodePool#node_location}
	NodeLocation *string `field:"required" json:"nodeLocation" yaml:"nodeLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.4.0/docs/resources/edgecontainer_node_pool#id EdgecontainerNodePool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels associated with this resource.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.4.0/docs/resources/edgecontainer_node_pool#labels EdgecontainerNodePool#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// local_disk_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.4.0/docs/resources/edgecontainer_node_pool#local_disk_encryption EdgecontainerNodePool#local_disk_encryption}
	LocalDiskEncryption *EdgecontainerNodePoolLocalDiskEncryption `field:"optional" json:"localDiskEncryption" yaml:"localDiskEncryption"`
	// Only machines matching this filter will be allowed to join the node pool.
	//
	// The filtering language accepts strings like "name=<name>", and is
	// documented in more detail in [AIP-160](https://google.aip.dev/160).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.4.0/docs/resources/edgecontainer_node_pool#machine_filter EdgecontainerNodePool#machine_filter}
	MachineFilter *string `field:"optional" json:"machineFilter" yaml:"machineFilter"`
	// node_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.4.0/docs/resources/edgecontainer_node_pool#node_config EdgecontainerNodePool#node_config}
	NodeConfig *EdgecontainerNodePoolNodeConfig `field:"optional" json:"nodeConfig" yaml:"nodeConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.4.0/docs/resources/edgecontainer_node_pool#project EdgecontainerNodePool#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.4.0/docs/resources/edgecontainer_node_pool#timeouts EdgecontainerNodePool#timeouts}
	Timeouts *EdgecontainerNodePoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

