// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package set

import (
	"sort"
	"sync"
)

// 参考： https://segmentfault.com/a/1190000010294041

type IntSet struct {
	sync.RWMutex
	m map[int]bool
}

// 新建集合对象
// 可以传入初始元素
func New(items ...int) *IntSet {
	s := &IntSet{
		m: make(map[int]bool, len(items)),
	}
	s.Add(items...)
	return s
}

// 创建副本
func (s *IntSet) Duplicate() *IntSet {
	s.Lock()
	defer s.Unlock()
	r := &IntSet{m: make(map[int]bool, len(s.m))}
	for e := range s.m {
		r.m[e] = true
	}
	return r
}

// 添加元素
func (s *IntSet) Add(items ...int) {
	s.Lock()
	defer s.Unlock()
	for _, v := range items {
		s.m[v] = true
	}
}

// 删除元素
func (s *IntSet) Remove(items ...int) {
	s.Lock()
	defer s.Unlock()
	for _, v := range items {
		delete(s.m, v)
	}
}

// 判断元素是否存在
func (s *IntSet) Has(items ...int) bool {
	s.RLock()
	defer s.RUnlock()
	for _, v := range items {
		if _, ok := s.m[v]; !ok {
			return false
		}
	}
	return true
}

// 统计元素个数
func (s *IntSet) Count() int {
	s.Lock()
	defer s.Unlock()
	return len(s.m)
}

// 清空集合
func (s *IntSet) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[int]bool{}
}

// 空集合判断
func (s *IntSet) Empty() bool {
	s.RLock()
	defer s.RUnlock()
	return len(s.m) == 0
}

// 获取元素列表（无序）
func (s *IntSet) List() []int {
	s.RLock()
	defer s.RUnlock()
	list := make([]int, 0, len(s.m))
	for item := range s.m {
		list = append(list, item)
	}
	return list
}

// 获取元素列表（有序）
func (s *IntSet) SortedList() []int {
	s.RLock()
	defer s.RUnlock()
	list := make([]int, 0, len(s.m))
	for item := range s.m {
		list = append(list, item)
	}
	sort.Ints(list)
	return list
}

// 并集
// 获取 s 与参数的并集，结果存入 s
func (s *IntSet) Union(sets ...*IntSet) {
	// 为了防止多例程死锁，不能同时锁定两个集合
	// 所以这里没有锁定 s，而是创建了一个临时集合
	r := s.Duplicate()
	// 获取并集
	for _, set := range sets {
		set.Lock()
		for e := range set.m {
			r.m[e] = true
		}
		set.Unlock()
	}
	// 将结果转入 s
	s.Lock()
	defer s.Unlock()
	s.m = map[int]bool{}
	for e := range r.m {
		s.m[e] = true
	}
}

// 并集（函数）
// 获取所有参数的并集，并返回
func Union(sets ...*IntSet) *IntSet {
	// 处理参数数量
	if len(sets) == 0 {
		return New()
	} else if len(sets) == 1 {
		return sets[0]
	}
	// 获取并集
	r := sets[0].Duplicate()
	for _, set := range sets[1:] {
		set.Lock()
		for e := range set.m {
			r.m[e] = true
		}
		set.Unlock()
	}
	return r
}

// 差集
// 获取 s 与所有参数的差集，结果存入 s
func (s *IntSet) Minus(sets ...*IntSet) {
	// 为了防止多例程死锁，不能同时锁定两个集合
	// 所以这里没有锁定 s，而是创建了一个临时集合
	r := s.Duplicate()
	// 获取差集
	for _, set := range sets {
		set.Lock()
		for e := range set.m {
			delete(r.m, e)
		}
		set.Unlock()
	}
	// 将结果转入 s
	s.Lock()
	defer s.Unlock()
	s.m = map[int]bool{}
	for e := range r.m {
		s.m[e] = true
	}
}

// 差集（函数）
// 获取第 1 个参数与其它参数的差集，并返回
func Minus(sets ...*IntSet) *IntSet {
	// 处理参数数量
	if len(sets) == 0 {
		return New()
	} else if len(sets) == 1 {
		return sets[0]
	}
	// 获取差集
	r := sets[0].Duplicate()
	for _, set := range sets[1:] {
		for e := range set.m {
			delete(r.m, e)
		}
	}
	return r
}

// 交集
// 获取 s 与其它参数的交集，结果存入 s
func (s *IntSet) Intersect(sets ...*IntSet) {
	// 为了防止多例程死锁，不能同时锁定两个集合
	// 所以这里没有锁定 s，而是创建了一个临时集合
	r := s.Duplicate()
	// 获取交集
	for _, set := range sets {
		set.Lock()
		for e := range r.m {
			if _, ok := set.m[e]; !ok {
				delete(r.m, e)
			}
		}
		set.Unlock()
	}
	// 将结果转入 s
	s.Lock()
	defer s.Unlock()
	s.m = map[int]bool{}
	for e := range r.m {
		s.m[e] = true
	}
}

// 交集（函数）
// 获取所有参数的交集，并返回
func Intersect(sets ...*IntSet) *IntSet {
	// 处理参数数量
	if len(sets) == 0 {
		return New()
	} else if len(sets) == 1 {
		return sets[0]
	}
	// 获取交集
	r := sets[0].Duplicate()
	for _, set := range sets[1:] {
		for e := range r.m {
			if _, ok := set.m[e]; !ok {
				delete(r.m, e)
			}
		}
	}
	return r
}

// 补集
// 获取 s 相对于 full 的补集，结果存入 s
func (s *IntSet) Complement(full *IntSet) {
	r := full.Duplicate()
	s.Lock()
	defer s.Unlock()
	// 获取补集
	for e := range s.m {
		delete(r.m, e)
	}
	// 将结果转入 s
	s.m = map[int]bool{}
	for e := range r.m {
		s.m[e] = true
	}
}

// 补集（函数）
// 获取 sub 相对于 full 的补集，并返回
func Complement(sub, full *IntSet) *IntSet {
	r := full.Duplicate()
	sub.Lock()
	defer sub.Unlock()
	for e := range sub.m {
		delete(r.m, e)
	}
	return r
}
