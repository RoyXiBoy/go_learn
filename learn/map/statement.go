package main

//声明
//var map1 map[keytype]valuetype
var map1 map[string]int

//未初始化的 map 的值是 nil。

/*
map 可以用 {key1: val1, key2: val2} 的描述方法来初始化，就像数组和结构体一样。

map 是 引用类型 的： 内存用 make 方法来分配。

map 的初始化：var map1 = make(map[keytype]valuetype)。

或者简写为：map1 := make(map[keytype]valuetype)。

不要使用 new，永远用 make 来构造 map

注意 如果你错误的使用 new() 分配了一个引用对象，你会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址：

 */
