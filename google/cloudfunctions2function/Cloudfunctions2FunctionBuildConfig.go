package cloudfunctions2function


type Cloudfunctions2FunctionBuildConfig struct {
	// User managed repository created in Artifact Registry optionally with a customer managed encryption key.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions2_function#docker_repository Cloudfunctions2Function#docker_repository}
	DockerRepository *string `field:"optional" json:"dockerRepository" yaml:"dockerRepository"`
	// The name of the function (as defined in source code) that will be executed.
	//
	// Defaults to the resource name suffix, if not specified. For backward
	// compatibility, if function with given name is not found, then the system
	// will try to use function named "function". For Node.js this is name of a
	// function exported by the module specified in source_location.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions2_function#entry_point Cloudfunctions2Function#entry_point}
	EntryPoint *string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// User-provided build-time environment variables for the function.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions2_function#environment_variables Cloudfunctions2Function#environment_variables}
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// The runtime in which to run the function. Required when deploying a new function, optional when updating an existing function.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions2_function#runtime Cloudfunctions2Function#runtime}
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions2_function#source Cloudfunctions2Function#source}
	Source *Cloudfunctions2FunctionBuildConfigSource `field:"optional" json:"source" yaml:"source"`
	// Name of the Cloud Build Custom Worker Pool that should be used to build the function.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/cloudfunctions2_function#worker_pool Cloudfunctions2Function#worker_pool}
	WorkerPool *string `field:"optional" json:"workerPool" yaml:"workerPool"`
}

