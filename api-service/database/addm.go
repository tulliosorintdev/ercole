// Copyright (c) 2019 Sorint.lab S.p.A.
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

package database

import (
	"context"
	"regexp"

	"github.com/amreo/ercole-services/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// SearchCurrentAddms search current addms
func (md *MongoDatabase) SearchCurrentAddms(keywords []string, sortBy string, sortDesc bool, page int, pageSize int) ([]interface{}, utils.AdvancedErrorInterface) {
	var out []interface{}
	var quotedKeywords []string
	for _, k := range keywords {
		quotedKeywords = append(quotedKeywords, regexp.QuoteMeta(k))
	}
	//Find the matching hostdata
	cur, err := md.Client.Database(md.Config.Mongodb.DBName).Collection("currentDatabases").Aggregate(
		context.TODO(),
		bson.A{
			bson.M{"$project": bson.M{
				"hostname":       true,
				"location":       true,
				"environment":    true,
				"created_at":     true,
				"database.name":  true,
				"database.addms": true,
			}},
			bson.M{"$unwind": "$database.addms"},
			bson.M{"$project": bson.M{
				"hostname":       true,
				"location":       true,
				"environment":    true,
				"created_at":     true,
				"dbname":         "$database.name",
				"action":         "$database.addms.action",
				"benefit":        "$database.addms.benefit",
				"finding":        "$database.addms.finding",
				"recommendation": "$database.addms.recommendation",
			}},
			optionalSortingStep(sortBy, sortDesc),
			optionalPagingStep(page, pageSize),
		},
	)
	if err != nil {
		return nil, utils.NewAdvancedErrorPtr(err, "DB ERROR")
	}

	//Decode the documents
	for cur.Next(context.TODO()) {
		var item map[string]interface{}
		if cur.Decode(&item) != nil {
			return nil, utils.NewAdvancedErrorPtr(err, "Decode ERROR")
		}
		out = append(out, &item)
	}
	return out, nil
}