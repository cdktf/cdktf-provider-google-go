package cloudrunservice


type CloudRunServiceTemplateSpecContainers struct {
	// Docker image name. This is most often a reference to a container located in the container registry, such as gcr.io/cloudrun/hello.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_run_service#image CloudRunService#image}
	Image *string `field:"required" json:"image" yaml:"image"`
	// Arguments to the entrypoint. The docker image's CMD is used if this is not provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_run_service#args CloudRunService#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_run_service#command CloudRunService#command}
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// env block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_run_service#env CloudRunService#env}
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// env_from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_run_service#env_from CloudRunService#env_from}
	EnvFrom interface{} `field:"optional" json:"envFrom" yaml:"envFrom"`
	// liveness_probe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_run_service#liveness_probe CloudRunService#liveness_probe}
	LivenessProbe *CloudRunServiceTemplateSpecContainersLivenessProbe `field:"optional" json:"livenessProbe" yaml:"livenessProbe"`
	// Name of the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_run_service#name CloudRunService#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_run_service#ports CloudRunService#ports}
	Ports interface{} `field:"optional" json:"ports" yaml:"ports"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_run_service#resources CloudRunService#resources}
	Resources *CloudRunServiceTemplateSpecContainersResources `field:"optional" json:"resources" yaml:"resources"`
	// startup_probe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_run_service#startup_probe CloudRunService#startup_probe}
	StartupProbe *CloudRunServiceTemplateSpecContainersStartupProbe `field:"optional" json:"startupProbe" yaml:"startupProbe"`
	// volume_mounts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_run_service#volume_mounts CloudRunService#volume_mounts}
	VolumeMounts interface{} `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
	// Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/cloud_run_service#working_dir CloudRunService#working_dir}
	WorkingDir *string `field:"optional" json:"workingDir" yaml:"workingDir"`
}

