// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package goox

import (
    "regexp"
)

// 是否匹配IP格式
func IsIP(str string) bool {
    regular := `((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)`
    return regexp.MustCompile(regular).MatchString(str)
}

// 是否匹配邮箱格式
func IsEmail(str string) bool {
    regular := `^[_a-zA-Z0-9-]+(\.[_a-zA-Z0-9-]+)*@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*(\.[a-zA-Z]{2,4})$`
    return regexp.MustCompile(regular).MatchString(str)
}

// 是否匹配电话格式
func IsTel(str string) bool {
    regular := `^(010|02\d{1}|0[3-9]\d{2})-\d{7,9}(-\d+)?$`
    return regexp.MustCompile(regular).MatchString(str)
}

// 是否匹配400电话
func Is400(str string) bool {
    regular := `^400(-\d{3,4}){2}$`
    return regexp.MustCompile(regular).MatchString(str)

}

// 是否匹配手机号格式
func IsPhone(str string) bool {
    regular := `^(\+?86-?)?(18|15|13)[0-9]{9}$`
    return regexp.MustCompile(regular).MatchString(str)
}

// 是否匹配年份格式
//  (?!0000)  闰年:2016-02-29
func IsYMD(str string) bool {
    regular := `^([0-9]{4}-((0[1-9]|1[0-2])-(0[1-9]|1[0-9]|2[0-8])|(0[13-9]|1[0-2])-(29|30)|(0[13578]|1[02])-31)|([0-9]{2}(0[48]|[2468][048]|[13579][26])|(0[48]|[2468][048]|[13579][26])00)-02-29)$`
    return regexp.MustCompile(regular).MatchString(str)
}

// 是否匹配日期格式(AP)
//  hh:mm:ss xx
func IsHMS_APM(str string) bool {
    regular := `(0[1-9]|1[0-2]):[0-5][0-9]:[0-5][0-9] ([AP]M)`
    return regexp.MustCompile(regular).MatchString(str)
}

// 是否匹配日期格式
//  hh:mm:ss
func IsHMS(str string) bool {
    regular := `(0[1-9]|1[0-2]):[0-5][0-9]:[0-5][0-9]`
    return regexp.MustCompile(regular).MatchString(str)
}

// 是否匹配日期时间格式
//  YYYY-MM-DD
//  (?!0000)  闰年:2016-02-29
func IsYMDHMS(str string) bool {
    regular := `^([0-9]{4}-((0[1-9]|1[0-2])-(0[1-9]|1[0-9]|2[0-8])|(0[13-9]|1[0-2])-(29|30)|(0[13578]|1[02])-31)|([0-9]{2}(0[48]|[2468][048]|[13579][26])|(0[48]|[2468][048]|[13579][26])00)-02-29) (0[1-9]|1[0-2]):[0-5][0-9]:[0-5][0-9]$`
    return regexp.MustCompile(regular).MatchString(str)
}

// 是否是数字
func IsNumber(str string) bool {
    regular := `^[0-9]*$`
    return regexp.MustCompile(regular).MatchString(str)
}

// 整数或者小数
func IsFloat(str string) bool {
    regular := `^[0-9]+([.]{0,1}[0-9]+){0,1}$`
    return regexp.MustCompile(regular).MatchString(str)
}

// 是否以特殊符号开头
//  !@#$%^&*()_+-={}|[]\:";'<>?,./ 等
func IsSpecialSymbols(str string) bool {
    regular := `^[!@#$%^&*()_+-={}|[]\:";'<>?,./]+`
    return regexp.MustCompile(regular).MatchString(str)
}

// 是否全是中文
func IsChineseCharacter(str string) bool {
    regular := `^[\u4E00-\u9FA5]+$`
    return regexp.MustCompile(regular).MatchString(str)
}
