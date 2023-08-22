package computeregionautoscaler


type ComputeRegionAutoscalerAutoscalingPolicyScalingSchedules struct {
	// The duration of time intervals (in seconds) for which this scaling schedule will be running.
	//
	// The minimum allowed value is 300.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_autoscaler#duration_sec ComputeRegionAutoscaler#duration_sec}
	DurationSec *float64 `field:"required" json:"durationSec" yaml:"durationSec"`
	// Minimum number of VM instances that autoscaler will recommend in time intervals starting according to schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_autoscaler#min_required_replicas ComputeRegionAutoscaler#min_required_replicas}
	MinRequiredReplicas *float64 `field:"required" json:"minRequiredReplicas" yaml:"minRequiredReplicas"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_autoscaler#name ComputeRegionAutoscaler#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The start timestamps of time intervals when this scaling schedule should provide a scaling signal.
	//
	// This field uses the extended cron format (with an optional year field).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_autoscaler#schedule ComputeRegionAutoscaler#schedule}
	Schedule *string `field:"required" json:"schedule" yaml:"schedule"`
	// A description of a scaling schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_autoscaler#description ComputeRegionAutoscaler#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A boolean value that specifies if a scaling schedule can influence autoscaler recommendations.
	//
	// If set to true, then a scaling schedule has no effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_autoscaler#disabled ComputeRegionAutoscaler#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// The time zone to be used when interpreting the schedule.
	//
	// The value of this field must be a time zone name from the tz database: http://en.wikipedia.org/wiki/Tz_database.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/compute_region_autoscaler#time_zone ComputeRegionAutoscaler#time_zone}
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

