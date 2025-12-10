// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package sqldatabaseinstance

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SqlDatabaseInstanceSettingsConnectionPoolConfigList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SqlDatabaseInstanceSettingsConnectionPoolConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SqlDatabaseInstanceSettingsConnectionPoolConfigList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SqlDatabaseInstanceSettingsConnectionPoolConfigList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SqlDatabaseInstanceSettingsConnectionPoolConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SqlDatabaseInstanceSettingsConnectionPoolConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SqlDatabaseInstanceSettingsConnectionPoolConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSqlDatabaseInstanceSettingsConnectionPoolConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

