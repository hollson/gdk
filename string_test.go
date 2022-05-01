package gdk

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

func ExampleCharCount() {
	str := "Hello, 世界"
	fmt.Println("bytes =", len(str))
	fmt.Println("runes =", CharCount(str))

	// Output:
	//
	// bytes = 13
	// runes = 9
}

// 遍历字符串
func TestStringRange(t *testing.T) {
	str := "Hello,世界"

	for i := 0; i < len(str); i++ {
		ch := str[i]
		fmt.Printf(string(ch))
	}

	fmt.Println()
	for _, ch1 := range str {
		fmt.Printf(string(ch1))
	}
}
