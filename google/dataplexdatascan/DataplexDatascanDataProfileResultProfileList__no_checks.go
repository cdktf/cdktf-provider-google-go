//go:build no_runtime_type_checking

package dataplexdatascan

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataplexDatascanDataProfileResultProfileList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataplexDatascanDataProfileResultProfileList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataplexDatascanDataProfileResultProfileList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataplexDatascanDataProfileResultProfileList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataplexDatascanDataProfileResultProfileList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataplexDatascanDataProfileResultProfileListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
