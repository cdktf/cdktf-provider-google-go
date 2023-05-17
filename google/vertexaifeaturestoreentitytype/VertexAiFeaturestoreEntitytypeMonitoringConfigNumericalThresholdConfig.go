package vertexaifeaturestoreentitytype


type VertexAiFeaturestoreEntitytypeMonitoringConfigNumericalThresholdConfig struct {
	// Specify a threshold value that can trigger the alert.
	//
	// For numerical feature, the distribution distance is calculated by Jensen–Shannon divergence. Each feature must have a non-zero threshold if they need to be monitored. Otherwise no alert will be triggered for that feature. The default value is 0.3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/vertex_ai_featurestore_entitytype#value VertexAiFeaturestoreEntitytype#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

