package dataproccluster


type DataprocClusterClusterConfigGceClusterConfigReservationAffinity struct {
	// Type of reservation to consume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_cluster#consume_reservation_type DataprocCluster#consume_reservation_type}
	ConsumeReservationType *string `field:"optional" json:"consumeReservationType" yaml:"consumeReservationType"`
	// Corresponds to the label key of reservation resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_cluster#key DataprocCluster#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Corresponds to the label values of reservation resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/dataproc_cluster#values DataprocCluster#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

