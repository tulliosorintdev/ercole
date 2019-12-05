/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/amreo/ercole-services/utils"

	"github.com/spf13/cobra"
)

// searchAlertsCmd represents the getHosts command
var searchAlertsCmd = &cobra.Command{
	Use:   "search-alerts",
	Short: "Search alerts",
	Long:  `search-alerts search the most matching alerts to the arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		params := url.Values{
			"search": []string{strings.Join(args, " ")},
		}

		if sortBy != "" {
			params.Set("sort-by", sortBy)
			params.Set("sort-desc", strconv.FormatBool(sortDesc))
		}

		resp, err := http.Get(
			utils.NewAPIUrl(
				ercoleConfig.APIService.RemoteEndpoint,
				ercoleConfig.APIService.UserUsername,
				ercoleConfig.APIService.UserPassword,
				"/alerts",
				params,
			).String())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to search alerts data: %v\n", err)
			os.Exit(1)
		} else if resp.StatusCode < 200 || resp.StatusCode > 299 {
			out, _ := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			fmt.Fprintf(os.Stderr, "Failed to search alerts data(Status: %d): %s\n", resp.StatusCode, string(out))
			os.Exit(1)
		} else {
			out, _ := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			var res []interface{}
			err = json.Unmarshal(out, &res)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to unmarshal response body: %v (%s)\n", err, string(out))
				os.Exit(1)
			}

			for _, item := range res {
				enc := json.NewEncoder(os.Stdout)
				enc.SetIndent("", "    ")
				enc.Encode(item)
			}
		}

	},
}

func init() {
	apiCmd.AddCommand(searchAlertsCmd)
	searchAlertsCmd.Flags().StringVar(&sortBy, "sort-by", "", "Sort by field")
	searchAlertsCmd.Flags().BoolVar(&sortDesc, "desc-order", false, "Sort descending")
}