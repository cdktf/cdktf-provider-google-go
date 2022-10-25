//go:build no_runtime_type_checking

package storageobjectaccesscontrol

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StorageObjectAccessControlProjectTeamList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_StorageObjectAccessControlProjectTeamList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_StorageObjectAccessControlProjectTeamList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_StorageObjectAccessControlProjectTeamList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_StorageObjectAccessControlProjectTeamList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewStorageObjectAccessControlProjectTeamListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

