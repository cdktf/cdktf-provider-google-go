// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type DataprocClusterIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_cluster_iam_binding#expression DataprocClusterIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_cluster_iam_binding#title DataprocClusterIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_cluster_iam_binding#description DataprocClusterIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

