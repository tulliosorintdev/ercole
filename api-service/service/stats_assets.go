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
	"time"

	"github.com/ercole-io/ercole/utils"
)

// GetTotalTechnologiesComplianceStats return the total compliance of all technologies
func (as *APIService) GetTotalTechnologiesComplianceStats(location string, environment string, olderThan time.Time) (map[string]interface{}, utils.AdvancedErrorInterface) {
	data, err := as.ListTechnologies("", false, location, environment, olderThan)
	if err != nil {
		return nil, err
	}

	hostsCount, err := as.GetHostsCountStats(location, environment, olderThan)
	if err != nil {
		return nil, err
	}

	totalUsed := float64(0.0)
	totalCount := float64(0.0)
	totalTotalCost := float64(0.0)
	totalPaidCost := float64(0.0)

	for _, ass := range data {
		totalUsed += ass.Used
		totalCount += ass.Count
		totalTotalCost += ass.TotalCost
		totalPaidCost += ass.PaidCost
	}
	if totalUsed == 0 {
		totalUsed = 1
		totalCount = 1
	}

	return map[string]interface{}{
		"unpaidDues": totalTotalCost - totalPaidCost,
		"compliance": totalCount / totalUsed,
		"hostsCount": hostsCount,
	}, nil
}
