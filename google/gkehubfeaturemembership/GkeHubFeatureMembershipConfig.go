// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gkehubfeaturemembership

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GkeHubFeatureMembershipConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/gke_hub_feature_membership#feature GkeHubFeatureMembership#feature}
	Feature *string `field:"required" json:"feature" yaml:"feature"`
	// The location of the feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/gke_hub_feature_membership#location GkeHubFeatureMembership#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of the membership.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/gke_hub_feature_membership#membership GkeHubFeatureMembership#membership}
	Membership *string `field:"required" json:"membership" yaml:"membership"`
	// configmanagement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/gke_hub_feature_membership#configmanagement GkeHubFeatureMembership#configmanagement}
	Configmanagement *GkeHubFeatureMembershipConfigmanagement `field:"optional" json:"configmanagement" yaml:"configmanagement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/gke_hub_feature_membership#id GkeHubFeatureMembership#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The location of the membership.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/gke_hub_feature_membership#membership_location GkeHubFeatureMembership#membership_location}
	MembershipLocation *string `field:"optional" json:"membershipLocation" yaml:"membershipLocation"`
	// mesh block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/gke_hub_feature_membership#mesh GkeHubFeatureMembership#mesh}
	Mesh *GkeHubFeatureMembershipMesh `field:"optional" json:"mesh" yaml:"mesh"`
	// policycontroller block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/gke_hub_feature_membership#policycontroller GkeHubFeatureMembership#policycontroller}
	Policycontroller *GkeHubFeatureMembershipPolicycontroller `field:"optional" json:"policycontroller" yaml:"policycontroller"`
	// The project of the feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/gke_hub_feature_membership#project GkeHubFeatureMembership#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.32.0/docs/resources/gke_hub_feature_membership#timeouts GkeHubFeatureMembership#timeouts}
	Timeouts *GkeHubFeatureMembershipTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

