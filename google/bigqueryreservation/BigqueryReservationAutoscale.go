package bigqueryreservation


type BigqueryReservationAutoscale struct {
	// Number of slots to be scaled when needed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_reservation#max_slots BigqueryReservation#max_slots}
	MaxSlots *float64 `field:"optional" json:"maxSlots" yaml:"maxSlots"`
}

