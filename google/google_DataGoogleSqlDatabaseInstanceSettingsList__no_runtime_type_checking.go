//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt google Provider for Terraform CDK (cdktf)
package google

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataGoogleSqlDatabaseInstanceSettingsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataGoogleSqlDatabaseInstanceSettingsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstanceSettingsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstanceSettingsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataGoogleSqlDatabaseInstanceSettingsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataGoogleSqlDatabaseInstanceSettingsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

