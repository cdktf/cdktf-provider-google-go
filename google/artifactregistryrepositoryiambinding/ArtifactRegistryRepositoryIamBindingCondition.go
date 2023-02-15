package artifactregistryrepositoryiambinding


type ArtifactRegistryRepositoryIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/artifact_registry_repository_iam_binding#expression ArtifactRegistryRepositoryIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/artifact_registry_repository_iam_binding#title ArtifactRegistryRepositoryIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/artifact_registry_repository_iam_binding#description ArtifactRegistryRepositoryIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

