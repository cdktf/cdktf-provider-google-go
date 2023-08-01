package datastoreindex


type DatastoreIndexTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/datastore_index#create DatastoreIndex#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.76.0/docs/resources/datastore_index#delete DatastoreIndex#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

