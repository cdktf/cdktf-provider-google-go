//go:build no_runtime_type_checking

package certificatemanagerdnsauthorization

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CertificateManagerDnsAuthorizationDnsResourceRecordList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CertificateManagerDnsAuthorizationDnsResourceRecordList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CertificateManagerDnsAuthorizationDnsResourceRecordList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CertificateManagerDnsAuthorizationDnsResourceRecordList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CertificateManagerDnsAuthorizationDnsResourceRecordList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCertificateManagerDnsAuthorizationDnsResourceRecordListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

