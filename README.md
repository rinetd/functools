##  Functools

[![Build Status](https://travis-ci.org/pytool/functools.svg?branch=master)](https://travis-ci.org/pytool/functools)

Functools is a simple Golang library including some commonly used functional programming tools. 
Reference [roman-kachanovsky/go-built-in]

**Features**

* High level functions such as Apply/Reduce/Filter etc.
* Rust-style Option type

**Install**

`go get github.com/pytool/functools`

## Usage
```
package main

import (
	"log"

	. "github.com/pytool/functools"
)

func PartialFun() {
	sum := func(a, b int) int { return a * b }
	sum10 := Partial(sum, 10)
	result := sum10(10)
	log.Println(result)
}
```
**Partial** 偏函数的功能就是：把一个函数的某些参数给固定住，返回一个新的函数 
```py
multiply(x, y)；
double = partial(multiply, y=2)；
double(3)
```

**Apply/Map** :对 sequence 中的 item 依次执行 function(item)，并将结果组成一个 List 返回 

**Reduce** 先将 sequence 的前两个 item 传给 function，即 function(item1, item2)，函数的返回值和 sequence 的下一个 item 再传给 function，即 function(function(item1, item2), item3)，如此迭代，直到 sequence 没有元素，如果有 initial，则作为初始值调用。
    `reduece(f, [x1, x2, x3, x4]) = f(f(f(x1, x2), x3), x4)`

**Filter** 将 function 依次作用于 sequnce 的每个 item，即 function(item)，将返回值为 True 的 item 组成一个 List 返回。
**Zip** 将两个一维的Slice合并成一个二维的Slice

**All** 必须所有的元素都为真 bool：true int:!0 
**Any** 只要有一个元素为真 
**Cmp** 比较两个元素的大小，结果为int8 [> 1] [= 0] [< -1] 
**ToBool** 将元素转换成Bool类型

**Sum** 计算数值型元素的总和
**Avg** 计算数值型元素的平均值
**Max** 计算数值型元素的最大值
**Min** 计算数值型元素的最小值

