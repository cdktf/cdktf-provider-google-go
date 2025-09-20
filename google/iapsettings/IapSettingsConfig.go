// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iapsettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IapSettingsConfig struct {
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
	// The resource name of the IAP protected resource.
	//
	// Name can have below resources:
	// * organizations/{organization_id}
	// * folders/{folder_id}
	// * projects/{project_id}
	// * projects/{project_id}/iap_web
	// * projects/{project_id}/iap_web/compute
	// * projects/{project_id}/iap_web/compute-{region}
	// * projects/{project_id}/iap_web/compute/services/{service_id}
	// * projects/{project_id}/iap_web/compute-{region}/services/{service_id}
	// * projects/{project_id}/iap_web/appengine-{app_id}
	// * projects/{project_id}/iap_web/appengine-{app_id}/services/{service_id}
	// * projects/{project_id}/iap_web/appengine-{app_id}/services/{service_id}/version/{version_id}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/iap_settings#name IapSettings#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// access_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/iap_settings#access_settings IapSettings#access_settings}
	AccessSettings *IapSettingsAccessSettings `field:"optional" json:"accessSettings" yaml:"accessSettings"`
	// application_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/iap_settings#application_settings IapSettings#application_settings}
	ApplicationSettings *IapSettingsApplicationSettings `field:"optional" json:"applicationSettings" yaml:"applicationSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/iap_settings#id IapSettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/iap_settings#timeouts IapSettings#timeouts}
	Timeouts *IapSettingsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

