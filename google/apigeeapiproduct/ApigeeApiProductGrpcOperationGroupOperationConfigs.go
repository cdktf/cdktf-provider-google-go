// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigeeapiproduct


type ApigeeApiProductGrpcOperationGroupOperationConfigs struct {
	// Required. Name of the API proxy with which the gRPC operation and quota are associated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/apigee_api_product#api_source ApigeeApiProduct#api_source}
	ApiSource *string `field:"optional" json:"apiSource" yaml:"apiSource"`
	// attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/apigee_api_product#attributes ApigeeApiProduct#attributes}
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// List of unqualified gRPC method names for the proxy to which quota will be applied.
	//
	// If this field is empty, the Quota will apply to all operations on the gRPC service defined on the proxy.
	//
	// Example: Given a proxy that is configured to serve com.petstore.PetService, the methods com.petstore.PetService.ListPets and com.petstore.PetService.GetPet would be specified here as simply ["ListPets", "GetPet"].
	//
	// Note: Currently, you can specify only a single GraphQLOperation. Specifying more than one will cause the operation to fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/apigee_api_product#methods ApigeeApiProduct#methods}
	Methods *[]*string `field:"optional" json:"methods" yaml:"methods"`
	// quota block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/apigee_api_product#quota ApigeeApiProduct#quota}
	Quota *ApigeeApiProductGrpcOperationGroupOperationConfigsQuota `field:"optional" json:"quota" yaml:"quota"`
	// Required.
	//
	// gRPC Service name associated to be associated with the API proxy, on which quota rules can be applied upon.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.0/docs/resources/apigee_api_product#service ApigeeApiProduct#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

