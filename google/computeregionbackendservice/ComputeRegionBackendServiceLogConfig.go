// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionbackendservice


type ComputeRegionBackendServiceLogConfig struct {
	// Whether to enable logging for the load balancer traffic served by this backend service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_region_backend_service#enable ComputeRegionBackendService#enable}
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// Specifies the fields to include in logging.
	//
	// This field can only be specified if logging is enabled for this backend service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_region_backend_service#optional_fields ComputeRegionBackendService#optional_fields}
	OptionalFields *[]*string `field:"optional" json:"optionalFields" yaml:"optionalFields"`
	// Specifies the optional logging mode for the load balancer traffic. Supported values: INCLUDE_ALL_OPTIONAL, EXCLUDE_ALL_OPTIONAL, CUSTOM. Possible values: ["INCLUDE_ALL_OPTIONAL", "EXCLUDE_ALL_OPTIONAL", "CUSTOM"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_region_backend_service#optional_mode ComputeRegionBackendService#optional_mode}
	OptionalMode *string `field:"optional" json:"optionalMode" yaml:"optionalMode"`
	// This field can only be specified if logging is enabled for this backend service.
	//
	// The value of
	// the field must be in [0, 1]. This configures the sampling rate of requests to the load balancer
	// where 1.0 means all logged requests are reported and 0.0 means no logged requests are reported.
	// The default value is 1.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.26.0/docs/resources/compute_region_backend_service#sample_rate ComputeRegionBackendService#sample_rate}
	SampleRate *float64 `field:"optional" json:"sampleRate" yaml:"sampleRate"`
}

