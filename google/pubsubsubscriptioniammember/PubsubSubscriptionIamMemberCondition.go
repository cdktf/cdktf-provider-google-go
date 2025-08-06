// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pubsubsubscriptioniammember


type PubsubSubscriptionIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/pubsub_subscription_iam_member#expression PubsubSubscriptionIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/pubsub_subscription_iam_member#title PubsubSubscriptionIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.47.0/docs/resources/pubsub_subscription_iam_member#description PubsubSubscriptionIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

