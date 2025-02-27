// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudtasksqueue


type CloudTasksQueueHttpTargetUriOverride struct {
	// Host override.
	//
	// When specified, replaces the host part of the task URL.
	// For example, if the task URL is "https://www.google.com", and host value
	// is set to "example.net", the overridden URI will be changed to "https://example.net".
	// Host value cannot be an empty string (INVALID_ARGUMENT).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/cloud_tasks_queue#host CloudTasksQueue#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// path_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/cloud_tasks_queue#path_override CloudTasksQueue#path_override}
	PathOverride *CloudTasksQueueHttpTargetUriOverridePathOverride `field:"optional" json:"pathOverride" yaml:"pathOverride"`
	// Port override.
	//
	// When specified, replaces the port part of the task URI.
	// For instance, for a URI http://www.google.com/foo and port=123, the overridden URI becomes http://www.google.com:123/foo.
	// Note that the port value must be a positive integer.
	// Setting the port to 0 (Zero) clears the URI port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/cloud_tasks_queue#port CloudTasksQueue#port}
	Port *string `field:"optional" json:"port" yaml:"port"`
	// query_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/cloud_tasks_queue#query_override CloudTasksQueue#query_override}
	QueryOverride *CloudTasksQueueHttpTargetUriOverrideQueryOverride `field:"optional" json:"queryOverride" yaml:"queryOverride"`
	// Scheme override.
	//
	// When specified, the task URI scheme is replaced by the provided value (HTTP or HTTPS). Possible values: ["HTTP", "HTTPS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/cloud_tasks_queue#scheme CloudTasksQueue#scheme}
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
	// URI Override Enforce Mode.
	//
	// When specified, determines the Target UriOverride mode. If not specified, it defaults to ALWAYS. Possible values: ["ALWAYS", "IF_NOT_EXISTS"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.23.0/docs/resources/cloud_tasks_queue#uri_override_enforce_mode CloudTasksQueue#uri_override_enforce_mode}
	UriOverrideEnforceMode *string `field:"optional" json:"uriOverrideEnforceMode" yaml:"uriOverrideEnforceMode"`
}

