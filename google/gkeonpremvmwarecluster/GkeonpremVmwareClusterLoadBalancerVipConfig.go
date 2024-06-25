// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonpremvmwarecluster


type GkeonpremVmwareClusterLoadBalancerVipConfig struct {
	// The VIP which you previously set aside for the Kubernetes API of this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.35.0/docs/resources/gkeonprem_vmware_cluster#control_plane_vip GkeonpremVmwareCluster#control_plane_vip}
	ControlPlaneVip *string `field:"optional" json:"controlPlaneVip" yaml:"controlPlaneVip"`
	// The VIP which you previously set aside for ingress traffic into this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.35.0/docs/resources/gkeonprem_vmware_cluster#ingress_vip GkeonpremVmwareCluster#ingress_vip}
	IngressVip *string `field:"optional" json:"ingressVip" yaml:"ingressVip"`
}

