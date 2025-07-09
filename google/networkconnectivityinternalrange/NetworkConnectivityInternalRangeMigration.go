// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkconnectivityinternalrange


type NetworkConnectivityInternalRangeMigration struct {
	// Resource path as an URI of the source resource, for example a subnet.
	//
	// The project for the source resource should match the project for the
	// InternalRange.
	// An example /projects/{project}/regions/{region}/subnetworks/{subnet}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/network_connectivity_internal_range#source NetworkConnectivityInternalRange#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// Resource path of the target resource.
	//
	// The target project can be
	// different, as in the cases when migrating to peer networks. The resource
	// may not exist yet.
	// For example /projects/{project}/regions/{region}/subnetworks/{subnet}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.43.0/docs/resources/network_connectivity_internal_range#target NetworkConnectivityInternalRange#target}
	Target *string `field:"required" json:"target" yaml:"target"`
}

