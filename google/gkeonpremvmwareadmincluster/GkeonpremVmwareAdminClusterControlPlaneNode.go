// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwareadmincluster


type GkeonpremVmwareAdminClusterControlPlaneNode struct {
	// The number of vCPUs for the control-plane node of the admin cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/gkeonprem_vmware_admin_cluster#cpus GkeonpremVmwareAdminCluster#cpus}
	Cpus *float64 `field:"optional" json:"cpus" yaml:"cpus"`
	// The number of mebibytes of memory for the control-plane node of the admin cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/gkeonprem_vmware_admin_cluster#memory GkeonpremVmwareAdminCluster#memory}
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
	// The number of control plane nodes for this VMware admin cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/gkeonprem_vmware_admin_cluster#replicas GkeonpremVmwareAdminCluster#replicas}
	Replicas *float64 `field:"optional" json:"replicas" yaml:"replicas"`
}

