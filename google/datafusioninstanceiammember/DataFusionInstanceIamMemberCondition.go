package datafusioninstanceiammember


type DataFusionInstanceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_fusion_instance_iam_member#expression DataFusionInstanceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_fusion_instance_iam_member#title DataFusionInstanceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/data_fusion_instance_iam_member#description DataFusionInstanceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

