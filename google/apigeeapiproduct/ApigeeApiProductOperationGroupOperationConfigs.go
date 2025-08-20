// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeeapiproduct


type ApigeeApiProductOperationGroupOperationConfigs struct {
	// Required. Name of the API proxy or remote service with which the resources, methods, and quota are associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/apigee_api_product#api_source ApigeeApiProduct#api_source}
	ApiSource *string `field:"optional" json:"apiSource" yaml:"apiSource"`
	// attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/apigee_api_product#attributes ApigeeApiProduct#attributes}
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// operations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/apigee_api_product#operations ApigeeApiProduct#operations}
	Operations interface{} `field:"optional" json:"operations" yaml:"operations"`
	// quota block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/apigee_api_product#quota ApigeeApiProduct#quota}
	Quota *ApigeeApiProductOperationGroupOperationConfigsQuota `field:"optional" json:"quota" yaml:"quota"`
}

