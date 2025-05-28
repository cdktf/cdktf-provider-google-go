// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterMonitoringConfigManagedPrometheusAutoMonitoringConfig struct {
	// The scope of auto-monitoring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/container_cluster#scope ContainerCluster#scope}
	Scope *string `field:"required" json:"scope" yaml:"scope"`
}

