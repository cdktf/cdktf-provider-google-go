// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package biglakedatabase


type BiglakeDatabaseHiveOptions struct {
	// Cloud Storage folder URI where the database data is stored, starting with "gs://".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/biglake_database#location_uri BiglakeDatabase#location_uri}
	LocationUri *string `field:"optional" json:"locationUri" yaml:"locationUri"`
	// Stores user supplied Hive database parameters.
	//
	// An object containing a
	// list of"key": value pairs.
	// Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.49.3/docs/resources/biglake_database#parameters BiglakeDatabase#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

