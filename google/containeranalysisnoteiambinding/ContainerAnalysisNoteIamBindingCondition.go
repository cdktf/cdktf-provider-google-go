package containeranalysisnoteiambinding


type ContainerAnalysisNoteIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_analysis_note_iam_binding#expression ContainerAnalysisNoteIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_analysis_note_iam_binding#title ContainerAnalysisNoteIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/container_analysis_note_iam_binding#description ContainerAnalysisNoteIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

