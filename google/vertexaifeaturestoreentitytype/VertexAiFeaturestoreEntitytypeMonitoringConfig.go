package vertexaifeaturestoreentitytype


type VertexAiFeaturestoreEntitytypeMonitoringConfig struct {
	// categorical_threshold_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/vertex_ai_featurestore_entitytype#categorical_threshold_config VertexAiFeaturestoreEntitytype#categorical_threshold_config}
	CategoricalThresholdConfig *VertexAiFeaturestoreEntitytypeMonitoringConfigCategoricalThresholdConfig `field:"optional" json:"categoricalThresholdConfig" yaml:"categoricalThresholdConfig"`
	// import_features_analysis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/vertex_ai_featurestore_entitytype#import_features_analysis VertexAiFeaturestoreEntitytype#import_features_analysis}
	ImportFeaturesAnalysis *VertexAiFeaturestoreEntitytypeMonitoringConfigImportFeaturesAnalysis `field:"optional" json:"importFeaturesAnalysis" yaml:"importFeaturesAnalysis"`
	// numerical_threshold_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/vertex_ai_featurestore_entitytype#numerical_threshold_config VertexAiFeaturestoreEntitytype#numerical_threshold_config}
	NumericalThresholdConfig *VertexAiFeaturestoreEntitytypeMonitoringConfigNumericalThresholdConfig `field:"optional" json:"numericalThresholdConfig" yaml:"numericalThresholdConfig"`
	// snapshot_analysis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/vertex_ai_featurestore_entitytype#snapshot_analysis VertexAiFeaturestoreEntitytype#snapshot_analysis}
	SnapshotAnalysis *VertexAiFeaturestoreEntitytypeMonitoringConfigSnapshotAnalysis `field:"optional" json:"snapshotAnalysis" yaml:"snapshotAnalysis"`
}

