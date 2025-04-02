// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventarcpipeline


type EventarcPipelineDestinationsOutputPayloadFormat struct {
	// avro block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/eventarc_pipeline#avro EventarcPipeline#avro}
	Avro *EventarcPipelineDestinationsOutputPayloadFormatAvro `field:"optional" json:"avro" yaml:"avro"`
	// json block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/eventarc_pipeline#json EventarcPipeline#json}
	Json *EventarcPipelineDestinationsOutputPayloadFormatJson `field:"optional" json:"json" yaml:"json"`
	// protobuf block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.28.0/docs/resources/eventarc_pipeline#protobuf EventarcPipeline#protobuf}
	Protobuf *EventarcPipelineDestinationsOutputPayloadFormatProtobuf `field:"optional" json:"protobuf" yaml:"protobuf"`
}

