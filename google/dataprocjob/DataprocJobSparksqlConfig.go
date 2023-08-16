package dataprocjob


type DataprocJobSparksqlConfig struct {
	// HCFS URIs of jar files to be added to the Spark CLASSPATH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_job#jar_file_uris DataprocJob#jar_file_uris}
	JarFileUris *[]*string `field:"optional" json:"jarFileUris" yaml:"jarFileUris"`
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_job#logging_config DataprocJob#logging_config}
	LoggingConfig *DataprocJobSparksqlConfigLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// A mapping of property names to values, used to configure Spark SQL's SparkConf.
	//
	// Properties that conflict with values set by the Cloud Dataproc API may be overwritten.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_job#properties DataprocJob#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// The HCFS URI of the script that contains SQL queries. Conflicts with query_list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_job#query_file_uri DataprocJob#query_file_uri}
	QueryFileUri *string `field:"optional" json:"queryFileUri" yaml:"queryFileUri"`
	// The list of SQL queries or statements to execute as part of the job. Conflicts with query_file_uri.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_job#query_list DataprocJob#query_list}
	QueryList *[]*string `field:"optional" json:"queryList" yaml:"queryList"`
	// Mapping of query variable names to values (equivalent to the Spark SQL command: SET name="value";).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataproc_job#script_variables DataprocJob#script_variables}
	ScriptVariables *map[string]*string `field:"optional" json:"scriptVariables" yaml:"scriptVariables"`
}

