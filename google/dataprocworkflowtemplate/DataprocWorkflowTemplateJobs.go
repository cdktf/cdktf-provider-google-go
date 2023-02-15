package dataprocworkflowtemplate


type DataprocWorkflowTemplateJobs struct {
	// Required.
	//
	// The step id. The id must be unique among all jobs within the template. The step id is used as prefix for job id, as job `goog-dataproc-workflow-step-id` label, and in prerequisiteStepIds field from other steps. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_workflow_template#step_id DataprocWorkflowTemplate#step_id}
	StepId *string `field:"required" json:"stepId" yaml:"stepId"`
	// hadoop_job block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_workflow_template#hadoop_job DataprocWorkflowTemplate#hadoop_job}
	HadoopJob *DataprocWorkflowTemplateJobsHadoopJob `field:"optional" json:"hadoopJob" yaml:"hadoopJob"`
	// hive_job block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_workflow_template#hive_job DataprocWorkflowTemplate#hive_job}
	HiveJob *DataprocWorkflowTemplateJobsHiveJob `field:"optional" json:"hiveJob" yaml:"hiveJob"`
	// Optional.
	//
	// The labels to associate with this job. Label keys must be between 1 and 63 characters long, and must conform to the following regular expression: p{Ll}p{Lo}{0,62} Label values must be between 1 and 63 characters long, and must conform to the following regular expression: [p{Ll}p{Lo}p{N}_-]{0,63} No more than 32 labels can be associated with a given job.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_workflow_template#labels DataprocWorkflowTemplate#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// pig_job block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_workflow_template#pig_job DataprocWorkflowTemplate#pig_job}
	PigJob *DataprocWorkflowTemplateJobsPigJob `field:"optional" json:"pigJob" yaml:"pigJob"`
	// Optional.
	//
	// The optional list of prerequisite job step_ids. If not specified, the job will start at the beginning of workflow.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_workflow_template#prerequisite_step_ids DataprocWorkflowTemplate#prerequisite_step_ids}
	PrerequisiteStepIds *[]*string `field:"optional" json:"prerequisiteStepIds" yaml:"prerequisiteStepIds"`
	// presto_job block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_workflow_template#presto_job DataprocWorkflowTemplate#presto_job}
	PrestoJob *DataprocWorkflowTemplateJobsPrestoJob `field:"optional" json:"prestoJob" yaml:"prestoJob"`
	// pyspark_job block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_workflow_template#pyspark_job DataprocWorkflowTemplate#pyspark_job}
	PysparkJob *DataprocWorkflowTemplateJobsPysparkJob `field:"optional" json:"pysparkJob" yaml:"pysparkJob"`
	// scheduling block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_workflow_template#scheduling DataprocWorkflowTemplate#scheduling}
	Scheduling *DataprocWorkflowTemplateJobsScheduling `field:"optional" json:"scheduling" yaml:"scheduling"`
	// spark_job block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_workflow_template#spark_job DataprocWorkflowTemplate#spark_job}
	SparkJob *DataprocWorkflowTemplateJobsSparkJob `field:"optional" json:"sparkJob" yaml:"sparkJob"`
	// spark_r_job block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_workflow_template#spark_r_job DataprocWorkflowTemplate#spark_r_job}
	SparkRJob *DataprocWorkflowTemplateJobsSparkRJob `field:"optional" json:"sparkRJob" yaml:"sparkRJob"`
	// spark_sql_job block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/dataproc_workflow_template#spark_sql_job DataprocWorkflowTemplate#spark_sql_job}
	SparkSqlJob *DataprocWorkflowTemplateJobsSparkSqlJob `field:"optional" json:"sparkSqlJob" yaml:"sparkSqlJob"`
}

