// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerattachedcluster


type ContainerAttachedClusterMonitoringConfig struct {
	// managed_prometheus_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.10.0/docs/resources/container_attached_cluster#managed_prometheus_config ContainerAttachedCluster#managed_prometheus_config}
	ManagedPrometheusConfig *ContainerAttachedClusterMonitoringConfigManagedPrometheusConfig `field:"optional" json:"managedPrometheusConfig" yaml:"managedPrometheusConfig"`
}

