package artifactregistryrepositoryiammember


type ArtifactRegistryRepositoryIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/artifact_registry_repository_iam_member#expression ArtifactRegistryRepositoryIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/artifact_registry_repository_iam_member#title ArtifactRegistryRepositoryIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/artifact_registry_repository_iam_member#description ArtifactRegistryRepositoryIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

