// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregionbackendservice


type ComputeRegionBackendServiceCircuitBreakers struct {
	// The maximum number of connections to the backend cluster. Defaults to 1024.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/compute_region_backend_service#max_connections ComputeRegionBackendService#max_connections}
	MaxConnections *float64 `field:"optional" json:"maxConnections" yaml:"maxConnections"`
	// The maximum number of pending requests to the backend cluster. Defaults to 1024.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/compute_region_backend_service#max_pending_requests ComputeRegionBackendService#max_pending_requests}
	MaxPendingRequests *float64 `field:"optional" json:"maxPendingRequests" yaml:"maxPendingRequests"`
	// The maximum number of parallel requests to the backend cluster. Defaults to 1024.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/compute_region_backend_service#max_requests ComputeRegionBackendService#max_requests}
	MaxRequests *float64 `field:"optional" json:"maxRequests" yaml:"maxRequests"`
	// Maximum requests for a single backend connection.
	//
	// This parameter
	// is respected by both the HTTP/1.1 and HTTP/2 implementations. If
	// not specified, there is no limit. Setting this parameter to 1
	// will effectively disable keep alive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/compute_region_backend_service#max_requests_per_connection ComputeRegionBackendService#max_requests_per_connection}
	MaxRequestsPerConnection *float64 `field:"optional" json:"maxRequestsPerConnection" yaml:"maxRequestsPerConnection"`
	// The maximum number of parallel retries to the backend cluster. Defaults to 3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.9.0/docs/resources/compute_region_backend_service#max_retries ComputeRegionBackendService#max_retries}
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
}

