package artifactregistryrepository


type ArtifactRegistryRepositoryTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/artifact_registry_repository#create ArtifactRegistryRepository#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/artifact_registry_repository#delete ArtifactRegistryRepository#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/artifact_registry_repository#update ArtifactRegistryRepository#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

