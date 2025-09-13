// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerawsnodepool


type ContainerAwsNodePoolUpdateSettings struct {
	// surge_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/container_aws_node_pool#surge_settings ContainerAwsNodePool#surge_settings}
	SurgeSettings *ContainerAwsNodePoolUpdateSettingsSurgeSettings `field:"optional" json:"surgeSettings" yaml:"surgeSettings"`
}

