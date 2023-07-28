package clouddeploydeliverypipeline


type ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes struct {
	// gateway_service_mesh block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.75.1/docs/resources/clouddeploy_delivery_pipeline#gateway_service_mesh ClouddeployDeliveryPipeline#gateway_service_mesh}
	GatewayServiceMesh *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh `field:"optional" json:"gatewayServiceMesh" yaml:"gatewayServiceMesh"`
	// service_networking block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.75.1/docs/resources/clouddeploy_delivery_pipeline#service_networking ClouddeployDeliveryPipeline#service_networking}
	ServiceNetworking *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking `field:"optional" json:"serviceNetworking" yaml:"serviceNetworking"`
}

