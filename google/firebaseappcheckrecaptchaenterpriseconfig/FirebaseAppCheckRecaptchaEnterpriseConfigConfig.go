// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package firebaseappcheckrecaptchaenterpriseconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FirebaseAppCheckRecaptchaEnterpriseConfigConfig struct {
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
	// The ID of an [Web App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/firebase_app_check_recaptcha_enterprise_config#app_id FirebaseAppCheckRecaptchaEnterpriseConfig#app_id}
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// The score-based site key created in reCAPTCHA Enterprise used to invoke reCAPTCHA and generate the reCAPTCHA tokens for your application.
	//
	// **Important**: This is not the siteSecret (as it is in reCAPTCHA v3), but rather your score-based reCAPTCHA Enterprise site key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/firebase_app_check_recaptcha_enterprise_config#site_key FirebaseAppCheckRecaptchaEnterpriseConfig#site_key}
	SiteKey *string `field:"required" json:"siteKey" yaml:"siteKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/firebase_app_check_recaptcha_enterprise_config#id FirebaseAppCheckRecaptchaEnterpriseConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/firebase_app_check_recaptcha_enterprise_config#project FirebaseAppCheckRecaptchaEnterpriseConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/firebase_app_check_recaptcha_enterprise_config#timeouts FirebaseAppCheckRecaptchaEnterpriseConfig#timeouts}
	Timeouts *FirebaseAppCheckRecaptchaEnterpriseConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Specifies the duration for which App Check tokens exchanged from reCAPTCHA Enterprise artifacts will be valid.
	//
	// If unset, a default value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive.
	//
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.22.0/docs/resources/firebase_app_check_recaptcha_enterprise_config#token_ttl FirebaseAppCheckRecaptchaEnterpriseConfig#token_ttl}
	TokenTtl *string `field:"optional" json:"tokenTtl" yaml:"tokenTtl"`
}

