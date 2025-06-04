// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionbackendservice


type ComputeRegionBackendServiceBackendCustomMetrics struct {
	// If true, the metric data is collected and reported to Cloud Monitoring, but is not used for load balancing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/compute_region_backend_service#dry_run ComputeRegionBackendService#dry_run}
	DryRun interface{} `field:"required" json:"dryRun" yaml:"dryRun"`
	// Name of a custom utilization signal.
	//
	// The name must be 1-64 characters
	// long and match the regular expression [a-z]([-_.a-z0-9]*[a-z0-9])? which
	// means the first character must be a lowercase letter, and all following
	// characters must be a dash, period, underscore, lowercase letter, or
	// digit, except the last character, which cannot be a dash, period, or
	// underscore. For usage guidelines, see Custom Metrics balancing mode. This
	// field can only be used for a global or regional backend service with the
	// loadBalancingScheme set to <code>EXTERNAL_MANAGED</code>,
	// <code>INTERNAL_MANAGED</code> <code>INTERNAL_SELF_MANAGED</code>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/compute_region_backend_service#name ComputeRegionBackendService#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional parameter to define a target utilization for the Custom Metrics balancing mode. The valid range is <code>[0.0, 1.0]</code>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.38.0/docs/resources/compute_region_backend_service#max_utilization ComputeRegionBackendService#max_utilization}
	MaxUtilization *float64 `field:"optional" json:"maxUtilization" yaml:"maxUtilization"`
}

