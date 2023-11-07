// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iapappengineversioniambinding


type IapAppEngineVersionIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.5.0/docs/resources/iap_app_engine_version_iam_binding#expression IapAppEngineVersionIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.5.0/docs/resources/iap_app_engine_version_iam_binding#title IapAppEngineVersionIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.5.0/docs/resources/iap_app_engine_version_iam_binding#description IapAppEngineVersionIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

