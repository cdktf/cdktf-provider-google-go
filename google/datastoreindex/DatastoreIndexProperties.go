package datastoreindex


type DatastoreIndexProperties struct {
	// The direction the index should optimize for sorting. Possible values: ["ASCENDING", "DESCENDING"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/datastore_index#direction DatastoreIndex#direction}
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// The property name to index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.65.2/docs/resources/datastore_index#name DatastoreIndex#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

