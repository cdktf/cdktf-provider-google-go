// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appengineflexibleappversion


type AppEngineFlexibleAppVersionAutomaticScalingRequestUtilization struct {
	// Target number of concurrent requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.12.0/docs/resources/app_engine_flexible_app_version#target_concurrent_requests AppEngineFlexibleAppVersion#target_concurrent_requests}
	TargetConcurrentRequests *float64 `field:"optional" json:"targetConcurrentRequests" yaml:"targetConcurrentRequests"`
	// Target requests per second.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.12.0/docs/resources/app_engine_flexible_app_version#target_request_count_per_second AppEngineFlexibleAppVersion#target_request_count_per_second}
	TargetRequestCountPerSecond *string `field:"optional" json:"targetRequestCountPerSecond" yaml:"targetRequestCountPerSecond"`
}

