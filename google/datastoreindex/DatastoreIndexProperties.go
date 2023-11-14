// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datastoreindex


type DatastoreIndexProperties struct {
	// The direction the index should optimize for sorting. Possible values: ["ASCENDING", "DESCENDING"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.6.0/docs/resources/datastore_index#direction DatastoreIndex#direction}
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// The property name to index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.6.0/docs/resources/datastore_index#name DatastoreIndex#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

