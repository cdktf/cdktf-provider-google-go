// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ArtifactRegistryRepositoryTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/artifact_registry_repository#create ArtifactRegistryRepository#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/artifact_registry_repository#delete ArtifactRegistryRepository#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/artifact_registry_repository#update ArtifactRegistryRepository#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

