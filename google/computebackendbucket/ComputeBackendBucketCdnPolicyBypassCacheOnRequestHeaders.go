// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computebackendbucket


type ComputeBackendBucketCdnPolicyBypassCacheOnRequestHeaders struct {
	// The header field name to match on when bypassing cache. Values are case-insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.12.0/docs/resources/compute_backend_bucket#header_name ComputeBackendBucket#header_name}
	HeaderName *string `field:"optional" json:"headerName" yaml:"headerName"`
}

