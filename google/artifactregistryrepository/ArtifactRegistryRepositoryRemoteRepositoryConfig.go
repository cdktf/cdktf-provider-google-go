package artifactregistryrepository


type ArtifactRegistryRepositoryRemoteRepositoryConfig struct {
	// The description of the remote source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/artifact_registry_repository#description ArtifactRegistryRepository#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// docker_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/artifact_registry_repository#docker_repository ArtifactRegistryRepository#docker_repository}
	DockerRepository *ArtifactRegistryRepositoryRemoteRepositoryConfigDockerRepository `field:"optional" json:"dockerRepository" yaml:"dockerRepository"`
	// maven_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/artifact_registry_repository#maven_repository ArtifactRegistryRepository#maven_repository}
	MavenRepository *ArtifactRegistryRepositoryRemoteRepositoryConfigMavenRepository `field:"optional" json:"mavenRepository" yaml:"mavenRepository"`
	// npm_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/artifact_registry_repository#npm_repository ArtifactRegistryRepository#npm_repository}
	NpmRepository *ArtifactRegistryRepositoryRemoteRepositoryConfigNpmRepository `field:"optional" json:"npmRepository" yaml:"npmRepository"`
	// python_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/artifact_registry_repository#python_repository ArtifactRegistryRepository#python_repository}
	PythonRepository *ArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepository `field:"optional" json:"pythonRepository" yaml:"pythonRepository"`
}

