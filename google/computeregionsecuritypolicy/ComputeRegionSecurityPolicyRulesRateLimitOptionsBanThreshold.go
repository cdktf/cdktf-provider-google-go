// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionsecuritypolicy


type ComputeRegionSecurityPolicyRulesRateLimitOptionsBanThreshold struct {
	// Number of HTTP(S) requests for calculating the threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/compute_region_security_policy#count ComputeRegionSecurityPolicy#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Interval over which the threshold is computed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/compute_region_security_policy#interval_sec ComputeRegionSecurityPolicy#interval_sec}
	IntervalSec *float64 `field:"optional" json:"intervalSec" yaml:"intervalSec"`
}

