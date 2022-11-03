package datastreamconnectionprofile


type DatastreamConnectionProfilePrivateConnectivity struct {
	// A reference to a private connection resource. Format: 'projects/{project}/locations/{location}/privateConnections/{name}'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/datastream_connection_profile#private_connection DatastreamConnectionProfile#private_connection}
	PrivateConnection *string `field:"required" json:"privateConnection" yaml:"privateConnection"`
}

