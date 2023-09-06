// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containeranalysisnote


type ContainerAnalysisNoteAttestationAuthorityHint struct {
	// The human readable name of this Attestation Authority, for example "qa".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.81.0/docs/resources/container_analysis_note#human_readable_name ContainerAnalysisNote#human_readable_name}
	HumanReadableName *string `field:"required" json:"humanReadableName" yaml:"humanReadableName"`
}

