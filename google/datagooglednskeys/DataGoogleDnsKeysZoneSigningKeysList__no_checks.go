// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datagooglednskeys

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataGoogleDnsKeysZoneSigningKeysList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataGoogleDnsKeysZoneSigningKeysList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataGoogleDnsKeysZoneSigningKeysList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataGoogleDnsKeysZoneSigningKeysList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataGoogleDnsKeysZoneSigningKeysList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataGoogleDnsKeysZoneSigningKeysListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

