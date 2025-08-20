// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeeapiproduct


type ApigeeApiProductGraphqlOperationGroupOperationConfigsOperations struct {
	// GraphQL operation name.
	//
	// The name and operation type will be used to apply quotas. If no name is specified, the quota will be applied to all GraphQL operations irrespective of their operation names in the payload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/apigee_api_product#operation ApigeeApiProduct#operation}
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
	// Required. GraphQL operation types. Valid values include query or mutation. Note: Apigee does not currently support subscription types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/apigee_api_product#operation_types ApigeeApiProduct#operation_types}
	OperationTypes *[]*string `field:"optional" json:"operationTypes" yaml:"operationTypes"`
}

