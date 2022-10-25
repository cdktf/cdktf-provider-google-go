//go:build no_runtime_type_checking

package memcacheinstance

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MemcacheInstanceMaintenanceScheduleList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MemcacheInstanceMaintenanceScheduleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MemcacheInstanceMaintenanceScheduleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MemcacheInstanceMaintenanceScheduleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MemcacheInstanceMaintenanceScheduleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMemcacheInstanceMaintenanceScheduleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

