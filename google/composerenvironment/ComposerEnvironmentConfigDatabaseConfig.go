// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package composerenvironment


type ComposerEnvironmentConfigDatabaseConfig struct {
	// Optional.
	//
	// Cloud SQL machine type used by Airflow database. It has to be one of: db-n1-standard-2, db-n1-standard-4, db-n1-standard-8 or db-n1-standard-16. If not specified, db-n1-standard-2 will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/composer_environment#machine_type ComposerEnvironment#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// Optional. Cloud SQL database preferred zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.30.0/docs/resources/composer_environment#zone ComposerEnvironment#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

