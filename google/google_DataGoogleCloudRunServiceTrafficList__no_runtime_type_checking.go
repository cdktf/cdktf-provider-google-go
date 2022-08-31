//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt google Provider for Terraform CDK (cdktf)
package google

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataGoogleCloudRunServiceTrafficList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataGoogleCloudRunServiceTrafficList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataGoogleCloudRunServiceTrafficList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataGoogleCloudRunServiceTrafficList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataGoogleCloudRunServiceTrafficList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataGoogleCloudRunServiceTrafficListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

