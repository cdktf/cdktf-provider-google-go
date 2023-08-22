package sqlsslcert


type SqlSslCertTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/sql_ssl_cert#create SqlSslCert#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.79.0/docs/resources/sql_ssl_cert#delete SqlSslCert#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

