// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pubsubsubscription


type PubsubSubscriptionMessageTransformsJavascriptUdf struct {
	// JavaScript code that contains a function 'function_name' with the following signature: ```   /**   * Transforms a Pub/Sub message.
	//
	// *
	// *.
	//
	// Returns: - To
	// * filter a message, return 'null'. To transform a message return a map
	// * with the following keys:
	// *   - (required) 'data' : {string}
	// *   - (optional) 'attributes' : {Object<string, string>}
	// * Returning empty 'attributes' will remove all attributes from the
	// * message.
	// *
	// *.
	Code *string `field:"required" json:"code" yaml:"code"`
	// Name of the JavaScript function that should be applied to Pub/Sub messages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.40.0/docs/resources/pubsub_subscription#function_name PubsubSubscription#function_name}
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
}

