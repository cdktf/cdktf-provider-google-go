package dataplextask


type DataplexTaskTriggerSpec struct {
	// Trigger type of the user-specified Task Possible values: ["ON_DEMAND", "RECURRING"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_task#type DataplexTask#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Prevent the task from executing.
	//
	// This does not cancel already running tasks. It is intended to temporarily disable RECURRING tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_task#disabled DataplexTask#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Number of retry attempts before aborting. Set to zero to never attempt to retry a failed task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_task#max_retries DataplexTask#max_retries}
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// Cron schedule (https://en.wikipedia.org/wiki/Cron) for running tasks periodically. To explicitly set a timezone to the cron tab, apply a prefix in the cron tab: 'CRON_TZ=${IANA_TIME_ZONE}' or 'TZ=${IANA_TIME_ZONE}'. The ${IANA_TIME_ZONE} may only be a valid string from IANA time zone database. For example, CRON_TZ=America/New_York 1 * * * *, or TZ=America/New_York 1 * * * *. This field is required for RECURRING tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_task#schedule DataplexTask#schedule}
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// The first run of the task will be after this time.
	//
	// If not specified, the task will run shortly after being submitted if ON_DEMAND and based on the schedule if RECURRING.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/dataplex_task#start_time DataplexTask#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

