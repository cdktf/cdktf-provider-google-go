// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterMonitoringConfig struct {
	// advanced_datapath_observability_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.24.0/docs/resources/container_cluster#advanced_datapath_observability_config ContainerCluster#advanced_datapath_observability_config}
	AdvancedDatapathObservabilityConfig *ContainerClusterMonitoringConfigAdvancedDatapathObservabilityConfig `field:"optional" json:"advancedDatapathObservabilityConfig" yaml:"advancedDatapathObservabilityConfig"`
	// GKE components exposing metrics. Valid values include SYSTEM_COMPONENTS, APISERVER, SCHEDULER, CONTROLLER_MANAGER, STORAGE, HPA, POD, DAEMONSET, DEPLOYMENT and STATEFULSET.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.24.0/docs/resources/container_cluster#enable_components ContainerCluster#enable_components}
	EnableComponents *[]*string `field:"optional" json:"enableComponents" yaml:"enableComponents"`
	// managed_prometheus block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.24.0/docs/resources/container_cluster#managed_prometheus ContainerCluster#managed_prometheus}
	ManagedPrometheus *ContainerClusterMonitoringConfigManagedPrometheus `field:"optional" json:"managedPrometheus" yaml:"managedPrometheus"`
}

