package dataprocautoscalingpolicy


type DataprocAutoscalingPolicyBasicAlgorithmYarnConfig struct {
	// Timeout for YARN graceful decommissioning of Node Managers.
	//
	// Specifies the
	// duration to wait for jobs to complete before forcefully removing workers
	// (and potentially interrupting jobs). Only applicable to downscaling operations.
	//
	// Bounds: [0s, 1d].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_autoscaling_policy#graceful_decommission_timeout DataprocAutoscalingPolicy#graceful_decommission_timeout}
	GracefulDecommissionTimeout *string `field:"required" json:"gracefulDecommissionTimeout" yaml:"gracefulDecommissionTimeout"`
	// Fraction of average pending memory in the last cooldown period for which to remove workers.
	//
	// A scale-down factor of 1 will result in scaling down so that there
	// is no available memory remaining after the update (more aggressive scaling).
	// A scale-down factor of 0 disables removing workers, which can be beneficial for
	// autoscaling a single job.
	//
	// Bounds: [0.0, 1.0].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_autoscaling_policy#scale_down_factor DataprocAutoscalingPolicy#scale_down_factor}
	ScaleDownFactor *float64 `field:"required" json:"scaleDownFactor" yaml:"scaleDownFactor"`
	// Fraction of average pending memory in the last cooldown period for which to add workers.
	//
	// A scale-up factor of 1.0 will result in scaling up so that there
	// is no pending memory remaining after the update (more aggressive scaling).
	// A scale-up factor closer to 0 will result in a smaller magnitude of scaling up
	// (less aggressive scaling).
	//
	// Bounds: [0.0, 1.0].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_autoscaling_policy#scale_up_factor DataprocAutoscalingPolicy#scale_up_factor}
	ScaleUpFactor *float64 `field:"required" json:"scaleUpFactor" yaml:"scaleUpFactor"`
	// Minimum scale-down threshold as a fraction of total cluster size before scaling occurs.
	//
	// For example, in a 20-worker cluster, a threshold of 0.1 means the autoscaler must
	// recommend at least a 2 worker scale-down for the cluster to scale. A threshold of 0
	// means the autoscaler will scale down on any recommended change.
	//
	// Bounds: [0.0, 1.0]. Default: 0.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_autoscaling_policy#scale_down_min_worker_fraction DataprocAutoscalingPolicy#scale_down_min_worker_fraction}
	ScaleDownMinWorkerFraction *float64 `field:"optional" json:"scaleDownMinWorkerFraction" yaml:"scaleDownMinWorkerFraction"`
	// Minimum scale-up threshold as a fraction of total cluster size before scaling occurs.
	//
	// For example, in a 20-worker cluster, a threshold of 0.1 means the autoscaler
	// must recommend at least a 2-worker scale-up for the cluster to scale. A threshold of
	// 0 means the autoscaler will scale up on any recommended change.
	//
	// Bounds: [0.0, 1.0]. Default: 0.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_autoscaling_policy#scale_up_min_worker_fraction DataprocAutoscalingPolicy#scale_up_min_worker_fraction}
	ScaleUpMinWorkerFraction *float64 `field:"optional" json:"scaleUpMinWorkerFraction" yaml:"scaleUpMinWorkerFraction"`
}

