package dataprocworkflowtemplate


type DataprocWorkflowTemplateJobsPysparkJob struct {
	// Required. The HCFS URI of the main Python file to use as the driver. Must be a .py file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_workflow_template#main_python_file_uri DataprocWorkflowTemplate#main_python_file_uri}
	MainPythonFileUri *string `field:"required" json:"mainPythonFileUri" yaml:"mainPythonFileUri"`
	// Optional.
	//
	// HCFS URIs of archives to be extracted into the working directory of each executor. Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_workflow_template#archive_uris DataprocWorkflowTemplate#archive_uris}
	ArchiveUris *[]*string `field:"optional" json:"archiveUris" yaml:"archiveUris"`
	// Optional.
	//
	// The arguments to pass to the driver. Do not include arguments, such as `--conf`, that can be set as job properties, since a collision may occur that causes an incorrect job submission.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_workflow_template#args DataprocWorkflowTemplate#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Optional.
	//
	// HCFS URIs of files to be placed in the working directory of each executor. Useful for naively parallel tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_workflow_template#file_uris DataprocWorkflowTemplate#file_uris}
	FileUris *[]*string `field:"optional" json:"fileUris" yaml:"fileUris"`
	// Optional. HCFS URIs of jar files to add to the CLASSPATHs of the Python driver and tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_workflow_template#jar_file_uris DataprocWorkflowTemplate#jar_file_uris}
	JarFileUris *[]*string `field:"optional" json:"jarFileUris" yaml:"jarFileUris"`
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_workflow_template#logging_config DataprocWorkflowTemplate#logging_config}
	LoggingConfig *DataprocWorkflowTemplateJobsPysparkJobLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// Optional.
	//
	// A mapping of property names to values, used to configure PySpark. Properties that conflict with values set by the Dataproc API may be overwritten. Can include properties set in /etc/spark/conf/spark-defaults.conf and classes in user code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_workflow_template#properties DataprocWorkflowTemplate#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// Optional.
	//
	// HCFS file URIs of Python files to pass to the PySpark framework. Supported file types: .py, .egg, and .zip.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_workflow_template#python_file_uris DataprocWorkflowTemplate#python_file_uris}
	PythonFileUris *[]*string `field:"optional" json:"pythonFileUris" yaml:"pythonFileUris"`
}

