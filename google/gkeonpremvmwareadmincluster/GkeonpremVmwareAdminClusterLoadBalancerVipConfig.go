// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwareadmincluster


type GkeonpremVmwareAdminClusterLoadBalancerVipConfig struct {
	// The VIP which you previously set aside for the Kubernetes API of this VMware Admin Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/gkeonprem_vmware_admin_cluster#control_plane_vip GkeonpremVmwareAdminCluster#control_plane_vip}
	ControlPlaneVip *string `field:"required" json:"controlPlaneVip" yaml:"controlPlaneVip"`
	// The VIP to configure the load balancer for add-ons.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.1/docs/resources/gkeonprem_vmware_admin_cluster#addons_vip GkeonpremVmwareAdminCluster#addons_vip}
	AddonsVip *string `field:"optional" json:"addonsVip" yaml:"addonsVip"`
}

