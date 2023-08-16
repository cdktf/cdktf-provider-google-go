package dataplextask


type DataplexTaskExecutionSpec struct {
	// Service account to use to execute a task.
	//
	// If not provided, the default Compute service account for the project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_task#service_account DataplexTask#service_account}
	ServiceAccount *string `field:"required" json:"serviceAccount" yaml:"serviceAccount"`
	// The arguments to pass to the task.
	//
	// The args can use placeholders of the format ${placeholder} as part of key/value string. These will be interpolated before passing the args to the driver. Currently supported placeholders: - ${taskId} - ${job_time} To pass positional args, set the key as TASK_ARGS. The value should be a comma-separated string of all the positional arguments. To use a delimiter other than comma, refer to https://cloud.google.com/sdk/gcloud/reference/topic/escaping. In case of other keys being present in the args, then TASK_ARGS will be passed as the last argument. An object containing a list of 'key': value pairs. Example: { 'name': 'wrench', 'mass': '1.3kg', 'count': '3' }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_task#args DataplexTask#args}
	Args *map[string]*string `field:"optional" json:"args" yaml:"args"`
	// The Cloud KMS key to use for encryption, of the form: projects/{project_number}/locations/{locationId}/keyRings/{key-ring-name}/cryptoKeys/{key-name}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_task#kms_key DataplexTask#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The maximum duration after which the job execution is expired.
	//
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: '3.5s'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_task#max_job_execution_lifetime DataplexTask#max_job_execution_lifetime}
	MaxJobExecutionLifetime *string `field:"optional" json:"maxJobExecutionLifetime" yaml:"maxJobExecutionLifetime"`
	// The project in which jobs are run.
	//
	// By default, the project containing the Lake is used. If a project is provided, the ExecutionSpec.service_account must belong to this project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_task#project DataplexTask#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

