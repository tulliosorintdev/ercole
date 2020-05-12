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

	"github.com/amreo/ercole-services/utils"
)

// GetTotalAssetsComplianceStats return the total compliance of all assets
func (as *APIService) GetTotalAssetsComplianceStats(location string, environment string, olderThan time.Time) (map[string]interface{}, utils.AdvancedErrorInterface) {
	data, err := as.ListAssets("", false, location, environment, olderThan)
	if err != nil {
		return nil, err
	}

	totalUsed := float32(0.0)
	totalCount := float32(0.0)
	totalCompliant := true
	totalCost := float32(0.0)

	for _, ass := range data {
		totalUsed += ass.Used
		totalCount += ass.Count
		totalCost += ass.Cost
		totalCompliant = totalCompliant && ass.Compliance
	}

	return map[string]interface{}{
		"Used":      totalUsed,
		"Count":     totalCount,
		"Compliant": totalCompliant,
		"Cost":      totalCost,
	}, nil
}