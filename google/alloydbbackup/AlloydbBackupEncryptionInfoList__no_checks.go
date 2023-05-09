//go:build no_runtime_type_checking

package alloydbbackup

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlloydbBackupEncryptionInfoList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlloydbBackupEncryptionInfoList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlloydbBackupEncryptionInfoList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlloydbBackupEncryptionInfoList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlloydbBackupEncryptionInfoList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlloydbBackupEncryptionInfoListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

