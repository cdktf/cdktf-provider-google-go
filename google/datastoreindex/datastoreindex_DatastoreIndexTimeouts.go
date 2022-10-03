package datastoreindex


type DatastoreIndexTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastore_index#create DatastoreIndex#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastore_index#delete DatastoreIndex#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

