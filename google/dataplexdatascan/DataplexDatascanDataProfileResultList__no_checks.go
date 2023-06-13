//go:build no_runtime_type_checking

package dataplexdatascan

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataplexDatascanDataProfileResultList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataplexDatascanDataProfileResultList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataplexDatascanDataProfileResultList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataplexDatascanDataProfileResultList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataplexDatascanDataProfileResultList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataplexDatascanDataProfileResultListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

