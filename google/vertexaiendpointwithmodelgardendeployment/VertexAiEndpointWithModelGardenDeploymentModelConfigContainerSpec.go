// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vertexaiendpointwithmodelgardendeployment


type VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpec struct {
	// URI of the Docker image to be used as the custom container for serving predictions.
	//
	// This URI must identify an image in Artifact Registry or
	// Container Registry. Learn more about the [container publishing
	// requirements](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#publishing),
	// including permissions requirements for the Vertex AI Service Agent.
	//
	// The container image is ingested upon ModelService.UploadModel, stored
	// internally, and this original path is afterwards not used.
	//
	// To learn about the requirements for the Docker image itself, see
	// [Custom container
	// requirements](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#).
	//
	// You can use the URI to one of Vertex AI's [pre-built container images for
	// prediction](https://cloud.google.com/vertex-ai/docs/predictions/pre-built-containers)
	// in this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#image_uri VertexAiEndpointWithModelGardenDeployment#image_uri}
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// Specifies arguments for the command that runs when the container starts.
	//
	// This overrides the container's
	// ['CMD'](https://docs.docker.com/engine/reference/builder/#cmd). Specify
	// this field as an array of executable and arguments, similar to a Docker
	// 'CMD''s "default parameters" form.
	//
	// If you don't specify this field but do specify the
	// command field, then the command from the
	// 'command' field runs without any additional arguments. See the
	// [Kubernetes documentation about how the
	// 'command' and 'args' fields interact with a container's 'ENTRYPOINT' and
	// 'CMD'](https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#notes).
	//
	// If you don't specify this field and don't specify the 'command' field,
	// then the container's
	// ['ENTRYPOINT'](https://docs.docker.com/engine/reference/builder/#cmd) and
	// 'CMD' determine what runs based on their default behavior. See the Docker
	// documentation about [how 'CMD' and 'ENTRYPOINT'
	// interact](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact).
	//
	// In this field, you can reference [environment variables
	// set by Vertex
	// AI](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables)
	// and environment variables set in the env field.
	// You cannot reference environment variables set in the Docker image. In
	// order for environment variables to be expanded, reference them by using the
	// following syntax:$(VARIABLE_NAME)
	// Note that this differs from Bash variable expansion, which does not use
	// parentheses. If a variable cannot be resolved, the reference in the input
	// string is used unchanged. To avoid variable expansion, you can escape this
	// syntax with '$$'; for example:$$(VARIABLE_NAME)
	// This field corresponds to the 'args' field of the Kubernetes Containers
	// [v1 core
	// API](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#args VertexAiEndpointWithModelGardenDeployment#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Specifies the command that runs when the container starts.
	//
	// This overrides
	// the container's
	// [ENTRYPOINT](https://docs.docker.com/engine/reference/builder/#entrypoint).
	// Specify this field as an array of executable and arguments, similar to a
	// Docker 'ENTRYPOINT''s "exec" form, not its "shell" form.
	//
	// If you do not specify this field, then the container's 'ENTRYPOINT' runs,
	// in conjunction with the args field or the
	// container's ['CMD'](https://docs.docker.com/engine/reference/builder/#cmd),
	// if either exists. If this field is not specified and the container does not
	// have an 'ENTRYPOINT', then refer to the Docker documentation about [how
	// 'CMD' and 'ENTRYPOINT'
	// interact](https://docs.docker.com/engine/reference/builder/#understand-how-cmd-and-entrypoint-interact).
	//
	// If you specify this field, then you can also specify the 'args' field to
	// provide additional arguments for this command. However, if you specify this
	// field, then the container's 'CMD' is ignored. See the
	// [Kubernetes documentation about how the
	// 'command' and 'args' fields interact with a container's 'ENTRYPOINT' and
	// 'CMD'](https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#notes).
	//
	// In this field, you can reference [environment variables set by Vertex
	// AI](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables)
	// and environment variables set in the env field.
	// You cannot reference environment variables set in the Docker image. In
	// order for environment variables to be expanded, reference them by using the
	// following syntax:$(VARIABLE_NAME)
	// Note that this differs from Bash variable expansion, which does not use
	// parentheses. If a variable cannot be resolved, the reference in the input
	// string is used unchanged. To avoid variable expansion, you can escape this
	// syntax with '$$'; for example:$$(VARIABLE_NAME)
	// This field corresponds to the 'command' field of the Kubernetes Containers
	// [v1 core
	// API](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#container-v1-core).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#command VertexAiEndpointWithModelGardenDeployment#command}
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Deployment timeout. Limit for deployment timeout is 2 hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#deployment_timeout VertexAiEndpointWithModelGardenDeployment#deployment_timeout}
	DeploymentTimeout *string `field:"optional" json:"deploymentTimeout" yaml:"deploymentTimeout"`
	// env block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#env VertexAiEndpointWithModelGardenDeployment#env}
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// grpc_ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#grpc_ports VertexAiEndpointWithModelGardenDeployment#grpc_ports}
	GrpcPorts interface{} `field:"optional" json:"grpcPorts" yaml:"grpcPorts"`
	// health_probe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#health_probe VertexAiEndpointWithModelGardenDeployment#health_probe}
	HealthProbe *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecHealthProbe `field:"optional" json:"healthProbe" yaml:"healthProbe"`
	// HTTP path on the container to send health checks to.
	//
	// Vertex AI
	// intermittently sends GET requests to this path on the container's IP
	// address and port to check that the container is healthy. Read more about
	// [health
	// checks](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#health).
	//
	// For example, if you set this field to '/bar', then Vertex AI
	// intermittently sends a GET request to the '/bar' path on the port of your
	// container specified by the first value of this 'ModelContainerSpec''s
	// ports field.
	//
	// If you don't specify this field, it defaults to the following value when
	// you deploy this Model to an Endpoint:/v1/endpoints/ENDPOINT/deployedModels/DEPLOYED_MODEL:predict
	// The placeholders in this value are replaced as follows:
	//
	// * ENDPOINT: The last segment (following 'endpoints/')of the
	// Endpoint.name][] field of the Endpoint where this Model has been
	// deployed. (Vertex AI makes this value available to your container code
	// as the ['AIP_ENDPOINT_ID' environment
	// variable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).)
	//
	// * DEPLOYED_MODEL: DeployedModel.id of the 'DeployedModel'.
	// (Vertex AI makes this value available to your container code as the
	// ['AIP_DEPLOYED_MODEL_ID' environment
	// variable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#health_route VertexAiEndpointWithModelGardenDeployment#health_route}
	HealthRoute *string `field:"optional" json:"healthRoute" yaml:"healthRoute"`
	// liveness_probe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#liveness_probe VertexAiEndpointWithModelGardenDeployment#liveness_probe}
	LivenessProbe *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecLivenessProbe `field:"optional" json:"livenessProbe" yaml:"livenessProbe"`
	// ports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#ports VertexAiEndpointWithModelGardenDeployment#ports}
	Ports interface{} `field:"optional" json:"ports" yaml:"ports"`
	// HTTP path on the container to send prediction requests to.
	//
	// Vertex AI
	// forwards requests sent using
	// projects.locations.endpoints.predict to this
	// path on the container's IP address and port. Vertex AI then returns the
	// container's response in the API response.
	//
	// For example, if you set this field to '/foo', then when Vertex AI
	// receives a prediction request, it forwards the request body in a POST
	// request to the '/foo' path on the port of your container specified by the
	// first value of this 'ModelContainerSpec''s
	// ports field.
	//
	// If you don't specify this field, it defaults to the following value when
	// you deploy this Model to an Endpoint:/v1/endpoints/ENDPOINT/deployedModels/DEPLOYED_MODEL:predict
	// The placeholders in this value are replaced as follows:
	//
	// * ENDPOINT: The last segment (following 'endpoints/')of the
	// Endpoint.name][] field of the Endpoint where this Model has been
	// deployed. (Vertex AI makes this value available to your container code
	// as the ['AIP_ENDPOINT_ID' environment
	// variable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).)
	//
	// * DEPLOYED_MODEL: DeployedModel.id of the 'DeployedModel'.
	// (Vertex AI makes this value available to your container code
	// as the ['AIP_DEPLOYED_MODEL_ID' environment
	// variable](https://cloud.google.com/vertex-ai/docs/predictions/custom-container-requirements#aip-variables).)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#predict_route VertexAiEndpointWithModelGardenDeployment#predict_route}
	PredictRoute *string `field:"optional" json:"predictRoute" yaml:"predictRoute"`
	// The amount of the VM memory to reserve as the shared memory for the model in megabytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#shared_memory_size_mb VertexAiEndpointWithModelGardenDeployment#shared_memory_size_mb}
	SharedMemorySizeMb *string `field:"optional" json:"sharedMemorySizeMb" yaml:"sharedMemorySizeMb"`
	// startup_probe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/vertex_ai_endpoint_with_model_garden_deployment#startup_probe VertexAiEndpointWithModelGardenDeployment#startup_probe}
	StartupProbe *VertexAiEndpointWithModelGardenDeploymentModelConfigContainerSpecStartupProbe `field:"optional" json:"startupProbe" yaml:"startupProbe"`
}

