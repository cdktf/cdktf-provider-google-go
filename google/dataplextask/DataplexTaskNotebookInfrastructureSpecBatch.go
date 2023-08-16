package dataplextask


type DataplexTaskNotebookInfrastructureSpecBatch struct {
	// Total number of job executors. Executor Count should be between 2 and 100. [Default=2].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_task#executors_count DataplexTask#executors_count}
	ExecutorsCount *float64 `field:"optional" json:"executorsCount" yaml:"executorsCount"`
	// Max configurable executors.
	//
	// If maxExecutorsCount > executorsCount, then auto-scaling is enabled. Max Executor Count should be between 2 and 1000. [Default=1000]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_task#max_executors_count DataplexTask#max_executors_count}
	MaxExecutorsCount *float64 `field:"optional" json:"maxExecutorsCount" yaml:"maxExecutorsCount"`
}

