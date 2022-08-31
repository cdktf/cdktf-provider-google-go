//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt google Provider for Terraform CDK (cdktf)
package google

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataGoogleComposerEnvironmentConfigDatabaseConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataGoogleComposerEnvironmentConfigDatabaseConfigList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigDatabaseConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigDatabaseConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataGoogleComposerEnvironmentConfigDatabaseConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataGoogleComposerEnvironmentConfigDatabaseConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

