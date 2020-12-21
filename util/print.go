// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package util

import (
    "encoding/json"
    "fmt"
)

// 将结构体o以json的格式输出
//  实现了类strings.String接口
func JsonFormat(o interface{}) string {
    dump, err := json.MarshalIndent(o, "", "\t")
    if err != nil {
        return fmt.Errorf("json format convert error,%v",err.Error()).Error()
    }
    return string(dump)
}
