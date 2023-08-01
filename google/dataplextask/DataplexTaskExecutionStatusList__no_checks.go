//go:build no_runtime_type_checking

package dataplextask

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataplexTaskExecutionStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataplexTaskExecutionStatusList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataplexTaskExecutionStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataplexTaskExecutionStatusList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataplexTaskExecutionStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataplexTaskExecutionStatusListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

