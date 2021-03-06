// Copyright (c) 2020 Sorint.lab S.p.A.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Package service is a package that provides methods for querying data
package service

import (
	"strings"
	"time"

	"github.com/ercole-io/ercole/utils"
)

// SearchOracleDatabaseAddms search addms
func (as *APIService) SearchOracleDatabaseAddms(search string, sortBy string, sortDesc bool, page int, pageSize int, location string, environment string, olderThan time.Time) ([]map[string]interface{}, utils.AdvancedErrorInterface) {
	return as.Database.SearchOracleDatabaseAddms(strings.Split(search, " "), sortBy, sortDesc, page, pageSize, location, environment, olderThan)
}

// SearchOracleDatabaseSegmentAdvisors search segment advisors
func (as *APIService) SearchOracleDatabaseSegmentAdvisors(search string, sortBy string, sortDesc bool, page int, pageSize int, location string, environment string, olderThan time.Time) ([]map[string]interface{}, utils.AdvancedErrorInterface) {
	return as.Database.SearchOracleDatabaseSegmentAdvisors(strings.Split(search, " "), sortBy, sortDesc, page, pageSize, location, environment, olderThan)
}

// SearchOracleDatabasePatchAdvisors search patch advisors
func (as *APIService) SearchOracleDatabasePatchAdvisors(search string, sortBy string, sortDesc bool, page int, pageSize int, windowTime time.Time, location string, environment string, olderThan time.Time, status string) ([]map[string]interface{}, utils.AdvancedErrorInterface) {
	return as.Database.SearchOracleDatabasePatchAdvisors(strings.Split(search, " "), sortBy, sortDesc, page, pageSize, windowTime, location, environment, olderThan, status)
}

// SearchOracleDatabases search databases
func (as *APIService) SearchOracleDatabases(full bool, search string, sortBy string, sortDesc bool, page int, pageSize int, location string, environment string, olderThan time.Time) ([]map[string]interface{}, utils.AdvancedErrorInterface) {
	return as.Database.SearchOracleDatabases(full, strings.Split(search, " "), sortBy, sortDesc, page, pageSize, location, environment, olderThan)
}

// ListLicenses list licenses
func (as *APIService) ListLicenses(full bool, sortBy string, sortDesc bool, page int, pageSize int, location string, environment string, olderThan time.Time) ([]interface{}, utils.AdvancedErrorInterface) {
	return as.Database.ListLicenses(full, sortBy, sortDesc, page, pageSize, location, environment, olderThan)
}

// GetLicense return the license specified in the name param
func (as *APIService) GetLicense(name string, olderThan time.Time) (interface{}, utils.AdvancedErrorInterface) {
	return as.Database.GetLicense(name, olderThan)
}

// SetLicenseCount set the count of a certain license
func (as *APIService) SetLicenseCount(name string, count int) utils.AdvancedErrorInterface {
	return as.Database.SetLicenseCount(name, count)
}

// SetLicenseCostPerProcessor set the cost per processor of a certain license
func (as *APIService) SetLicenseCostPerProcessor(name string, costPerProcessor float64) utils.AdvancedErrorInterface {
	return as.Database.SetLicenseCostPerProcessor(name, costPerProcessor)
}

// SetLicenseUnlimitedStatus set the unlimited status of a certain license
func (as *APIService) SetLicenseUnlimitedStatus(name string, unlimitedStatus bool) utils.AdvancedErrorInterface {
	return as.Database.SetLicenseUnlimitedStatus(name, unlimitedStatus)
}

// SetLicensesCount set the count of all licenses in newLicenses
// It assumes that newLicenses maps contain the string _id and the int Count
func (as *APIService) SetLicensesCount(newLicenses []map[string]interface{}) utils.AdvancedErrorInterface {
	for _, lic := range newLicenses {
		err := as.Database.SetLicenseCount(lic["_id"].(string), lic["count"].(int))
		if err != nil {
			return err
		}
	}
	return nil
}
