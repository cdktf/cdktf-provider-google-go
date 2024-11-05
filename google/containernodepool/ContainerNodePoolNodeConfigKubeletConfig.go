// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containernodepool


type ContainerNodePoolNodeConfigKubeletConfig struct {
	// Enable CPU CFS quota enforcement for containers that specify CPU limits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/container_node_pool#cpu_cfs_quota ContainerNodePool#cpu_cfs_quota}
	CpuCfsQuota interface{} `field:"optional" json:"cpuCfsQuota" yaml:"cpuCfsQuota"`
	// Set the CPU CFS quota period value 'cpu.cfs_period_us'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/container_node_pool#cpu_cfs_quota_period ContainerNodePool#cpu_cfs_quota_period}
	CpuCfsQuotaPeriod *string `field:"optional" json:"cpuCfsQuotaPeriod" yaml:"cpuCfsQuotaPeriod"`
	// Control the CPU management policy on the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/container_node_pool#cpu_manager_policy ContainerNodePool#cpu_manager_policy}
	CpuManagerPolicy *string `field:"optional" json:"cpuManagerPolicy" yaml:"cpuManagerPolicy"`
	// Controls whether the kubelet read-only port is enabled.
	//
	// It is strongly recommended to set this to `FALSE`. Possible values: `TRUE`, `FALSE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/container_node_pool#insecure_kubelet_readonly_port_enabled ContainerNodePool#insecure_kubelet_readonly_port_enabled}
	InsecureKubeletReadonlyPortEnabled *string `field:"optional" json:"insecureKubeletReadonlyPortEnabled" yaml:"insecureKubeletReadonlyPortEnabled"`
	// Controls the maximum number of processes allowed to run in a pod.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.10.0/docs/resources/container_node_pool#pod_pids_limit ContainerNodePool#pod_pids_limit}
	PodPidsLimit *float64 `field:"optional" json:"podPidsLimit" yaml:"podPidsLimit"`
}

