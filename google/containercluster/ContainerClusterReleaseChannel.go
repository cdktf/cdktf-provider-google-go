// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterReleaseChannel struct {
	// The selected release channel.
	//
	// Accepted values are:
	// * UNSPECIFIED: Not set.
	// * RAPID: Weekly upgrade cadence; Early testers and developers who requires new features.
	// * REGULAR: Multiple per month upgrade cadence; Production users who need features not yet offered in the Stable channel.
	// * STABLE: Every few months upgrade cadence; Production users who need stability above all else, and for whom frequent upgrades are too risky.
	// * EXTENDED: GKE provides extended support for Kubernetes minor versions through the Extended channel. With this channel, you can stay on a minor version for up to 24 months.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.20.0/docs/resources/container_cluster#channel ContainerCluster#channel}
	Channel *string `field:"required" json:"channel" yaml:"channel"`
}

