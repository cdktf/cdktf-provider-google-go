// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ContainerAnalysisNoteRelatedUrl struct {
	// Specific URL associated with the resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_analysis_note#url ContainerAnalysisNote#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// Label to describe usage of the URL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_analysis_note#label ContainerAnalysisNote#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
}

