// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterLoadBalancerPortConfig struct {
	// The port that control plane hosted load balancers will listen on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.31.1/docs/resources/gkeonprem_bare_metal_cluster#control_plane_load_balancer_port GkeonpremBareMetalCluster#control_plane_load_balancer_port}
	ControlPlaneLoadBalancerPort *float64 `field:"required" json:"controlPlaneLoadBalancerPort" yaml:"controlPlaneLoadBalancerPort"`
}

