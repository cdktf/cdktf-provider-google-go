package gameservicesgameserverconfig


type GameServicesGameServerConfigScalingConfigsSchedules struct {
	// The duration for the cron job event. The duration of the event is effective after the cron job's start time.
	//
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/game_services_game_server_config#cron_job_duration GameServicesGameServerConfig#cron_job_duration}
	CronJobDuration *string `field:"optional" json:"cronJobDuration" yaml:"cronJobDuration"`
	// The cron definition of the scheduled event. See https://en.wikipedia.org/wiki/Cron. Cron spec specifies the local time as defined by the realm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/game_services_game_server_config#cron_spec GameServicesGameServerConfig#cron_spec}
	CronSpec *string `field:"optional" json:"cronSpec" yaml:"cronSpec"`
	// The end time of the event.
	//
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/game_services_game_server_config#end_time GameServicesGameServerConfig#end_time}
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The start time of the event.
	//
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/game_services_game_server_config#start_time GameServicesGameServerConfig#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

