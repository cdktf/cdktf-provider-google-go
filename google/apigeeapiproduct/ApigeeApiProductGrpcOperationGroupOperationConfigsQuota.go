// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeeapiproduct


type ApigeeApiProductGrpcOperationGroupOperationConfigsQuota struct {
	// Required. Time interval over which the number of request messages is calculated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/apigee_api_product#interval ApigeeApiProduct#interval}
	Interval *string `field:"optional" json:"interval" yaml:"interval"`
	// Required. Upper limit allowed for the time interval and time unit specified. Requests exceeding this limit will be rejected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/apigee_api_product#limit ApigeeApiProduct#limit}
	Limit *string `field:"optional" json:"limit" yaml:"limit"`
	// Time unit defined for the interval.
	//
	// Valid values include second, minute, hour, day, month or year. If limit and interval are valid, the default value is hour; otherwise, the default is null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.2/docs/resources/apigee_api_product#time_unit ApigeeApiProduct#time_unit}
	TimeUnit *string `field:"optional" json:"timeUnit" yaml:"timeUnit"`
}

