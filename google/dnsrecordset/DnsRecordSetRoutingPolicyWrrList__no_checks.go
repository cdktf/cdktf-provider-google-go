//go:build no_runtime_type_checking

package dnsrecordset

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DnsRecordSetRoutingPolicyWrrList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DnsRecordSetRoutingPolicyWrrList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyWrrList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyWrrList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyWrrList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyWrrList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDnsRecordSetRoutingPolicyWrrListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

