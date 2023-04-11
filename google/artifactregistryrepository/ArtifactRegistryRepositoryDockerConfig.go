package artifactregistryrepository


type ArtifactRegistryRepositoryDockerConfig struct {
	// The repository which enabled this flag prevents all tags from being modified, moved or deleted.
	//
	// This does not prevent tags from being created.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/artifact_registry_repository#immutable_tags ArtifactRegistryRepository#immutable_tags}
	ImmutableTags interface{} `field:"optional" json:"immutableTags" yaml:"immutableTags"`
}

