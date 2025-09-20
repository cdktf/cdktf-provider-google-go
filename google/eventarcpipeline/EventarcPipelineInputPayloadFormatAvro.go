// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventarcpipeline


type EventarcPipelineInputPayloadFormatAvro struct {
	// The entire schema definition is stored in this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/eventarc_pipeline#schema_definition EventarcPipeline#schema_definition}
	SchemaDefinition *string `field:"optional" json:"schemaDefinition" yaml:"schemaDefinition"`
}

