//go:build no_runtime_type_checking

package dnsresponsepolicy

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DnsResponsePolicyNetworksList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DnsResponsePolicyNetworksList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DnsResponsePolicyNetworksList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DnsResponsePolicyNetworksList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DnsResponsePolicyNetworksList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DnsResponsePolicyNetworksList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDnsResponsePolicyNetworksListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

