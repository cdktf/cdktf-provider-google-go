// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterLoadBalancerVipConfig struct {
	// The VIP which you previously set aside for the Kubernetes API of this Bare Metal User Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/gkeonprem_bare_metal_cluster#control_plane_vip GkeonpremBareMetalCluster#control_plane_vip}
	ControlPlaneVip *string `field:"required" json:"controlPlaneVip" yaml:"controlPlaneVip"`
	// The VIP which you previously set aside for ingress traffic into this Bare Metal User Cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/gkeonprem_bare_metal_cluster#ingress_vip GkeonpremBareMetalCluster#ingress_vip}
	IngressVip *string `field:"required" json:"ingressVip" yaml:"ingressVip"`
}

