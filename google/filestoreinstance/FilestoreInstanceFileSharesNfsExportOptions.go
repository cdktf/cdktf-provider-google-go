package filestoreinstance


type FilestoreInstanceFileSharesNfsExportOptions struct {
	// Either READ_ONLY, for allowing only read requests on the exported directory, or READ_WRITE, for allowing both read and write requests.
	//
	// The default is READ_WRITE. Default value: "READ_WRITE" Possible values: ["READ_ONLY", "READ_WRITE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/filestore_instance#access_mode FilestoreInstance#access_mode}
	AccessMode *string `field:"optional" json:"accessMode" yaml:"accessMode"`
	// An integer representing the anonymous group id with a default value of 65534.
	//
	// Anon_gid may only be set with squashMode of ROOT_SQUASH. An error will be returned
	// if this field is specified for other squashMode settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/filestore_instance#anon_gid FilestoreInstance#anon_gid}
	AnonGid *float64 `field:"optional" json:"anonGid" yaml:"anonGid"`
	// An integer representing the anonymous user id with a default value of 65534.
	//
	// Anon_uid may only be set with squashMode of ROOT_SQUASH. An error will be returned
	// if this field is specified for other squashMode settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/filestore_instance#anon_uid FilestoreInstance#anon_uid}
	AnonUid *float64 `field:"optional" json:"anonUid" yaml:"anonUid"`
	// List of either IPv4 addresses, or ranges in CIDR notation which may mount the file share.
	//
	// Overlapping IP ranges are not allowed, both within and across NfsExportOptions. An error will be returned.
	// The limit is 64 IP ranges/addresses for each FileShareConfig among all NfsExportOptions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/filestore_instance#ip_ranges FilestoreInstance#ip_ranges}
	IpRanges *[]*string `field:"optional" json:"ipRanges" yaml:"ipRanges"`
	// Either NO_ROOT_SQUASH, for allowing root access on the exported directory, or ROOT_SQUASH, for not allowing root access.
	//
	// The default is NO_ROOT_SQUASH. Default value: "NO_ROOT_SQUASH" Possible values: ["NO_ROOT_SQUASH", "ROOT_SQUASH"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.78.0/docs/resources/filestore_instance#squash_mode FilestoreInstance#squash_mode}
	SquashMode *string `field:"optional" json:"squashMode" yaml:"squashMode"`
}

