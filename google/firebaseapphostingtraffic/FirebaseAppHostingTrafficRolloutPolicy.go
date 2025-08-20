// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firebaseapphostingtraffic


type FirebaseAppHostingTrafficRolloutPolicy struct {
	// Specifies a branch that triggers a new build to be started with this policy.
	//
	// If not set, no automatic rollouts will happen.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/firebase_app_hosting_traffic#codebase_branch FirebaseAppHostingTraffic#codebase_branch}
	CodebaseBranch *string `field:"optional" json:"codebaseBranch" yaml:"codebaseBranch"`
	// A flag that, if true, prevents rollouts from being created via this RolloutPolicy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/firebase_app_hosting_traffic#disabled FirebaseAppHostingTraffic#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}

