// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeregiontargethttpproxy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeRegionTargetHttpProxyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Name of the resource.
	//
	// Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.3.0/docs/resources/compute_region_target_http_proxy#name ComputeRegionTargetHttpProxy#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A reference to the RegionUrlMap resource that defines the mapping from URL to the BackendService.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.3.0/docs/resources/compute_region_target_http_proxy#url_map ComputeRegionTargetHttpProxy#url_map}
	UrlMap *string `field:"required" json:"urlMap" yaml:"urlMap"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.3.0/docs/resources/compute_region_target_http_proxy#description ComputeRegionTargetHttpProxy#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies how long to keep a connection open, after completing a response, while there is no matching traffic (in seconds).
	//
	// If an HTTP keepalive is
	// not specified, a default value (600 seconds) will be used. For Regional
	// HTTP(S) load balancer, the minimum allowed value is 5 seconds and the
	// maximum allowed value is 600 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.3.0/docs/resources/compute_region_target_http_proxy#http_keep_alive_timeout_sec ComputeRegionTargetHttpProxy#http_keep_alive_timeout_sec}
	HttpKeepAliveTimeoutSec *float64 `field:"optional" json:"httpKeepAliveTimeoutSec" yaml:"httpKeepAliveTimeoutSec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.3.0/docs/resources/compute_region_target_http_proxy#id ComputeRegionTargetHttpProxy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.3.0/docs/resources/compute_region_target_http_proxy#project ComputeRegionTargetHttpProxy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The Region in which the created target https proxy should reside.
	//
	// If it is not provided, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.3.0/docs/resources/compute_region_target_http_proxy#region ComputeRegionTargetHttpProxy#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.3.0/docs/resources/compute_region_target_http_proxy#timeouts ComputeRegionTargetHttpProxy#timeouts}
	Timeouts *ComputeRegionTargetHttpProxyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

