package lookerinstance


type LookerInstanceDenyMaintenancePeriod struct {
	// end_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/looker_instance#end_date LookerInstance#end_date}
	EndDate *LookerInstanceDenyMaintenancePeriodEndDate `field:"required" json:"endDate" yaml:"endDate"`
	// start_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/looker_instance#start_date LookerInstance#start_date}
	StartDate *LookerInstanceDenyMaintenancePeriodStartDate `field:"required" json:"startDate" yaml:"startDate"`
	// time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/looker_instance#time LookerInstance#time}
	Time *LookerInstanceDenyMaintenancePeriodTime `field:"required" json:"time" yaml:"time"`
}

