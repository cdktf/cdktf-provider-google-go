//go:build no_runtime_type_checking

package dataplexzone

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataplexZoneAssetStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataplexZoneAssetStatusList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataplexZoneAssetStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataplexZoneAssetStatusList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataplexZoneAssetStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataplexZoneAssetStatusListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

