package dataprocworkflowtemplate


type DataprocWorkflowTemplateJobsSparkSqlJob struct {
	// Optional. HCFS URIs of jar files to be added to the Spark CLASSPATH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataproc_workflow_template#jar_file_uris DataprocWorkflowTemplate#jar_file_uris}
	JarFileUris *[]*string `field:"optional" json:"jarFileUris" yaml:"jarFileUris"`
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataproc_workflow_template#logging_config DataprocWorkflowTemplate#logging_config}
	LoggingConfig *DataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// Optional.
	//
	// A mapping of property names to values, used to configure Spark SQL's SparkConf. Properties that conflict with values set by the Dataproc API may be overwritten.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataproc_workflow_template#properties DataprocWorkflowTemplate#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// The HCFS URI of the script that contains SQL queries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataproc_workflow_template#query_file_uri DataprocWorkflowTemplate#query_file_uri}
	QueryFileUri *string `field:"optional" json:"queryFileUri" yaml:"queryFileUri"`
	// query_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataproc_workflow_template#query_list DataprocWorkflowTemplate#query_list}
	QueryList *DataprocWorkflowTemplateJobsSparkSqlJobQueryListStruct `field:"optional" json:"queryList" yaml:"queryList"`
	// Optional. Mapping of query variable names to values (equivalent to the Spark SQL command: SET `name="value";`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/dataproc_workflow_template#script_variables DataprocWorkflowTemplate#script_variables}
	ScriptVariables *map[string]*string `field:"optional" json:"scriptVariables" yaml:"scriptVariables"`
}

