// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeeapiproduct


type ApigeeApiProductOperationGroupOperationConfigsOperations struct {
	// Methods refers to the REST verbs, when none specified, all verb types are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/apigee_api_product#methods ApigeeApiProduct#methods}
	Methods *[]*string `field:"optional" json:"methods" yaml:"methods"`
	// Required. REST resource path associated with the API proxy or remote service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/apigee_api_product#resource ApigeeApiProduct#resource}
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
}

