package artifactregistryrepository


type ArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepository struct {
	// Address of the remote repository. Default value: "PYPI" Possible values: ["PYPI"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/artifact_registry_repository#public_repository ArtifactRegistryRepository#public_repository}
	PublicRepository *string `field:"optional" json:"publicRepository" yaml:"publicRepository"`
}

