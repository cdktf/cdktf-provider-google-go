package lookerinstance


type LookerInstanceMaintenanceWindowStartTime struct {
	// Hours of day in 24 hour format. Should be from 0 to 23.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/looker_instance#hours LookerInstance#hours}
	Hours *float64 `field:"optional" json:"hours" yaml:"hours"`
	// Minutes of hour of day. Must be from 0 to 59.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/looker_instance#minutes LookerInstance#minutes}
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
	// Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/looker_instance#nanos LookerInstance#nanos}
	Nanos *float64 `field:"optional" json:"nanos" yaml:"nanos"`
	// Seconds of minutes of the time. Must normally be from 0 to 59.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/looker_instance#seconds LookerInstance#seconds}
	Seconds *float64 `field:"optional" json:"seconds" yaml:"seconds"`
}

