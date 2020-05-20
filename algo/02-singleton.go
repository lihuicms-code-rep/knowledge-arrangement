package main

import "sync"

//题目2:实现Singleton模式
//设计一个类,我们只能生成该类的一个实例
//涉及到的内容:单例模式
/*
分四步进行解决
1.问题:设计一个类,只能生成该类的一个实例
2.分析:
定义：“一个类有且仅有一个实例，并且自行实例化向整个系统提供。”
特点：
单例类只能有一个实例。
单例类必须自己创建自己的唯一实例。
单例类必须给所有其他对象提供这一实例。
要解决的问题：避免一个全局使用的类频繁地创建与销毁；保证一个类仅有一个实例，并提供一个访问它的全局访问点。

在 go 语言中，主要的实现方式有以下四种：
懒汉式：线程(不)安全，需要实例的时候才会主动去创建实例，并且可以通过 Sync.Mutex 加锁来实现线程安全。
双重检查：线程安全，在懒汉式（线程安全）的基础上再进行忧化，在未创建实例的时候才会加锁。保证线程安全同时不影响性能。
饿汉式：线程安全，导入包的同时会创建实例，持续占用内存。
sync.Once：通过 sync.Once 来确保创建对象的方法只执行一次，Once 内部也是实现的双重检查。
 */

type Instance struct {
	name string
}

var instance *Instance //私有实例指针
var lock sync.Mutex    //锁

//懒汉式,线程不安全
//只适用于单线程的环境
//比如两个go或者线程同时判断到==nil,那么就会被创建2次
func GetInstance1() *Instance {
	if instance == nil {
		instance = new(Instance)
	}

	return instance
}

//懒汉式,线程安全,
//通过使用Sync.Mutex同步锁来实现
//缺点:加锁是一个非常耗时的过程
func GetInstance2() *Instance {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = new(Instance)
	}

	return instance
}


//双重检查
//实际上是在上面加锁的方法上提升效率
//如果实例已经创建了则不需要加锁

func GetInstance3() *Instance {
	if instance == nil {
		lock.Lock()
		if instance == nil {
			instance = new(Instance)
		}
		lock.Unlock()
	}

	return instance
}

//饿汉式,导入包的同时创建实例,线程安全

//通过sync.Once
var once sync.Once

func GetInstance4() *Instance {
	once.Do(func() {
		instance = new(Instance)
	})

	return instance
}
