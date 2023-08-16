package dataplextask


type DataplexTaskNotebook struct {
	// Path to input notebook.
	//
	// This can be the Cloud Storage URI of the notebook file or the path to a Notebook Content. The execution args are accessible as environment variables (TASK_key=value).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_task#notebook DataplexTask#notebook}
	Notebook *string `field:"required" json:"notebook" yaml:"notebook"`
	// Cloud Storage URIs of archives to be extracted into the working directory of each executor.
	//
	// Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_task#archive_uris DataplexTask#archive_uris}
	ArchiveUris *[]*string `field:"optional" json:"archiveUris" yaml:"archiveUris"`
	// Cloud Storage URIs of files to be placed in the working directory of each executor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_task#file_uris DataplexTask#file_uris}
	FileUris *[]*string `field:"optional" json:"fileUris" yaml:"fileUris"`
	// infrastructure_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_task#infrastructure_spec DataplexTask#infrastructure_spec}
	InfrastructureSpec *DataplexTaskNotebookInfrastructureSpec `field:"optional" json:"infrastructureSpec" yaml:"infrastructureSpec"`
}

