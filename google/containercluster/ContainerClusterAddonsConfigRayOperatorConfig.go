// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterAddonsConfigRayOperatorConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/container_cluster#enabled ContainerCluster#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// ray_cluster_logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/container_cluster#ray_cluster_logging_config ContainerCluster#ray_cluster_logging_config}
	RayClusterLoggingConfig *ContainerClusterAddonsConfigRayOperatorConfigRayClusterLoggingConfig `field:"optional" json:"rayClusterLoggingConfig" yaml:"rayClusterLoggingConfig"`
	// ray_cluster_monitoring_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.34.0/docs/resources/container_cluster#ray_cluster_monitoring_config ContainerCluster#ray_cluster_monitoring_config}
	RayClusterMonitoringConfig *ContainerClusterAddonsConfigRayOperatorConfigRayClusterMonitoringConfig `field:"optional" json:"rayClusterMonitoringConfig" yaml:"rayClusterMonitoringConfig"`
}

