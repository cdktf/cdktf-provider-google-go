// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pubsubschemaiambinding


type PubsubSchemaIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/pubsub_schema_iam_binding#expression PubsubSchemaIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/pubsub_schema_iam_binding#title PubsubSchemaIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.21.0/docs/resources/pubsub_schema_iam_binding#description PubsubSchemaIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

