// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package sqldatabaseinstance

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SqlDatabaseInstanceDnsNamesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SqlDatabaseInstanceDnsNamesList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SqlDatabaseInstanceDnsNamesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SqlDatabaseInstanceDnsNamesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SqlDatabaseInstanceDnsNamesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SqlDatabaseInstanceDnsNamesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSqlDatabaseInstanceDnsNamesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

