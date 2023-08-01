//go:build no_runtime_type_checking

package dataplextask

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataplexTaskExecutionStatusLatestJobList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataplexTaskExecutionStatusLatestJobList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataplexTaskExecutionStatusLatestJobList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataplexTaskExecutionStatusLatestJobList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataplexTaskExecutionStatusLatestJobList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataplexTaskExecutionStatusLatestJobListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

