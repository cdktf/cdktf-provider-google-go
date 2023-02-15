package datafusioninstanceiammember


type DataFusionInstanceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_fusion_instance_iam_member#expression DataFusionInstanceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_fusion_instance_iam_member#title DataFusionInstanceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/data_fusion_instance_iam_member#description DataFusionInstanceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

