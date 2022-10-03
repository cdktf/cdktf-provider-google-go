package containeranalysisnote


type ContainerAnalysisNoteTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_analysis_note#create ContainerAnalysisNote#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_analysis_note#delete ContainerAnalysisNote#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_analysis_note#update ContainerAnalysisNote#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

