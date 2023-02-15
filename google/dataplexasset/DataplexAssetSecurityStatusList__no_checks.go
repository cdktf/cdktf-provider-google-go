//go:build no_runtime_type_checking

package dataplexasset

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataplexAssetSecurityStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataplexAssetSecurityStatusList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataplexAssetSecurityStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataplexAssetSecurityStatusList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataplexAssetSecurityStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataplexAssetSecurityStatusListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

