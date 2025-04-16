// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterMonitoringConfigManagedPrometheus struct {
	// Whether or not the managed collection is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// auto_monitoring_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/container_cluster#auto_monitoring_config ContainerCluster#auto_monitoring_config}
	AutoMonitoringConfig *ContainerClusterMonitoringConfigManagedPrometheusAutoMonitoringConfig `field:"optional" json:"autoMonitoringConfig" yaml:"autoMonitoringConfig"`
}

