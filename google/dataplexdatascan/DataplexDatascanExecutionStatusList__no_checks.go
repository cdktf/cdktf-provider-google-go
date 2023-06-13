//go:build no_runtime_type_checking

package dataplexdatascan

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataplexDatascanExecutionStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataplexDatascanExecutionStatusList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataplexDatascanExecutionStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataplexDatascanExecutionStatusList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataplexDatascanExecutionStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataplexDatascanExecutionStatusListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

