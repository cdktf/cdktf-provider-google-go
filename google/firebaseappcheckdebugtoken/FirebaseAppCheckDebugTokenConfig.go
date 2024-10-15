// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firebaseappcheckdebugtoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirebaseAppCheckDebugTokenConfig struct {
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
	// The ID of a [Web App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id), [Apple App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.iosApps#IosApp.FIELDS.app_id), or [Android App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.androidApps#AndroidApp.FIELDS.app_id).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/firebase_app_check_debug_token#app_id FirebaseAppCheckDebugToken#app_id}
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// A human readable display name used to identify this debug token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/firebase_app_check_debug_token#display_name FirebaseAppCheckDebugToken#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The secret token itself.
	//
	// Must be provided during creation, and must be a UUID4,
	// case insensitive. You may use a method of your choice such as random/random_uuid
	// to generate the token.
	//
	// This field is immutable once set, and cannot be updated. You can, however, delete
	// this debug token to revoke it.
	//
	// For security reasons, this field will never be populated in any response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/firebase_app_check_debug_token#token FirebaseAppCheckDebugToken#token}
	Token *string `field:"required" json:"token" yaml:"token"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/firebase_app_check_debug_token#id FirebaseAppCheckDebugToken#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/firebase_app_check_debug_token#project FirebaseAppCheckDebugToken#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.7.0/docs/resources/firebase_app_check_debug_token#timeouts FirebaseAppCheckDebugToken#timeouts}
	Timeouts *FirebaseAppCheckDebugTokenTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

