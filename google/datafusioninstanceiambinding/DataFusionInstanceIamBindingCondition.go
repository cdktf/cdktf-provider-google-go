package datafusioninstanceiambinding


type DataFusionInstanceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_fusion_instance_iam_binding#expression DataFusionInstanceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_fusion_instance_iam_binding#title DataFusionInstanceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/data_fusion_instance_iam_binding#description DataFusionInstanceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

