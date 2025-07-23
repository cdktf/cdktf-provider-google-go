// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterPodAutoscaling struct {
	// HPA Profile is used to configure the Horizontal Pod Autoscaler (HPA) profile for the cluster.
	//
	// Available options include:
	// - NONE: Customers explicitly opt-out of HPA profiles.
	// - PERFORMANCE: PERFORMANCE is used when customers opt-in to the performance HPA profile. In this profile we support a higher number of HPAs per cluster and faster metrics collection for workload autoscaling.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_cluster#hpa_profile ContainerCluster#hpa_profile}
	HpaProfile *string `field:"required" json:"hpaProfile" yaml:"hpaProfile"`
}

