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

package service

import (
	"testing"

	"github.com/ercole-io/ercole/utils"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListTechnologies_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	db := NewMockMongoDatabaseInterface(mockCtrl)
	as := APIService{
		Database: db,
	}

	expectedRes := []map[string]interface{}{
		{
			"compliance": 1.0,
			"unpaidDues": 0,
			"product":    "Oracle/Exadata",
			"hostsCount": 2,
		},
		{
			"compliance": 7.0 / 10,
			"unpaidDues": 45,
			"product":    "Oracle/Database",
			"hostsCount": 8,
		},
	}

	getTechnologiesUsageRes := map[string]float64{
		"Oracle/Database_hostsCount": 8,
		"Oracle/Exadata":             2,
	}
	db.EXPECT().
		GetTechnologiesUsage("Italy", "PROD", utils.P("2020-12-05T14:02:03Z")).
		Return(getTechnologiesUsageRes, nil)
	listLicensesRes := []interface{}{
		map[string]interface{}{
			"compliance":       false,
			"count":            4,
			"used":             4,
			"_id":              "Partitioning",
			"totalCost":        40,
			"paidCost":         40,
			"costPerProcessor": 10,
		},
		map[string]interface{}{
			"compliance":       false,
			"count":            3,
			"used":             6,
			"_id":              "Diagnostics Pack",
			"totalCost":        90,
			"paidCost":         45,
			"costPerProcessor": 15,
		},
		map[string]interface{}{
			"compliance":       true,
			"count":            5,
			"used":             0,
			"_id":              "Advanced Analytics",
			"totalCost":        0,
			"paidCost":         5,
			"costPerProcessor": 1,
		},
	}
	db.EXPECT().
		ListLicenses(false, "", false, -1, -1, "Italy", "PROD", utils.P("2020-12-05T14:02:03Z")).
		Return(listLicensesRes, nil)

	res, err := as.ListTechnologies(
		"Count", true,
		"Italy", "PROD", utils.P("2020-12-05T14:02:03Z"),
	)
	require.NoError(t, err)
	assert.JSONEq(t, utils.ToJSON(expectedRes), utils.ToJSON(res))
}

func TestListTechnologies_SuccessEmpty(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	db := NewMockMongoDatabaseInterface(mockCtrl)
	as := APIService{
		Database: db,
	}

	expectedRes := []map[string]interface{}{
		{"compliance": 1, "hostsCount": 0, "product": "Oracle/Exadata", "unpaidDues": 0},
		{"compliance": 1, "hostsCount": 0, "product": "Oracle/Database", "unpaidDues": 0},
	}

	getTechnologiesUsageRes := map[string]float64{}
	db.EXPECT().
		GetTechnologiesUsage("Italy", "PROD", utils.P("2020-12-05T14:02:03Z")).
		Return(getTechnologiesUsageRes, nil)
	listLicensesRes := []interface{}{
		map[string]interface{}{
			"compliance": false,
			"count":      10,
			"used":       0,
			"_id":        "Partitioning",
		},
	}
	db.EXPECT().
		ListLicenses(false, "", false, -1, -1, "Italy", "PROD", utils.P("2020-12-05T14:02:03Z")).
		Return(listLicensesRes, nil)

	res, err := as.ListTechnologies(
		"Count", true,
		"Italy", "PROD", utils.P("2020-12-05T14:02:03Z"),
	)

	require.NoError(t, err)
	assert.JSONEq(t, utils.ToJSON(expectedRes), utils.ToJSON(res))
}

func TestListTechnologies_FailInternalServerError1(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	db := NewMockMongoDatabaseInterface(mockCtrl)
	as := APIService{
		Database: db,
	}

	db.EXPECT().
		GetTechnologiesUsage("Italy", "PROD", utils.P("2020-12-05T14:02:03Z")).
		Return(nil, aerrMock)

	_, err := as.ListTechnologies(
		"Count", true,
		"Italy", "PROD", utils.P("2020-12-05T14:02:03Z"),
	)

	require.Equal(t, aerrMock, err)
}

func TestListTechnologies_FailInternalServerError2(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	db := NewMockMongoDatabaseInterface(mockCtrl)
	as := APIService{
		Database: db,
	}

	getTechnologiesUsageRes := map[string]float64{}
	db.EXPECT().
		GetTechnologiesUsage("Italy", "PROD", utils.P("2020-12-05T14:02:03Z")).
		Return(getTechnologiesUsageRes, nil)
	db.EXPECT().
		ListLicenses(false, "", false, -1, -1, "Italy", "PROD", utils.P("2020-12-05T14:02:03Z")).
		Return(nil, aerrMock)

	_, err := as.ListTechnologies(
		"Count", true,
		"Italy", "PROD", utils.P("2020-12-05T14:02:03Z"),
	)

	require.Equal(t, aerrMock, err)
}
