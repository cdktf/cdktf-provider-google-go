// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appenginestandardappversion


type AppEngineStandardAppVersionDeploymentZip struct {
	// Source URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/app_engine_standard_app_version#source_url AppEngineStandardAppVersion#source_url}
	SourceUrl *string `field:"required" json:"sourceUrl" yaml:"sourceUrl"`
	// files count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.17.0/docs/resources/app_engine_standard_app_version#files_count AppEngineStandardAppVersion#files_count}
	FilesCount *float64 `field:"optional" json:"filesCount" yaml:"filesCount"`
}

