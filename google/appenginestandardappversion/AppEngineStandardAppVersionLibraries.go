// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appenginestandardappversion


type AppEngineStandardAppVersionLibraries struct {
	// Name of the library. Example "django".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/app_engine_standard_app_version#name AppEngineStandardAppVersion#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Version of the library to select, or "latest".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.39.0/docs/resources/app_engine_standard_app_version#version AppEngineStandardAppVersion#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

