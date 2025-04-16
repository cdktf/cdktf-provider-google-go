// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubmembership


type GkeHubMembershipAuthority struct {
	// A JSON Web Token (JWT) issuer URI.
	//
	// 'issuer' must start with 'https://' and // be a valid
	// with length <2000 characters. For example: 'https://container.googleapis.com/v1/projects/my-project/locations/us-west1/clusters/my-cluster'. If the cluster is provisioned with Terraform, this is '"https://container.googleapis.com/v1/${google_container_cluster.my-cluster.id}"'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/gke_hub_membership#issuer GkeHubMembership#issuer}
	Issuer *string `field:"required" json:"issuer" yaml:"issuer"`
}

