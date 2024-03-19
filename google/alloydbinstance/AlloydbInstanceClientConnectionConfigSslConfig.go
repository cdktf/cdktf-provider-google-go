// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alloydbinstance


type AlloydbInstanceClientConnectionConfigSslConfig struct {
	// SSL mode. Specifies client-server SSL/TLS connection behavior. Possible values: ["ENCRYPTED_ONLY", "ALLOW_UNENCRYPTED_AND_ENCRYPTED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/alloydb_instance#ssl_mode AlloydbInstance#ssl_mode}
	SslMode *string `field:"optional" json:"sslMode" yaml:"sslMode"`
}

