package dataprocworkflowtemplate


type DataprocWorkflowTemplateJobsHiveJob struct {
	// Optional.
	//
	// Whether to continue executing queries if a query fails. The default value is `false`. Setting to `true` can be useful when executing independent parallel queries.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_workflow_template#continue_on_failure DataprocWorkflowTemplate#continue_on_failure}
	ContinueOnFailure interface{} `field:"optional" json:"continueOnFailure" yaml:"continueOnFailure"`
	// Optional.
	//
	// HCFS URIs of jar files to add to the CLASSPATH of the Hive server and Hadoop MapReduce (MR) tasks. Can contain Hive SerDes and UDFs.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_workflow_template#jar_file_uris DataprocWorkflowTemplate#jar_file_uris}
	JarFileUris *[]*string `field:"optional" json:"jarFileUris" yaml:"jarFileUris"`
	// Optional.
	//
	// A mapping of property names and values, used to configure Hive. Properties that conflict with values set by the Dataproc API may be overwritten. Can include properties set in /etc/hadoop/conf/*-site.xml, /etc/hive/conf/hive-site.xml, and classes in user code.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_workflow_template#properties DataprocWorkflowTemplate#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// The HCFS URI of the script that contains Hive queries.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_workflow_template#query_file_uri DataprocWorkflowTemplate#query_file_uri}
	QueryFileUri *string `field:"optional" json:"queryFileUri" yaml:"queryFileUri"`
	// query_list block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_workflow_template#query_list DataprocWorkflowTemplate#query_list}
	QueryList *DataprocWorkflowTemplateJobsHiveJobQueryList `field:"optional" json:"queryList" yaml:"queryList"`
	// Optional. Mapping of query variable names to values (equivalent to the Hive command: `SET name="value";`).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_workflow_template#script_variables DataprocWorkflowTemplate#script_variables}
	ScriptVariables *map[string]*string `field:"optional" json:"scriptVariables" yaml:"scriptVariables"`
}

