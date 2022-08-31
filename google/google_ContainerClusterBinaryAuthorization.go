// Prebuilt google Provider for Terraform CDK (cdktf)
package google


type ContainerClusterBinaryAuthorization struct {
	// Enable Binary Authorization for this cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Mode of operation for Binary Authorization policy evaluation.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/container_cluster#evaluation_mode ContainerCluster#evaluation_mode}
	EvaluationMode *string `field:"optional" json:"evaluationMode" yaml:"evaluationMode"`
}

