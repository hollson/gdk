// Copyright 2020 Hollson. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package kkmap

import (
    "fmt"
    "sync"
)

// Object with two identity identifiers
type KKobject interface {
    Key1() interface{}
    Key2() interface{}
    fmt.Stringer
}

// map with double key
//  e.g. TcpPacket(uuid,remoteRul)
type Map struct {
    positive map[interface{}]KKobject
    negative map[interface{}]KKobject
    lock     sync.RWMutex
}

func NewMap() *Map {
    return &Map{
        positive: map[interface{}]KKobject{},
        negative: map[interface{}]KKobject{},
        lock:     sync.RWMutex{},
    }
}

// add or modify value of kkmap
func (k *Map) Set(kk KKobject) {
    k.lock.Lock()
    defer k.lock.Unlock()
    k.positive[kk.Key1()] = kk
    k.negative[kk.Key2()] = kk
}

// get value by  key1
func (k *Map) GetByKey1(k1 interface{}) KKobject {
    k.lock.RLock()
    defer k.lock.RUnlock()
    v, ok := k.positive[k1]
    if !ok {
        return nil
    }
    return v
}

// get value by  key2
func (k *Map) GetByKey2(k2 interface{}) KKobject {
    k.lock.RLock()
    defer k.lock.RUnlock()
    v, ok := k.negative[k2]
    if !ok {
        return nil
    }
    return v
}

// Check that the length of key1 and key1 are equal
func (k *Map) Check() bool {
    return len(k.positive) == len(k.negative)
}

// get length of kkmap
func (k *Map) Length() (k1len, k2len int) {
    return len(k.positive), len(k.negative)
}

// delete element of kkmap
func (k *Map) Delete(kk KKobject) {
    k.lock.Lock()
    defer k.lock.Unlock()
    delete(k.positive, kk.Key1())
    delete(k.negative, kk.Key2())
}

// delete element of kkmap by key1
func (k *Map) DeleteByKey1(k1 interface{}) {
    k.lock.Lock()
    defer k.lock.Unlock()
    if v, ok := k.positive[k1]; ok {
        delete(k.negative, v.(KKobject).Key2())
    }
    delete(k.positive, k1)
}

// delete element of kkmap by key2
func (k *Map) DeleteByKey2(k2 interface{}) {
    k.lock.Lock()
    defer k.lock.Unlock()
    if v, ok := k.negative[k2]; ok {
        delete(k.positive, v.(KKobject).Key1())
    }
    delete(k.negative, k2)
}

// range kkmap by key1
func (k *Map) RangeByKey1(f func(KKobject)) {
    for _, v := range k.positive {
        f(v)
    }
}

// range kkmap by key1
func (k *Map) RangeByKey2(f func(KKobject)) {
    for _, v := range k.negative {
        f(v)
    }
}

func (k *Map) String() string {
    return fmt.Sprintf("%v\n%v", k.positive, k.negative)
}
