package sqluser


type SqlUserSqlServerUserDetails struct {
	// If the user has been disabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_user#disabled SqlUser#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// The server roles for this user in the database.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_user#server_roles SqlUser#server_roles}
	ServerRoles *[]*string `field:"optional" json:"serverRoles" yaml:"serverRoles"`
}

