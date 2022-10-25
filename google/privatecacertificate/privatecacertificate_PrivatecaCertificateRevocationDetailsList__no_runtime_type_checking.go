//go:build no_runtime_type_checking

package privatecacertificate

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrivatecaCertificateRevocationDetailsList) validateGetParameters(index *float64) error {
	return nil
}

func (p *jsiiProxy_PrivatecaCertificateRevocationDetailsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_PrivatecaCertificateRevocationDetailsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_PrivatecaCertificateRevocationDetailsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_PrivatecaCertificateRevocationDetailsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewPrivatecaCertificateRevocationDetailsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

