package cloudfunctionsfunction

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudfunctionsFunctionConfig struct {
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
	// A user-defined name of the function. Function names must be unique globally.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#name CloudfunctionsFunction#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The runtime in which the function is going to run. Eg. "nodejs12", "nodejs14", "python37", "go111".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#runtime CloudfunctionsFunction#runtime}
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// Memory (in MB), available to the function. Default value is 256. Possible values include 128, 256, 512, 1024, etc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#available_memory_mb CloudfunctionsFunction#available_memory_mb}
	AvailableMemoryMb *float64 `field:"optional" json:"availableMemoryMb" yaml:"availableMemoryMb"`
	// A set of key/value environment variable pairs available during build time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#build_environment_variables CloudfunctionsFunction#build_environment_variables}
	BuildEnvironmentVariables *map[string]*string `field:"optional" json:"buildEnvironmentVariables" yaml:"buildEnvironmentVariables"`
	// Name of the Cloud Build Custom Worker Pool that should be used to build the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#build_worker_pool CloudfunctionsFunction#build_worker_pool}
	BuildWorkerPool *string `field:"optional" json:"buildWorkerPool" yaml:"buildWorkerPool"`
	// Description of the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#description CloudfunctionsFunction#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docker Registry to use for storing the function's Docker images. Allowed values are CONTAINER_REGISTRY (default) and ARTIFACT_REGISTRY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#docker_registry CloudfunctionsFunction#docker_registry}
	DockerRegistry *string `field:"optional" json:"dockerRegistry" yaml:"dockerRegistry"`
	// User managed repository created in Artifact Registry optionally with a customer managed encryption key.
	//
	// If specified, deployments will use Artifact Registry for storing images built with Cloud Build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#docker_repository CloudfunctionsFunction#docker_repository}
	DockerRepository *string `field:"optional" json:"dockerRepository" yaml:"dockerRepository"`
	// Name of the function that will be executed when the Google Cloud Function is triggered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#entry_point CloudfunctionsFunction#entry_point}
	EntryPoint *string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// A set of key/value environment variable pairs to assign to the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#environment_variables CloudfunctionsFunction#environment_variables}
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// event_trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#event_trigger CloudfunctionsFunction#event_trigger}
	EventTrigger *CloudfunctionsFunctionEventTrigger `field:"optional" json:"eventTrigger" yaml:"eventTrigger"`
	// The security level for the function. Defaults to SECURE_OPTIONAL. Valid only if trigger_http is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#https_trigger_security_level CloudfunctionsFunction#https_trigger_security_level}
	HttpsTriggerSecurityLevel *string `field:"optional" json:"httpsTriggerSecurityLevel" yaml:"httpsTriggerSecurityLevel"`
	// URL which triggers function execution. Returned only if trigger_http is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#https_trigger_url CloudfunctionsFunction#https_trigger_url}
	HttpsTriggerUrl *string `field:"optional" json:"httpsTriggerUrl" yaml:"httpsTriggerUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#id CloudfunctionsFunction#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// String value that controls what traffic can reach the function.
	//
	// Allowed values are ALLOW_ALL and ALLOW_INTERNAL_ONLY. Changes to this field will recreate the cloud function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#ingress_settings CloudfunctionsFunction#ingress_settings}
	IngressSettings *string `field:"optional" json:"ingressSettings" yaml:"ingressSettings"`
	// Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#kms_key_name CloudfunctionsFunction#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// A set of key/value label pairs to assign to the function. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#labels CloudfunctionsFunction#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The limit on the maximum number of function instances that may coexist at a given time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#max_instances CloudfunctionsFunction#max_instances}
	MaxInstances *float64 `field:"optional" json:"maxInstances" yaml:"maxInstances"`
	// The limit on the minimum number of function instances that may coexist at a given time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#min_instances CloudfunctionsFunction#min_instances}
	MinInstances *float64 `field:"optional" json:"minInstances" yaml:"minInstances"`
	// Project of the function. If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#project CloudfunctionsFunction#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Region of function. If it is not provided, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#region CloudfunctionsFunction#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// secret_environment_variables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#secret_environment_variables CloudfunctionsFunction#secret_environment_variables}
	SecretEnvironmentVariables interface{} `field:"optional" json:"secretEnvironmentVariables" yaml:"secretEnvironmentVariables"`
	// secret_volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#secret_volumes CloudfunctionsFunction#secret_volumes}
	SecretVolumes interface{} `field:"optional" json:"secretVolumes" yaml:"secretVolumes"`
	// If provided, the self-provided service account to run the function with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#service_account_email CloudfunctionsFunction#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// The GCS bucket containing the zip archive which contains the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#source_archive_bucket CloudfunctionsFunction#source_archive_bucket}
	SourceArchiveBucket *string `field:"optional" json:"sourceArchiveBucket" yaml:"sourceArchiveBucket"`
	// The source archive object (file) in archive bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#source_archive_object CloudfunctionsFunction#source_archive_object}
	SourceArchiveObject *string `field:"optional" json:"sourceArchiveObject" yaml:"sourceArchiveObject"`
	// source_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#source_repository CloudfunctionsFunction#source_repository}
	SourceRepository *CloudfunctionsFunctionSourceRepository `field:"optional" json:"sourceRepository" yaml:"sourceRepository"`
	// Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#timeout CloudfunctionsFunction#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#timeouts CloudfunctionsFunction#timeouts}
	Timeouts *CloudfunctionsFunctionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Boolean variable.
	//
	// Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as https_trigger_url. Cannot be used with trigger_bucket and trigger_topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#trigger_http CloudfunctionsFunction#trigger_http}
	TriggerHttp interface{} `field:"optional" json:"triggerHttp" yaml:"triggerHttp"`
	// The VPC Network Connector that this cloud function can connect to.
	//
	// It can be either the fully-qualified URI, or the short name of the network connector resource. The format of this field is projects/*\/locations/*\/connectors/*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#vpc_connector CloudfunctionsFunction#vpc_connector}
	VpcConnector *string `field:"optional" json:"vpcConnector" yaml:"vpcConnector"`
	// The egress settings for the connector, controlling what traffic is diverted through it.
	//
	// Allowed values are ALL_TRAFFIC and PRIVATE_RANGES_ONLY. Defaults to PRIVATE_RANGES_ONLY. If unset, this field preserves the previously set value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudfunctions_function#vpc_connector_egress_settings CloudfunctionsFunction#vpc_connector_egress_settings}
	VpcConnectorEgressSettings *string `field:"optional" json:"vpcConnectorEgressSettings" yaml:"vpcConnectorEgressSettings"`
}

