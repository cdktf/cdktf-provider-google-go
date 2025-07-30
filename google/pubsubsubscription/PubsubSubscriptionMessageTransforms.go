// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pubsubsubscription


type PubsubSubscriptionMessageTransforms struct {
	// Controls whether or not to use this transform.
	//
	// If not set or 'false',
	// the transform will be applied to messages. Default: 'true'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/pubsub_subscription#disabled PubsubSubscription#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// javascript_udf block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.46.0/docs/resources/pubsub_subscription#javascript_udf PubsubSubscription#javascript_udf}
	JavascriptUdf *PubsubSubscriptionMessageTransformsJavascriptUdf `field:"optional" json:"javascriptUdf" yaml:"javascriptUdf"`
}

