
[TOC]


# 算法类
- [雪花算法](./algorithm/snowflake.go)


# 代码设计
- [错误重试](./algorithm/retry.go)

func ReverseInts(data []int) {
   for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
      data[i], data[j] = data[j], data[i]
   }
}