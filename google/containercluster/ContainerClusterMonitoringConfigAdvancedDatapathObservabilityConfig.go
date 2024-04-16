// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterMonitoringConfigAdvancedDatapathObservabilityConfig struct {
	// Whether or not the advanced datapath metrics are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.25.0/docs/resources/container_cluster#enable_metrics ContainerCluster#enable_metrics}
	EnableMetrics interface{} `field:"required" json:"enableMetrics" yaml:"enableMetrics"`
	// Whether or not Relay is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.25.0/docs/resources/container_cluster#enable_relay ContainerCluster#enable_relay}
	EnableRelay interface{} `field:"optional" json:"enableRelay" yaml:"enableRelay"`
	// Mode used to make Relay available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.25.0/docs/resources/container_cluster#relay_mode ContainerCluster#relay_mode}
	RelayMode *string `field:"optional" json:"relayMode" yaml:"relayMode"`
}

