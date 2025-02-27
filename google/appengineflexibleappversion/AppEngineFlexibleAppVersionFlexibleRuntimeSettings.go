// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appengineflexibleappversion


type AppEngineFlexibleAppVersionFlexibleRuntimeSettings struct {
	// Operating System of the application runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/app_engine_flexible_app_version#operating_system AppEngineFlexibleAppVersion#operating_system}
	OperatingSystem *string `field:"optional" json:"operatingSystem" yaml:"operatingSystem"`
	// The runtime version of an App Engine flexible application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/app_engine_flexible_app_version#runtime_version AppEngineFlexibleAppVersion#runtime_version}
	RuntimeVersion *string `field:"optional" json:"runtimeVersion" yaml:"runtimeVersion"`
}

