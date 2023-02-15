//go:build no_runtime_type_checking

package filestoreinstance

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FilestoreInstanceNetworksList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FilestoreInstanceNetworksList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FilestoreInstanceNetworksList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FilestoreInstanceNetworksList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FilestoreInstanceNetworksList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FilestoreInstanceNetworksList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFilestoreInstanceNetworksListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

