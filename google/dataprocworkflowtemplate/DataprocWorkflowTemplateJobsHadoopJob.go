package dataprocworkflowtemplate


type DataprocWorkflowTemplateJobsHadoopJob struct {
	// Optional.
	//
	// HCFS URIs of archives to be extracted in the working directory of Hadoop drivers and tasks. Supported file types: .jar, .tar, .tar.gz, .tgz, or .zip.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_workflow_template#archive_uris DataprocWorkflowTemplate#archive_uris}
	ArchiveUris *[]*string `field:"optional" json:"archiveUris" yaml:"archiveUris"`
	// Optional.
	//
	// The arguments to pass to the driver. Do not include arguments, such as `-libjars` or `-Dfoo=bar`, that can be set as job properties, since a collision may occur that causes an incorrect job submission.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_workflow_template#args DataprocWorkflowTemplate#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Optional.
	//
	// HCFS (Hadoop Compatible Filesystem) URIs of files to be copied to the working directory of Hadoop drivers and distributed tasks. Useful for naively parallel tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_workflow_template#file_uris DataprocWorkflowTemplate#file_uris}
	FileUris *[]*string `field:"optional" json:"fileUris" yaml:"fileUris"`
	// Optional. Jar file URIs to add to the CLASSPATHs of the Hadoop driver and tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_workflow_template#jar_file_uris DataprocWorkflowTemplate#jar_file_uris}
	JarFileUris *[]*string `field:"optional" json:"jarFileUris" yaml:"jarFileUris"`
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_workflow_template#logging_config DataprocWorkflowTemplate#logging_config}
	LoggingConfig *DataprocWorkflowTemplateJobsHadoopJobLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// The name of the driver's main class.
	//
	// The jar file containing the class must be in the default CLASSPATH or specified in `jar_file_uris`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_workflow_template#main_class DataprocWorkflowTemplate#main_class}
	MainClass *string `field:"optional" json:"mainClass" yaml:"mainClass"`
	// The HCFS URI of the jar file containing the main class. Examples: 'gs://foo-bucket/analytics-binaries/extract-useful-metrics-mr.jar' 'hdfs:/tmp/test-samples/custom-wordcount.jar' 'file:///home/usr/lib/hadoop-mapreduce/hadoop-mapreduce-examples.jar'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_workflow_template#main_jar_file_uri DataprocWorkflowTemplate#main_jar_file_uri}
	MainJarFileUri *string `field:"optional" json:"mainJarFileUri" yaml:"mainJarFileUri"`
	// Optional.
	//
	// A mapping of property names to values, used to configure Hadoop. Properties that conflict with values set by the Dataproc API may be overwritten. Can include properties set in /etc/hadoop/conf/*-site and classes in user code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/dataproc_workflow_template#properties DataprocWorkflowTemplate#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
}

