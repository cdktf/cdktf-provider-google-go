package artifactregistryrepository


type ArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepository struct {
	// Address of the remote repository. Default value: "MAVEN_CENTRAL" Possible values: ["MAVEN_CENTRAL"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/artifact_registry_repository#public_repository ArtifactRegistryRepository#public_repository}
	PublicRepository *string `field:"optional" json:"publicRepository" yaml:"publicRepository"`
}

