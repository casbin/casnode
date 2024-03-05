// Copyright 2021 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package casdoor

import (
	"fmt"

	"github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

func GetUsers() []*casdoorsdk.User {
	if adapter != nil {
		return getUsers()
	} else {
		users, err := casdoorsdk.GetUsers()
		if err != nil {
			panic(err)
		}

		return users
	}
}

func GetSortedUsers(sorter string, limit int) []*casdoorsdk.User {
	if adapter != nil {
		return getSortedUsers(sorter, limit)
	} else {
		users, err := casdoorsdk.GetSortedUsers(sorter, limit)
		if err != nil {
			panic(err)
		}

		return users
	}
}

func GetUserCount() int {
	if adapter != nil {
		return getUserCount()
	} else {
		count, err := casdoorsdk.GetUserCount("")
		if err != nil {
			panic(err)
		}

		return count
	}
}

func GetOnlineUserCount() int {
	if adapter != nil {
		return getOnlineUserCount()
	} else {
		count, err := casdoorsdk.GetUserCount("1")
		if err != nil {
			fmt.Printf("GetOnlineUserCount() error, %s\n", err.Error())
			return -1
		}

		return count
	}
}

func GetUserByEmail(email string) *casdoorsdk.User {
	if adapter != nil {
		return getUserByEmail(email)
	} else {
		user, err := casdoorsdk.GetUserByEmail(email)
		if err != nil {
			panic(err)
		}

		return user
	}
}
