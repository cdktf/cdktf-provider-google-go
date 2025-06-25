// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firebaseappcheckserviceconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirebaseAppCheckServiceConfigConfig struct {
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
	// The identifier of the service to configure enforcement.
	//
	// Currently, the following service IDs are supported:
	//   firebasestorage.googleapis.com (Cloud Storage for Firebase)
	//   firebasedatabase.googleapis.com (Firebase Realtime Database)
	//   firestore.googleapis.com (Cloud Firestore)
	//   identitytoolkit.googleapis.com (Authentication)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/firebase_app_check_service_config#service_id FirebaseAppCheckServiceConfig#service_id}
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
	// The App Check enforcement mode for a service supported by App Check. Valid values are.
	//
	// (Unset)
	// Firebase App Check is not enforced for the service, nor are App Check metrics collected.
	// Though the service is not protected by App Check in this mode, other applicable protections,
	// such as user authorization, are still enforced. An unconfigured service is in this mode by default.
	// This is equivalent to OFF in the REST API. Deleting the Terraform resource will also switch the
	// enforcement to OFF for this service.
	//
	// UNENFORCED
	// Firebase App Check is not enforced for the service. App Check metrics are collected to help you
	// decide when to turn on enforcement for the service. Though the service is not protected by App Check
	// in this mode, other applicable protections, such as user authorization, are still enforced.
	//
	// ENFORCED
	// Firebase App Check is enforced for the service. The service will reject any request that attempts to
	// access your project's resources if it does not have valid App Check token attached, with some exceptions
	// depending on the service; for example, some services will still allow requests bearing the developer's
	// privileged service account credentials without an App Check token. App Check metrics continue to be
	// collected to help you detect issues with your App Check integration and monitor the composition of your
	// callers. While the service is protected by App Check, other applicable protections, such as user
	// authorization, continue to be enforced at the same time.
	//
	// Use caution when choosing to enforce App Check on a Firebase service. If your users have not updated
	// to an App Check capable version of your app, their apps will no longer be able to use your Firebase
	// services that are enforcing App Check. App Check metrics can help you decide whether to enforce App
	// Check on your Firebase services.
	//
	// If your app has not launched yet, you should enable enforcement immediately, since there are no outdated
	// clients in use. Possible values: ["UNENFORCED", "ENFORCED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/firebase_app_check_service_config#enforcement_mode FirebaseAppCheckServiceConfig#enforcement_mode}
	EnforcementMode *string `field:"optional" json:"enforcementMode" yaml:"enforcementMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/firebase_app_check_service_config#id FirebaseAppCheckServiceConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/firebase_app_check_service_config#project FirebaseAppCheckServiceConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.41.0/docs/resources/firebase_app_check_service_config#timeouts FirebaseAppCheckServiceConfig#timeouts}
	Timeouts *FirebaseAppCheckServiceConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

