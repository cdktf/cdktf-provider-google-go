//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt google Provider for Terraform CDK (cdktf)
package google

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataGoogleComputeInstanceTemplateShieldedInstanceConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataGoogleComputeInstanceTemplateShieldedInstanceConfigList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateShieldedInstanceConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateShieldedInstanceConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataGoogleComputeInstanceTemplateShieldedInstanceConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataGoogleComputeInstanceTemplateShieldedInstanceConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

