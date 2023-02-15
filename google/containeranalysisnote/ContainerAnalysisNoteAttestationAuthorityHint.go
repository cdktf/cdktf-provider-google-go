package containeranalysisnote


type ContainerAnalysisNoteAttestationAuthorityHint struct {
	// The human readable name of this Attestation Authority, for example "qa".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_analysis_note#human_readable_name ContainerAnalysisNote#human_readable_name}
	HumanReadableName *string `field:"required" json:"humanReadableName" yaml:"humanReadableName"`
}

