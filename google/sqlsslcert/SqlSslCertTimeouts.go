package sqlsslcert


type SqlSslCertTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_ssl_cert#create SqlSslCert#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/sql_ssl_cert#delete SqlSslCert#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

