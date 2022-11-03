package vertexaifeaturestoreentitytype


type VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysis struct {
	// The monitoring schedule for snapshot analysis.
	//
	// For EntityType-level config: unset / disabled = true indicates disabled by default for Features under it; otherwise by default enable snapshot analysis monitoring with monitoringInterval for Features under it.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/vertex_ai_featurestore_entitytype#disabled VertexAiFeaturestoreEntitytype#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}

