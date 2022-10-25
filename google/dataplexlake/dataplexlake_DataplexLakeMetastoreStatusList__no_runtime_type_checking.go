//go:build no_runtime_type_checking

package dataplexlake

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataplexLakeMetastoreStatusList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataplexLakeMetastoreStatusList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataplexLakeMetastoreStatusList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataplexLakeMetastoreStatusList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataplexLakeMetastoreStatusList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataplexLakeMetastoreStatusListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

