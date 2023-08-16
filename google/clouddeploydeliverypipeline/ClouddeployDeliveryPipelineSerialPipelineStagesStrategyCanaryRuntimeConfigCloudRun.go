package clouddeploydeliverypipeline


type ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigCloudRun struct {
	// Whether Cloud Deploy should update the traffic stanza in a Cloud Run Service on the user's behalf to facilitate traffic splitting.
	//
	// This is required to be true for CanaryDeployments, but optional for CustomCanaryDeployments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/clouddeploy_delivery_pipeline#automatic_traffic_control ClouddeployDeliveryPipeline#automatic_traffic_control}
	AutomaticTrafficControl interface{} `field:"optional" json:"automaticTrafficControl" yaml:"automaticTrafficControl"`
}

