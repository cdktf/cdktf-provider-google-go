//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt google Provider for Terraform CDK (cdktf)
package google

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataGoogleCloudfunctionsFunctionEventTriggerList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataGoogleCloudfunctionsFunctionEventTriggerList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataGoogleCloudfunctionsFunctionEventTriggerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataGoogleCloudfunctionsFunctionEventTriggerList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataGoogleCloudfunctionsFunctionEventTriggerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataGoogleCloudfunctionsFunctionEventTriggerListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

