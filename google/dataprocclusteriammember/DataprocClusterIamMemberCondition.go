package dataprocclusteriammember


type DataprocClusterIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_cluster_iam_member#expression DataprocClusterIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_cluster_iam_member#title DataprocClusterIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_cluster_iam_member#description DataprocClusterIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

