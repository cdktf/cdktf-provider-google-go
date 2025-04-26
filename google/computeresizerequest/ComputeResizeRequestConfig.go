// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeresizerequest

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeResizeRequestConfig struct {
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
	// The reference of the instance group manager this ResizeRequest is a part of.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/compute_resize_request#instance_group_manager ComputeResizeRequest#instance_group_manager}
	InstanceGroupManager *string `field:"required" json:"instanceGroupManager" yaml:"instanceGroupManager"`
	// The name of this resize request. The name must be 1-63 characters long, and comply with RFC1035.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/compute_resize_request#name ComputeResizeRequest#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The number of instances to be created by this resize request.
	//
	// The group's target size will be increased by this number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/compute_resize_request#resize_by ComputeResizeRequest#resize_by}
	ResizeBy *float64 `field:"required" json:"resizeBy" yaml:"resizeBy"`
	// An optional description of this resize-request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/compute_resize_request#description ComputeResizeRequest#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/compute_resize_request#id ComputeResizeRequest#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/compute_resize_request#project ComputeResizeRequest#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// requested_run_duration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/compute_resize_request#requested_run_duration ComputeResizeRequest#requested_run_duration}
	RequestedRunDuration *ComputeResizeRequestRequestedRunDuration `field:"optional" json:"requestedRunDuration" yaml:"requestedRunDuration"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/compute_resize_request#timeouts ComputeResizeRequest#timeouts}
	Timeouts *ComputeResizeRequestTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The reference of the compute zone scoping this request. If it is not provided, the provider zone is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.32.0/docs/resources/compute_resize_request#zone ComputeResizeRequest#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

