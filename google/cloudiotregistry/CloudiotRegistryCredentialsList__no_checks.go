//go:build no_runtime_type_checking

package cloudiotregistry

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudiotRegistryCredentialsList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudiotRegistryCredentialsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudiotRegistryCredentialsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudiotRegistryCredentialsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudiotRegistryCredentialsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudiotRegistryCredentialsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudiotRegistryCredentialsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

