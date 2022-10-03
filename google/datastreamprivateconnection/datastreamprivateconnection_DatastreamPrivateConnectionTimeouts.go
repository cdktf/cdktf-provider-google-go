package datastreamprivateconnection


type DatastreamPrivateConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastream_private_connection#create DatastreamPrivateConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastream_private_connection#delete DatastreamPrivateConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

