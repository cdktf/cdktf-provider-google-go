// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firebaseapphostingtraffic

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirebaseAppHostingTrafficConfig struct {
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
	// Id of the backend that this Traffic config applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/firebase_app_hosting_traffic#backend FirebaseAppHostingTraffic#backend}
	Backend *string `field:"required" json:"backend" yaml:"backend"`
	// The location the Backend that this Traffic config applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/firebase_app_hosting_traffic#location FirebaseAppHostingTraffic#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/firebase_app_hosting_traffic#id FirebaseAppHostingTraffic#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/firebase_app_hosting_traffic#project FirebaseAppHostingTraffic#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// rollout_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/firebase_app_hosting_traffic#rollout_policy FirebaseAppHostingTraffic#rollout_policy}
	RolloutPolicy *FirebaseAppHostingTrafficRolloutPolicy `field:"optional" json:"rolloutPolicy" yaml:"rolloutPolicy"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/firebase_app_hosting_traffic#target FirebaseAppHostingTraffic#target}
	Target *FirebaseAppHostingTrafficTarget `field:"optional" json:"target" yaml:"target"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/firebase_app_hosting_traffic#timeouts FirebaseAppHostingTraffic#timeouts}
	Timeouts *FirebaseAppHostingTrafficTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

