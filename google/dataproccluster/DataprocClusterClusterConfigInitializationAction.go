package dataproccluster


type DataprocClusterClusterConfigInitializationAction struct {
	// The script to be executed during initialization of the cluster.
	//
	// The script must be a GCS file with a gs:// prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_cluster#script DataprocCluster#script}
	Script *string `field:"required" json:"script" yaml:"script"`
	// The maximum duration (in seconds) which script is allowed to take to execute its action.
	//
	// GCP will default to a predetermined computed value if not set (currently 300).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_cluster#timeout_sec DataprocCluster#timeout_sec}
	TimeoutSec *float64 `field:"optional" json:"timeoutSec" yaml:"timeoutSec"`
}

