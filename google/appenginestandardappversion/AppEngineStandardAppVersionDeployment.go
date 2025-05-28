// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appenginestandardappversion


type AppEngineStandardAppVersionDeployment struct {
	// files block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/app_engine_standard_app_version#files AppEngineStandardAppVersion#files}
	Files interface{} `field:"optional" json:"files" yaml:"files"`
	// zip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.37.0/docs/resources/app_engine_standard_app_version#zip AppEngineStandardAppVersion#zip}
	Zip *AppEngineStandardAppVersionDeploymentZip `field:"optional" json:"zip" yaml:"zip"`
}

