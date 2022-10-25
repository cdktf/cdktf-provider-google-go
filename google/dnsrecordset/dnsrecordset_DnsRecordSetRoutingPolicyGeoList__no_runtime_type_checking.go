//go:build no_runtime_type_checking

package dnsrecordset

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DnsRecordSetRoutingPolicyGeoList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DnsRecordSetRoutingPolicyGeoList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyGeoList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyGeoList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyGeoList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DnsRecordSetRoutingPolicyGeoList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDnsRecordSetRoutingPolicyGeoListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

