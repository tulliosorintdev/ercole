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

func init() {
	searchSegmentAdvisorsCmd := simpleAPIRequestCommand("search-segment-advisors",
		"Search current segment advisors",
		`search-segment-advisors search the most matching segment advisors to the arguments`,
		false, false, false, true, true, true, false,
		"/segment-advisors",
		"Failed to search segment advisors data: %v\n",
		"Failed to search segment advisors data(Status: %d): %s\n",
	)

	apiCmd.AddCommand(searchSegmentAdvisorsCmd)
}