// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containernodepool


type ContainerNodePoolNodeConfigKubeletConfig struct {
	// Defines a comma-separated allowlist of unsafe sysctls or sysctl patterns which can be set on the Pods.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/container_node_pool#allowed_unsafe_sysctls ContainerNodePool#allowed_unsafe_sysctls}
	AllowedUnsafeSysctls *[]*string `field:"optional" json:"allowedUnsafeSysctls" yaml:"allowedUnsafeSysctls"`
	// Defines the maximum number of container log files that can be present for a container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/container_node_pool#container_log_max_files ContainerNodePool#container_log_max_files}
	ContainerLogMaxFiles *float64 `field:"optional" json:"containerLogMaxFiles" yaml:"containerLogMaxFiles"`
	// Defines the maximum size of the container log file before it is rotated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/container_node_pool#container_log_max_size ContainerNodePool#container_log_max_size}
	ContainerLogMaxSize *string `field:"optional" json:"containerLogMaxSize" yaml:"containerLogMaxSize"`
	// Enable CPU CFS quota enforcement for containers that specify CPU limits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/container_node_pool#cpu_cfs_quota ContainerNodePool#cpu_cfs_quota}
	CpuCfsQuota interface{} `field:"optional" json:"cpuCfsQuota" yaml:"cpuCfsQuota"`
	// Set the CPU CFS quota period value 'cpu.cfs_period_us'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/container_node_pool#cpu_cfs_quota_period ContainerNodePool#cpu_cfs_quota_period}
	CpuCfsQuotaPeriod *string `field:"optional" json:"cpuCfsQuotaPeriod" yaml:"cpuCfsQuotaPeriod"`
	// Control the CPU management policy on the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/container_node_pool#cpu_manager_policy ContainerNodePool#cpu_manager_policy}
	CpuManagerPolicy *string `field:"optional" json:"cpuManagerPolicy" yaml:"cpuManagerPolicy"`
	// Defines the maximum allowed grace period (in seconds) to use when terminating pods in response to a soft eviction threshold being met.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/container_node_pool#eviction_max_pod_grace_period_seconds ContainerNodePool#eviction_max_pod_grace_period_seconds}
	EvictionMaxPodGracePeriodSeconds *float64 `field:"optional" json:"evictionMaxPodGracePeriodSeconds" yaml:"evictionMaxPodGracePeriodSeconds"`
	// eviction_minimum_reclaim block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/container_node_pool#eviction_minimum_reclaim ContainerNodePool#eviction_minimum_reclaim}
	EvictionMinimumReclaim *ContainerNodePoolNodeConfigKubeletConfigEvictionMinimumReclaim `field:"optional" json:"evictionMinimumReclaim" yaml:"evictionMinimumReclaim"`
	// eviction_soft block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/container_node_pool#eviction_soft ContainerNodePool#eviction_soft}
	EvictionSoft *ContainerNodePoolNodeConfigKubeletConfigEvictionSoft `field:"optional" json:"evictionSoft" yaml:"evictionSoft"`
	// eviction_soft_grace_period block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/container_node_pool#eviction_soft_grace_period ContainerNodePool#eviction_soft_grace_period}
	EvictionSoftGracePeriod *ContainerNodePoolNodeConfigKubeletConfigEvictionSoftGracePeriod `field:"optional" json:"evictionSoftGracePeriod" yaml:"evictionSoftGracePeriod"`
	// Defines the percent of disk usage after which image garbage collection is always run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/container_node_pool#image_gc_high_threshold_percent ContainerNodePool#image_gc_high_threshold_percent}
	ImageGcHighThresholdPercent *float64 `field:"optional" json:"imageGcHighThresholdPercent" yaml:"imageGcHighThresholdPercent"`
	// Defines the percent of disk usage before which image garbage collection is never run.
	//
	// Lowest disk usage to garbage collect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/container_node_pool#image_gc_low_threshold_percent ContainerNodePool#image_gc_low_threshold_percent}
	ImageGcLowThresholdPercent *float64 `field:"optional" json:"imageGcLowThresholdPercent" yaml:"imageGcLowThresholdPercent"`
	// Defines the maximum age an image can be unused before it is garbage collected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/container_node_pool#image_maximum_gc_age ContainerNodePool#image_maximum_gc_age}
	ImageMaximumGcAge *string `field:"optional" json:"imageMaximumGcAge" yaml:"imageMaximumGcAge"`
	// Defines the minimum age for an unused image before it is garbage collected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/container_node_pool#image_minimum_gc_age ContainerNodePool#image_minimum_gc_age}
	ImageMinimumGcAge *string `field:"optional" json:"imageMinimumGcAge" yaml:"imageMinimumGcAge"`
	// Controls whether the kubelet read-only port is enabled.
	//
	// It is strongly recommended to set this to `FALSE`. Possible values: `TRUE`, `FALSE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/container_node_pool#insecure_kubelet_readonly_port_enabled ContainerNodePool#insecure_kubelet_readonly_port_enabled}
	InsecureKubeletReadonlyPortEnabled *string `field:"optional" json:"insecureKubeletReadonlyPortEnabled" yaml:"insecureKubeletReadonlyPortEnabled"`
	// Set the maximum number of image pulls in parallel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/container_node_pool#max_parallel_image_pulls ContainerNodePool#max_parallel_image_pulls}
	MaxParallelImagePulls *float64 `field:"optional" json:"maxParallelImagePulls" yaml:"maxParallelImagePulls"`
	// Controls the maximum number of processes allowed to run in a pod.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/container_node_pool#pod_pids_limit ContainerNodePool#pod_pids_limit}
	PodPidsLimit *float64 `field:"optional" json:"podPidsLimit" yaml:"podPidsLimit"`
	// Defines whether to enable single process OOM killer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/container_node_pool#single_process_oom_kill ContainerNodePool#single_process_oom_kill}
	SingleProcessOomKill interface{} `field:"optional" json:"singleProcessOomKill" yaml:"singleProcessOomKill"`
}

