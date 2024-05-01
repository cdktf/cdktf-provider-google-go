// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerattachedcluster


type ContainerAttachedClusterLoggingConfigComponentConfig struct {
	// The components to be enabled. Possible values: ["SYSTEM_COMPONENTS", "WORKLOADS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.27.0/docs/resources/container_attached_cluster#enable_components ContainerAttachedCluster#enable_components}
	EnableComponents *[]*string `field:"optional" json:"enableComponents" yaml:"enableComponents"`
}

