//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package dataplexasset

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataplexAssetDiscoveryStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataplexAssetDiscoveryStatusList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataplexAssetDiscoveryStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataplexAssetDiscoveryStatusList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataplexAssetDiscoveryStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataplexAssetDiscoveryStatusListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

