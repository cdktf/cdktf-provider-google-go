// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appengineapplication


type AppEngineApplicationFeatureSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/app_engine_application#split_health_checks AppEngineApplication#split_health_checks}.
	SplitHealthChecks interface{} `field:"required" json:"splitHealthChecks" yaml:"splitHealthChecks"`
}

