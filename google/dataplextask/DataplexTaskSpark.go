package dataplextask


type DataplexTaskSpark struct {
	// Cloud Storage URIs of archives to be extracted into the working directory of each executor.
	//
	// Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataplex_task#archive_uris DataplexTask#archive_uris}
	ArchiveUris *[]*string `field:"optional" json:"archiveUris" yaml:"archiveUris"`
	// Cloud Storage URIs of files to be placed in the working directory of each executor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataplex_task#file_uris DataplexTask#file_uris}
	FileUris *[]*string `field:"optional" json:"fileUris" yaml:"fileUris"`
	// infrastructure_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataplex_task#infrastructure_spec DataplexTask#infrastructure_spec}
	InfrastructureSpec *DataplexTaskSparkInfrastructureSpec `field:"optional" json:"infrastructureSpec" yaml:"infrastructureSpec"`
	// The name of the driver's main class.
	//
	// The jar file that contains the class must be in the default CLASSPATH or specified in jar_file_uris. The execution args are passed in as a sequence of named process arguments (--key=value).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataplex_task#main_class DataplexTask#main_class}
	MainClass *string `field:"optional" json:"mainClass" yaml:"mainClass"`
	// The Cloud Storage URI of the jar file that contains the main class.
	//
	// The execution args are passed in as a sequence of named process arguments (--key=value).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataplex_task#main_jar_file_uri DataplexTask#main_jar_file_uri}
	MainJarFileUri *string `field:"optional" json:"mainJarFileUri" yaml:"mainJarFileUri"`
	// The Gcloud Storage URI of the main Python file to use as the driver.
	//
	// Must be a .py file. The execution args are passed in as a sequence of named process arguments (--key=value).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataplex_task#python_script_file DataplexTask#python_script_file}
	PythonScriptFile *string `field:"optional" json:"pythonScriptFile" yaml:"pythonScriptFile"`
	// The query text. The execution args are used to declare a set of script variables (set key='value';).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataplex_task#sql_script DataplexTask#sql_script}
	SqlScript *string `field:"optional" json:"sqlScript" yaml:"sqlScript"`
	// A reference to a query file.
	//
	// This can be the Cloud Storage URI of the query file or it can the path to a SqlScript Content. The execution args are used to declare a set of script variables (set key='value';).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataplex_task#sql_script_file DataplexTask#sql_script_file}
	SqlScriptFile *string `field:"optional" json:"sqlScriptFile" yaml:"sqlScriptFile"`
}

