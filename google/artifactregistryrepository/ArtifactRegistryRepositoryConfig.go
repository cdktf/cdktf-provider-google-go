package artifactregistryrepository

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ArtifactRegistryRepositoryConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The format of packages that are stored in the repository.
	//
	// Supported formats
	// can be found [here](https://cloud.google.com/artifact-registry/docs/supported-formats).
	// You can only create alpha formats if you are a member of the
	// [alpha user group](https://cloud.google.com/artifact-registry/docs/supported-formats#alpha-access).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/artifact_registry_repository#format ArtifactRegistryRepository#format}
	Format *string `field:"required" json:"format" yaml:"format"`
	// The last part of the repository name, for example: "repo1".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/artifact_registry_repository#repository_id ArtifactRegistryRepository#repository_id}
	RepositoryId *string `field:"required" json:"repositoryId" yaml:"repositoryId"`
	// The user-provided description of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/artifact_registry_repository#description ArtifactRegistryRepository#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// docker_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/artifact_registry_repository#docker_config ArtifactRegistryRepository#docker_config}
	DockerConfig *ArtifactRegistryRepositoryDockerConfig `field:"optional" json:"dockerConfig" yaml:"dockerConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/artifact_registry_repository#id ArtifactRegistryRepository#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The Cloud KMS resource name of the customer managed encryption key that’s used to encrypt the contents of the Repository.
	//
	// Has the form:
	// 'projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key'.
	// This value may not be changed after the Repository has been created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/artifact_registry_repository#kms_key_name ArtifactRegistryRepository#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// Labels with user-defined metadata.
	//
	// This field may contain up to 64 entries. Label keys and values may be no
	// longer than 63 characters. Label keys must begin with a lowercase letter
	// and may only contain lowercase letters, numeric characters, underscores,
	// and dashes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/artifact_registry_repository#labels ArtifactRegistryRepository#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The name of the location this repository is located in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/artifact_registry_repository#location ArtifactRegistryRepository#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// maven_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/artifact_registry_repository#maven_config ArtifactRegistryRepository#maven_config}
	MavenConfig *ArtifactRegistryRepositoryMavenConfig `field:"optional" json:"mavenConfig" yaml:"mavenConfig"`
	// The mode configures the repository to serve artifacts from different sources. Default value: "STANDARD_REPOSITORY" Possible values: ["STANDARD_REPOSITORY", "VIRTUAL_REPOSITORY", "REMOTE_REPOSITORY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/artifact_registry_repository#mode ArtifactRegistryRepository#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/artifact_registry_repository#project ArtifactRegistryRepository#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// remote_repository_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/artifact_registry_repository#remote_repository_config ArtifactRegistryRepository#remote_repository_config}
	RemoteRepositoryConfig *ArtifactRegistryRepositoryRemoteRepositoryConfig `field:"optional" json:"remoteRepositoryConfig" yaml:"remoteRepositoryConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/artifact_registry_repository#timeouts ArtifactRegistryRepository#timeouts}
	Timeouts *ArtifactRegistryRepositoryTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// virtual_repository_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/artifact_registry_repository#virtual_repository_config ArtifactRegistryRepository#virtual_repository_config}
	VirtualRepositoryConfig *ArtifactRegistryRepositoryVirtualRepositoryConfig `field:"optional" json:"virtualRepositoryConfig" yaml:"virtualRepositoryConfig"`
}

