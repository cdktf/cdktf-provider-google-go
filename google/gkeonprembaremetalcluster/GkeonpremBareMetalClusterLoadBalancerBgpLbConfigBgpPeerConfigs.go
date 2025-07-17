// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterLoadBalancerBgpLbConfigBgpPeerConfigs struct {
	// BGP autonomous system number (ASN) for the network that contains the external peer device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/gkeonprem_bare_metal_cluster#asn GkeonpremBareMetalCluster#asn}
	Asn *float64 `field:"required" json:"asn" yaml:"asn"`
	// The IP address of the external peer device.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/gkeonprem_bare_metal_cluster#ip_address GkeonpremBareMetalCluster#ip_address}
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// The IP address of the control plane node that connects to the external peer.
	//
	// If you don't specify any control plane nodes, all control plane nodes
	// can connect to the external peer. If you specify one or more IP addresses,
	// only the nodes specified participate in peering sessions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.44.0/docs/resources/gkeonprem_bare_metal_cluster#control_plane_nodes GkeonpremBareMetalCluster#control_plane_nodes}
	ControlPlaneNodes *[]*string `field:"optional" json:"controlPlaneNodes" yaml:"controlPlaneNodes"`
}

