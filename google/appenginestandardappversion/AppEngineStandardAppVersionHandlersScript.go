// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appenginestandardappversion


type AppEngineStandardAppVersionHandlersScript struct {
	// Path to the script from the application root directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.40.0/docs/resources/app_engine_standard_app_version#script_path AppEngineStandardAppVersion#script_path}
	ScriptPath *string `field:"required" json:"scriptPath" yaml:"scriptPath"`
}

