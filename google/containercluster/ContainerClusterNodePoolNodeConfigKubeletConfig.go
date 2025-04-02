// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodePoolNodeConfigKubeletConfig struct {
	// Defines a comma-separated allowlist of unsafe sysctls or sysctl patterns which can be set on the Pods.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/container_cluster#allowed_unsafe_sysctls ContainerCluster#allowed_unsafe_sysctls}
	AllowedUnsafeSysctls *[]*string `field:"optional" json:"allowedUnsafeSysctls" yaml:"allowedUnsafeSysctls"`
	// Defines the maximum number of container log files that can be present for a container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/container_cluster#container_log_max_files ContainerCluster#container_log_max_files}
	ContainerLogMaxFiles *float64 `field:"optional" json:"containerLogMaxFiles" yaml:"containerLogMaxFiles"`
	// Defines the maximum size of the container log file before it is rotated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/container_cluster#container_log_max_size ContainerCluster#container_log_max_size}
	ContainerLogMaxSize *string `field:"optional" json:"containerLogMaxSize" yaml:"containerLogMaxSize"`
	// Enable CPU CFS quota enforcement for containers that specify CPU limits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/container_cluster#cpu_cfs_quota ContainerCluster#cpu_cfs_quota}
	CpuCfsQuota interface{} `field:"optional" json:"cpuCfsQuota" yaml:"cpuCfsQuota"`
	// Set the CPU CFS quota period value 'cpu.cfs_period_us'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/container_cluster#cpu_cfs_quota_period ContainerCluster#cpu_cfs_quota_period}
	CpuCfsQuotaPeriod *string `field:"optional" json:"cpuCfsQuotaPeriod" yaml:"cpuCfsQuotaPeriod"`
	// Control the CPU management policy on the node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/container_cluster#cpu_manager_policy ContainerCluster#cpu_manager_policy}
	CpuManagerPolicy *string `field:"optional" json:"cpuManagerPolicy" yaml:"cpuManagerPolicy"`
	// Defines the percent of disk usage after which image garbage collection is always run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/container_cluster#image_gc_high_threshold_percent ContainerCluster#image_gc_high_threshold_percent}
	ImageGcHighThresholdPercent *float64 `field:"optional" json:"imageGcHighThresholdPercent" yaml:"imageGcHighThresholdPercent"`
	// Defines the percent of disk usage before which image garbage collection is never run.
	//
	// Lowest disk usage to garbage collect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/container_cluster#image_gc_low_threshold_percent ContainerCluster#image_gc_low_threshold_percent}
	ImageGcLowThresholdPercent *float64 `field:"optional" json:"imageGcLowThresholdPercent" yaml:"imageGcLowThresholdPercent"`
	// Defines the maximum age an image can be unused before it is garbage collected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/container_cluster#image_maximum_gc_age ContainerCluster#image_maximum_gc_age}
	ImageMaximumGcAge *string `field:"optional" json:"imageMaximumGcAge" yaml:"imageMaximumGcAge"`
	// Defines the minimum age for an unused image before it is garbage collected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/container_cluster#image_minimum_gc_age ContainerCluster#image_minimum_gc_age}
	ImageMinimumGcAge *string `field:"optional" json:"imageMinimumGcAge" yaml:"imageMinimumGcAge"`
	// Controls whether the kubelet read-only port is enabled.
	//
	// It is strongly recommended to set this to `FALSE`. Possible values: `TRUE`, `FALSE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/container_cluster#insecure_kubelet_readonly_port_enabled ContainerCluster#insecure_kubelet_readonly_port_enabled}
	InsecureKubeletReadonlyPortEnabled *string `field:"optional" json:"insecureKubeletReadonlyPortEnabled" yaml:"insecureKubeletReadonlyPortEnabled"`
	// Controls the maximum number of processes allowed to run in a pod.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/container_cluster#pod_pids_limit ContainerCluster#pod_pids_limit}
	PodPidsLimit *float64 `field:"optional" json:"podPidsLimit" yaml:"podPidsLimit"`
}

