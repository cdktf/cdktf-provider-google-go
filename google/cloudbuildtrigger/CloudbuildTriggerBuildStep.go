package cloudbuildtrigger


type CloudbuildTriggerBuildStep struct {
	// The name of the container image that will run this particular build step.
	//
	// If the image is available in the host's Docker daemon's cache, it will be
	// run directly. If not, the host will attempt to pull the image first, using
	// the builder service account's credentials if necessary.
	//
	// The Docker daemon's cache will already have the latest versions of all of
	// the officially supported build steps (see https://github.com/GoogleCloudPlatform/cloud-builders
	// for images and examples).
	// The Docker daemon will also have cached many of the layers for some popular
	// images, like "ubuntu", "debian", but they will be refreshed at the time
	// you attempt to use them.
	//
	// If you built an image in a previous build step, it will be stored in the
	// host's Docker daemon's cache and is available to use as the name for a
	// later build step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_trigger#name CloudbuildTrigger#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Allow this build step to fail without failing the entire build if and only if the exit code is one of the specified codes.
	//
	// If 'allowFailure' is also specified, this field will take precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_trigger#allow_exit_codes CloudbuildTrigger#allow_exit_codes}
	AllowExitCodes *[]*float64 `field:"optional" json:"allowExitCodes" yaml:"allowExitCodes"`
	// Allow this build step to fail without failing the entire build.
	//
	// If false, the entire build will fail if this step fails. Otherwise, the
	// build will succeed, but this step will still have a failure status.
	// Error information will be reported in the 'failureDetail' field.
	//
	// 'allowExitCodes' takes precedence over this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_trigger#allow_failure CloudbuildTrigger#allow_failure}
	AllowFailure interface{} `field:"optional" json:"allowFailure" yaml:"allowFailure"`
	// A list of arguments that will be presented to the step when it is started.
	//
	// If the image used to run the step's container has an entrypoint, the args
	// are used as arguments to that entrypoint. If the image does not define an
	// entrypoint, the first element in args is used as the entrypoint, and the
	// remainder will be used as arguments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_trigger#args CloudbuildTrigger#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Working directory to use when running this step's container.
	//
	// If this value is a relative path, it is relative to the build's working
	// directory. If this value is absolute, it may be outside the build's working
	// directory, in which case the contents of the path may not be persisted
	// across build step executions, unless a 'volume' for that path is specified.
	//
	// If the build specifies a 'RepoSource' with 'dir' and a step with a
	// 'dir',
	// which specifies an absolute path, the 'RepoSource' 'dir' is ignored
	// for the step's execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_trigger#dir CloudbuildTrigger#dir}
	Dir *string `field:"optional" json:"dir" yaml:"dir"`
	// Entrypoint to be used instead of the build step image's default entrypoint. If unset, the image's default entrypoint is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_trigger#entrypoint CloudbuildTrigger#entrypoint}
	Entrypoint *string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// A list of environment variable definitions to be used when running a step.
	//
	// The elements are of the form "KEY=VALUE" for the environment variable
	// "KEY" being given the value "VALUE".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_trigger#env CloudbuildTrigger#env}
	Env *[]*string `field:"optional" json:"env" yaml:"env"`
	// Unique identifier for this build step, used in 'wait_for' to reference this build step as a dependency.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_trigger#id CloudbuildTrigger#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// A shell script to be executed in the step.
	//
	// When script is provided, the user cannot specify the entrypoint or args.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_trigger#script CloudbuildTrigger#script}
	Script *string `field:"optional" json:"script" yaml:"script"`
	// A list of environment variables which are encrypted using a Cloud Key Management Service crypto key.
	//
	// These values must be specified in
	// the build's 'Secret'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_trigger#secret_env CloudbuildTrigger#secret_env}
	SecretEnv *[]*string `field:"optional" json:"secretEnv" yaml:"secretEnv"`
	// Time limit for executing this build step.
	//
	// If not defined,
	// the step has no
	// time limit and will be allowed to continue to run until either it
	// completes or the build itself times out.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_trigger#timeout CloudbuildTrigger#timeout}
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
	// Output only. Stores timing information for executing this build step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_trigger#timing CloudbuildTrigger#timing}
	Timing *string `field:"optional" json:"timing" yaml:"timing"`
	// volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_trigger#volumes CloudbuildTrigger#volumes}
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
	// The ID(s) of the step(s) that this build step depends on.
	//
	// This build step will not start until all the build steps in 'wait_for'
	// have completed successfully. If 'wait_for' is empty, this build step
	// will start when all previous build steps in the 'Build.Steps' list
	// have completed successfully.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloudbuild_trigger#wait_for CloudbuildTrigger#wait_for}
	WaitFor *[]*string `field:"optional" json:"waitFor" yaml:"waitFor"`
}

