// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeeinstance


type ApigeeInstanceAccessLoggingConfig struct {
	// Boolean flag that specifies whether the customer access log feature is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/apigee_instance#enabled ApigeeInstance#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Ship the access log entries that match the statusCode defined in the filter.
	//
	// The statusCode is the only expected/supported filter field. (Ex: statusCode)
	// The filter will parse it to the Common Expression Language semantics for expression
	// evaluation to build the filter condition. (Ex: "filter": statusCode >= 200 && statusCode < 300 )
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.48.0/docs/resources/apigee_instance#filter ApigeeInstance#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
}

