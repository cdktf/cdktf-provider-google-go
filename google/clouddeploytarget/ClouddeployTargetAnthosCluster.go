// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package clouddeploytarget


type ClouddeployTargetAnthosCluster struct {
	// Membership of the GKE Hub-registered cluster to which to apply the Skaffold configuration. Format is `projects/{project}/locations/{location}/memberships/{membership_name}`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.37.0/docs/resources/clouddeploy_target#membership ClouddeployTarget#membership}
	Membership *string `field:"optional" json:"membership" yaml:"membership"`
}

