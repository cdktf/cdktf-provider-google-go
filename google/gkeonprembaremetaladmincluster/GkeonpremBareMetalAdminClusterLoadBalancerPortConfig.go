// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetaladmincluster


type GkeonpremBareMetalAdminClusterLoadBalancerPortConfig struct {
	// The port that control plane hosted load balancers will listen on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.3.0/docs/resources/gkeonprem_bare_metal_admin_cluster#control_plane_load_balancer_port GkeonpremBareMetalAdminCluster#control_plane_load_balancer_port}
	ControlPlaneLoadBalancerPort *float64 `field:"required" json:"controlPlaneLoadBalancerPort" yaml:"controlPlaneLoadBalancerPort"`
}

