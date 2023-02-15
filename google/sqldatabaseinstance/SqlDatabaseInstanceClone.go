package sqldatabaseinstance


type SqlDatabaseInstanceClone struct {
	// The name of the instance from which the point in time should be restored.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#source_instance_name SqlDatabaseInstance#source_instance_name}
	SourceInstanceName *string `field:"required" json:"sourceInstanceName" yaml:"sourceInstanceName"`
	// The name of the allocated ip range for the private ip CloudSQL instance.
	//
	// For example: "google-managed-services-default". If set, the cloned instance ip will be created in the allocated range. The range name must comply with [RFC 1035](https://tools.ietf.org/html/rfc1035). Specifically, the name must be 1-63 characters long and match the regular expression [a-z]([-a-z0-9]*[a-z0-9])?.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#allocated_ip_range SqlDatabaseInstance#allocated_ip_range}
	AllocatedIpRange *string `field:"optional" json:"allocatedIpRange" yaml:"allocatedIpRange"`
	// The timestamp of the point in time that should be restored.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_database_instance#point_in_time SqlDatabaseInstance#point_in_time}
	PointInTime *string `field:"optional" json:"pointInTime" yaml:"pointInTime"`
}

