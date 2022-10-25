//go:build no_runtime_type_checking

package dataplexlake

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataplexLakeAssetStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataplexLakeAssetStatusList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataplexLakeAssetStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataplexLakeAssetStatusList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataplexLakeAssetStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataplexLakeAssetStatusListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

