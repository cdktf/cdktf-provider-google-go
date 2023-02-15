package containeranalysisnote


type ContainerAnalysisNoteAttestationAuthority struct {
	// hint block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_analysis_note#hint ContainerAnalysisNote#hint}
	Hint *ContainerAnalysisNoteAttestationAuthorityHint `field:"required" json:"hint" yaml:"hint"`
}

