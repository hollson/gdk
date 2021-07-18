// Copyright 2021 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goox

import (
	"fmt"
	"testing"
)

func TestTrimSpace(t *testing.T) {
	fmt.Printf("[%s]\n", TrimSpace(`
	select *   
  from  \n\t
wherw \n 1=1
order by \t\t\t\ id desc       `))

	fmt.Printf("[%s]\n", TrimSpace("select \n * from \t\r student  where 1=1        \n"))
}

func TestPascal(t *testing.T) {
	fmt.Println(Pascal("_user_name"))
	fmt.Println(Pascal("_username"))
	fmt.Println(Pascal("_last_user_name"))
	fmt.Println(Pascal("_last_100_user_name"))
	fmt.Println(Pascal("user_name"))
	fmt.Println(Pascal("User_Name"))
	fmt.Println(Pascal("USERNAME"))
	fmt.Println(Pascal("username"))
	fmt.Println(Pascal("user__name"))
	fmt.Println(Pascal("_user_name_"))
	fmt.Println(Pascal("_user_name_99"))
}


func TestHasSub(t *testing.T) {
	fmt.Println(HasAny("datetime", "date", "time")) // true
	fmt.Println(HasAny("datetime", "date"))         // true
	fmt.Println(HasAny("datetime", "time", "abc"))  // true
	fmt.Println(HasAny("datetime", "time", "xyz"))  // true
	fmt.Println(HasAny("datetime", "xyz"))          // false
	fmt.Println(HasAny("datetime", "xxx", "yyy"))   // false
}
