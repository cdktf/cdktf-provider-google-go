// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerawsnodepool


type ContainerAwsNodePoolKubeletConfig struct {
	// Whether or not to enable CPU CFS quota. Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/container_aws_node_pool#cpu_cfs_quota ContainerAwsNodePool#cpu_cfs_quota}
	CpuCfsQuota interface{} `field:"optional" json:"cpuCfsQuota" yaml:"cpuCfsQuota"`
	// Optional. The CPU CFS quota period to use for the node. Defaults to "100ms".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/container_aws_node_pool#cpu_cfs_quota_period ContainerAwsNodePool#cpu_cfs_quota_period}
	CpuCfsQuotaPeriod *string `field:"optional" json:"cpuCfsQuotaPeriod" yaml:"cpuCfsQuotaPeriod"`
	// The CpuManagerPolicy to use for the node. Defaults to "none".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/container_aws_node_pool#cpu_manager_policy ContainerAwsNodePool#cpu_manager_policy}
	CpuManagerPolicy *string `field:"optional" json:"cpuManagerPolicy" yaml:"cpuManagerPolicy"`
	// Optional.
	//
	// The maximum number of PIDs in each pod running on the node. The limit scales automatically based on underlying machine size if left unset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.27.0/docs/resources/container_aws_node_pool#pod_pids_limit ContainerAwsNodePool#pod_pids_limit}
	PodPidsLimit *float64 `field:"optional" json:"podPidsLimit" yaml:"podPidsLimit"`
}

