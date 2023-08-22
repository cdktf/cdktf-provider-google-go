package computeregionautoscaler


type ComputeRegionAutoscalerAutoscalingPolicy struct {
	// The maximum number of instances that the autoscaler can scale up to.
	//
	// This is required when creating or updating an autoscaler. The
	// maximum number of replicas should not be lower than minimal number
	// of replicas.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_autoscaler#max_replicas ComputeRegionAutoscaler#max_replicas}
	MaxReplicas *float64 `field:"required" json:"maxReplicas" yaml:"maxReplicas"`
	// The minimum number of replicas that the autoscaler can scale down to.
	//
	// This cannot be less than 0. If not provided, autoscaler will
	// choose a default value depending on maximum number of instances
	// allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_autoscaler#min_replicas ComputeRegionAutoscaler#min_replicas}
	MinReplicas *float64 `field:"required" json:"minReplicas" yaml:"minReplicas"`
	// The number of seconds that the autoscaler should wait before it starts collecting information from a new instance.
	//
	// This prevents
	// the autoscaler from collecting information when the instance is
	// initializing, during which the collected usage would not be
	// reliable. The default time autoscaler waits is 60 seconds.
	//
	// Virtual machine initialization times might vary because of
	// numerous factors. We recommend that you test how long an
	// instance may take to initialize. To do this, create an instance
	// and time the startup process.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_autoscaler#cooldown_period ComputeRegionAutoscaler#cooldown_period}
	CooldownPeriod *float64 `field:"optional" json:"cooldownPeriod" yaml:"cooldownPeriod"`
	// cpu_utilization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_autoscaler#cpu_utilization ComputeRegionAutoscaler#cpu_utilization}
	CpuUtilization *ComputeRegionAutoscalerAutoscalingPolicyCpuUtilization `field:"optional" json:"cpuUtilization" yaml:"cpuUtilization"`
	// load_balancing_utilization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_autoscaler#load_balancing_utilization ComputeRegionAutoscaler#load_balancing_utilization}
	LoadBalancingUtilization *ComputeRegionAutoscalerAutoscalingPolicyLoadBalancingUtilization `field:"optional" json:"loadBalancingUtilization" yaml:"loadBalancingUtilization"`
	// metric block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_autoscaler#metric ComputeRegionAutoscaler#metric}
	Metric interface{} `field:"optional" json:"metric" yaml:"metric"`
	// Defines operating mode for this policy. Default value: "ON" Possible values: ["OFF", "ONLY_UP", "ON"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_autoscaler#mode ComputeRegionAutoscaler#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// scale_in_control block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_autoscaler#scale_in_control ComputeRegionAutoscaler#scale_in_control}
	ScaleInControl *ComputeRegionAutoscalerAutoscalingPolicyScaleInControl `field:"optional" json:"scaleInControl" yaml:"scaleInControl"`
	// scaling_schedules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_autoscaler#scaling_schedules ComputeRegionAutoscaler#scaling_schedules}
	ScalingSchedules interface{} `field:"optional" json:"scalingSchedules" yaml:"scalingSchedules"`
}

