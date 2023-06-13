package vertexaifeaturestoreentitytype


type VertexAiFeaturestoreEntitytypeMonitoringConfigCategoricalThresholdConfig struct {
	// Specify a threshold value that can trigger the alert.
	//
	// For categorical feature, the distribution distance is calculated by L-inifinity norm. Each feature must have a non-zero threshold if they need to be monitored. Otherwise no alert will be triggered for that feature. The default value is 0.3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/vertex_ai_featurestore_entitytype#value VertexAiFeaturestoreEntitytype#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

