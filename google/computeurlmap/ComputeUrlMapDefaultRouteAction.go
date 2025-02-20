// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeurlmap


type ComputeUrlMapDefaultRouteAction struct {
	// cors_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/compute_url_map#cors_policy ComputeUrlMap#cors_policy}
	CorsPolicy *ComputeUrlMapDefaultRouteActionCorsPolicy `field:"optional" json:"corsPolicy" yaml:"corsPolicy"`
	// fault_injection_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/compute_url_map#fault_injection_policy ComputeUrlMap#fault_injection_policy}
	FaultInjectionPolicy *ComputeUrlMapDefaultRouteActionFaultInjectionPolicy `field:"optional" json:"faultInjectionPolicy" yaml:"faultInjectionPolicy"`
	// request_mirror_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/compute_url_map#request_mirror_policy ComputeUrlMap#request_mirror_policy}
	RequestMirrorPolicy *ComputeUrlMapDefaultRouteActionRequestMirrorPolicy `field:"optional" json:"requestMirrorPolicy" yaml:"requestMirrorPolicy"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/compute_url_map#retry_policy ComputeUrlMap#retry_policy}
	RetryPolicy *ComputeUrlMapDefaultRouteActionRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/compute_url_map#timeout ComputeUrlMap#timeout}
	Timeout *ComputeUrlMapDefaultRouteActionTimeout `field:"optional" json:"timeout" yaml:"timeout"`
	// url_rewrite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/compute_url_map#url_rewrite ComputeUrlMap#url_rewrite}
	UrlRewrite *ComputeUrlMapDefaultRouteActionUrlRewrite `field:"optional" json:"urlRewrite" yaml:"urlRewrite"`
	// weighted_backend_services block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/compute_url_map#weighted_backend_services ComputeUrlMap#weighted_backend_services}
	WeightedBackendServices interface{} `field:"optional" json:"weightedBackendServices" yaml:"weightedBackendServices"`
}

