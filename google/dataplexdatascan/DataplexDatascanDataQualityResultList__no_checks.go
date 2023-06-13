//go:build no_runtime_type_checking

package dataplexdatascan

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataplexDatascanDataQualityResultList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataplexDatascanDataQualityResultList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataplexDatascanDataQualityResultList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataplexDatascanDataQualityResultList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataplexDatascanDataQualityResultList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataplexDatascanDataQualityResultListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

