// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetaladmincluster


type GkeonpremBareMetalAdminClusterLoadBalancerVipConfig struct {
	// The VIP which you previously set aside for the Kubernetes API of this Bare Metal Admin Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/gkeonprem_bare_metal_admin_cluster#control_plane_vip GkeonpremBareMetalAdminCluster#control_plane_vip}
	ControlPlaneVip *string `field:"required" json:"controlPlaneVip" yaml:"controlPlaneVip"`
}

