//go:build no_runtime_type_checking

package dnspolicy

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DnsPolicyNetworksList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DnsPolicyNetworksList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DnsPolicyNetworksList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DnsPolicyNetworksList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DnsPolicyNetworksList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DnsPolicyNetworksList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDnsPolicyNetworksListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

