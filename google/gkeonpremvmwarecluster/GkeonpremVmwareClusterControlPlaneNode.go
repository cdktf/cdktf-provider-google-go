// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwarecluster


type GkeonpremVmwareClusterControlPlaneNode struct {
	// auto_resize_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/gkeonprem_vmware_cluster#auto_resize_config GkeonpremVmwareCluster#auto_resize_config}
	AutoResizeConfig *GkeonpremVmwareClusterControlPlaneNodeAutoResizeConfig `field:"optional" json:"autoResizeConfig" yaml:"autoResizeConfig"`
	// The number of CPUs for each admin cluster node that serve as control planes for this VMware User Cluster.
	//
	// (default: 4 CPUs)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/gkeonprem_vmware_cluster#cpus GkeonpremVmwareCluster#cpus}
	Cpus *float64 `field:"optional" json:"cpus" yaml:"cpus"`
	// The megabytes of memory for each admin cluster node that serves as a control plane for this VMware User Cluster (default: 8192 MB memory).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/gkeonprem_vmware_cluster#memory GkeonpremVmwareCluster#memory}
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// The number of control plane nodes for this VMware User Cluster. (default: 1 replica).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.2.0/docs/resources/gkeonprem_vmware_cluster#replicas GkeonpremVmwareCluster#replicas}
	Replicas *float64 `field:"optional" json:"replicas" yaml:"replicas"`
}

