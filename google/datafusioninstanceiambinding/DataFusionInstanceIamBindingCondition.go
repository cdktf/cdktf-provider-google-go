package datafusioninstanceiambinding


type DataFusionInstanceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_fusion_instance_iam_binding#expression DataFusionInstanceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_fusion_instance_iam_binding#title DataFusionInstanceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_fusion_instance_iam_binding#description DataFusionInstanceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

