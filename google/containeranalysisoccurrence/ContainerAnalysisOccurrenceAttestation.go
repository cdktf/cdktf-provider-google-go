// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containeranalysisoccurrence


type ContainerAnalysisOccurrenceAttestation struct {
	// The serialized payload that is verified by one or more signatures. A base64-encoded string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.81.0/docs/resources/container_analysis_occurrence#serialized_payload ContainerAnalysisOccurrence#serialized_payload}
	SerializedPayload *string `field:"required" json:"serializedPayload" yaml:"serializedPayload"`
	// signatures block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.81.0/docs/resources/container_analysis_occurrence#signatures ContainerAnalysisOccurrence#signatures}
	Signatures interface{} `field:"required" json:"signatures" yaml:"signatures"`
}

