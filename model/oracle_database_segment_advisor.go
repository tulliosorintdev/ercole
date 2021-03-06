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

package model

import (
	"reflect"

	godynstruct "github.com/amreo/go-dyn-struct"
	"go.mongodb.org/mongo-driver/bson"
)

// OracleDatabaseSegmentAdvisor holds information about a segment advisor
type OracleDatabaseSegmentAdvisor struct {
	SegmentOwner   string                 `json:"segmentOwner" bson:"segmentOwner"`
	SegmentName    string                 `json:"segmentName" bson:"segmentName"`
	SegmentType    string                 `json:"segmentType" bson:"segmentType"`
	PartitionName  string                 `json:"partitionName" bson:"partitionName"`
	Reclaimable    float64                `json:"reclaimable" bson:"reclaimable"`
	Recommendation string                 `json:"recommendation" bson:"recommendation"`
	OtherInfo      map[string]interface{} `json:"-" bson:"-"`
}

// MarshalJSON return the JSON rappresentation of this
func (v OracleDatabaseSegmentAdvisor) MarshalJSON() ([]byte, error) {
	return godynstruct.DynMarshalJSON(reflect.ValueOf(v), v.OtherInfo, "OtherInfo")
}

// UnmarshalJSON parse the JSON content in data and set the fields in v appropriately
func (v *OracleDatabaseSegmentAdvisor) UnmarshalJSON(data []byte) error {
	return godynstruct.DynUnmarshalJSON(data, reflect.ValueOf(v), &v.OtherInfo, "OtherInfo")
}

// MarshalBSON return the BSON rappresentation of this
func (v OracleDatabaseSegmentAdvisor) MarshalBSON() ([]byte, error) {
	return godynstruct.DynMarshalBSON(reflect.ValueOf(v), v.OtherInfo, "OtherInfo")
}

// UnmarshalBSON parse the BSON content in data and set the fields in v appropriately
func (v *OracleDatabaseSegmentAdvisor) UnmarshalBSON(data []byte) error {
	return godynstruct.DynUnmarshalBSON(data, reflect.ValueOf(v), &v.OtherInfo, "OtherInfo")
}

// OracleDatabaseSegmentAdvisorBsonValidatorRules contains mongodb validation rules for OracleDatabaseSegmentAdvisor
var OracleDatabaseSegmentAdvisorBsonValidatorRules = bson.M{
	"bsonType": "object",
	"required": bson.A{
		"segmentOwner",
		"segmentName",
		"segmentType",
		"partitionName",
		"reclaimable",
		"recommendation",
	},
	"properties": bson.M{
		"segmentOwner": bson.M{
			"bsonType":  "string",
			"minLength": 1,
			"maxLength": 32,
		},
		"segmentName": bson.M{
			"bsonType":  "string",
			"minLength": 1,
			"maxLength": 32,
		},
		"segmentType": bson.M{
			"bsonType":  "string",
			"minLength": 1,
			"maxLength": 32,
		},
		"partitionName": bson.M{
			"bsonType":  "string",
			"maxLength": 32,
		},
		"reclaimable": bson.M{
			"bsonType": "number",
			"minimum":  0,
		},
		"recommendation": bson.M{
			"bsonType":  "string",
			"minLength": 1,
			"maxLength": 256,
		},
	},
}
