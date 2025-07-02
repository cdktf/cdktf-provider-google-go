// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventarcpipeline


type EventarcPipelineInputPayloadFormat struct {
	// avro block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/eventarc_pipeline#avro EventarcPipeline#avro}
	Avro *EventarcPipelineInputPayloadFormatAvro `field:"optional" json:"avro" yaml:"avro"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/eventarc_pipeline#json EventarcPipeline#json}
	Json *EventarcPipelineInputPayloadFormatJson `field:"optional" json:"json" yaml:"json"`
	// protobuf block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.42.0/docs/resources/eventarc_pipeline#protobuf EventarcPipeline#protobuf}
	Protobuf *EventarcPipelineInputPayloadFormatProtobuf `field:"optional" json:"protobuf" yaml:"protobuf"`
}

