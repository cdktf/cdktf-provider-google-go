package lookerinstance


type LookerInstanceDenyMaintenancePeriodEndDate struct {
	// Day of a month.
	//
	// Must be from 1 to 31 and valid for the year and month, or 0
	// to specify a year by itself or a year and month where the day isn't significant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/looker_instance#day LookerInstance#day}
	Day *float64 `field:"optional" json:"day" yaml:"day"`
	// Month of a year.
	//
	// Must be from 1 to 12, or 0 to specify a year without a
	// month and day.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/looker_instance#month LookerInstance#month}
	Month *float64 `field:"optional" json:"month" yaml:"month"`
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without a year.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/looker_instance#year LookerInstance#year}
	Year *float64 `field:"optional" json:"year" yaml:"year"`
}

