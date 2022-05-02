#include<iostream>
using namespace std;

// ref: https://zhuanlan.zhihu.com/p/37469260

// 总结：
// - Eager Singleton 虽然是线程安全的，但存在潜在问题；
// - Lazy Singleton通常需要加锁来保证线程安全，但局部静态变量版本在C++11后是线程安全的；
// - 局部静态变量版本（Meyers Singleton）最优雅。


// ============================================================================
// 教学版，即懒汉版（Lazy Singleton）：单例实例在第一次被使用时才进行初始化,
// 这叫做延迟初始化。
// version 1.0
// ============================================================================
class LazySingleton
{
    private:
        static LazySingleton* instance;

    private:
        LazySingleton() {};
        ~LazySingleton() {};
        LazySingleton(const LazySingleton&);
        LazySingleton& operator=(const LazySingleton&);

    public:
        static LazySingleton* getInstance() 
        {
            if(instance == NULL) 
                instance = new LazySingleton();
            return instance;
        }
};

// init static member
LazySingleton* LazySingleton::instance = NULL;

// 问题1：Lazy Singleton存在内存泄露的问题，有两种解决方法：
// 使用智能指针
// 使用静态的嵌套类对象


// ============================================================================
// version 1.1
// ============================================================================
class Singleton
{
    private:
        static Singleton* instance;

    private:
        Singleton() { };
        ~Singleton() { };
        Singleton(const Singleton&);
        Singleton& operator=(const Singleton&);

    private:
// 在程序运行结束时，系统会调用静态成员deletor的析构函数，该析构函数会删除单例
// 的唯一实例。使用这种方法释放单例对象有以下特征：
// - 在单例类内部定义专有的嵌套类。
// - 在单例类内定义私有的专门用于释放的静态成员。
// - 利用程序在结束时析构全局变量的特性，选择最终的释放时机。
// 在单例类内再定义一个嵌套类，总是感觉很麻烦。
        class Deletor {
            public:
                ~Deletor() {
                    if(Singleton::instance != NULL)
                        delete Singleton::instance;
                }
        };
        static Deletor deletor;

    public:

// 问题2：这个代码在单线程环境下是正确无误的，但是当拿到多线程环境下时这份代码就
// 会出现race condition，注意version 1.0与version 1.1都不是线程安全的。
// 要使其线程安全，能在多线程环境下实现单例模式，我们首先想到的是利用同步机制来
// 正确的保护我们的shared data。
// 可以使用双检测锁模式（DCL: Double-Checked Locking Pattern）：
        static Singleton* getInstance() {
            if(instance == NULL) {
                instance = new Singleton();
            }
            return instance;
        }

        // DCL
        static Singleton* getInstanceThreadSafe() {
            if(instance == NULL) {
                Lock lock;  // 基于作用域的加锁，超出作用域，自动调用析构函数解锁
                if(instance == NULL) {
                    instance = new Singleton();
                }
            }
            return instance;
        }
};


// init static member
Singleton* Singleton::instance = NULL;


// ============================================================================
// version 1.2 (Best of all)
//
// C++11规定了local static在多线程条件下的初始化行为，要求编译器保证了内部静态
// 变量的线程安全性。在C++11标准下，《Effective C++》提出了一种更优雅的单例模式
// 实现，使用函数内的 local static 对象。这样，只有当第一次访问getInstance()
// 方法时才创建实例。这种方法也被称为Meyers' Singleton。
//
// XXX: C++0x之后该实现是线程安全的，C++0x之前仍需加锁。
// ============================================================================
class Singleton
{
    private:
        Singleton() { };
        ~Singleton() { };
        Singleton(const Singleton&);
        Singleton& operator=(const Singleton&);
    public:
        static Singleton& getInstance() 
        {
            static Singleton instance;
            return instance;
        }
};


// ============================================================================
// version 1.3 
// 饿汉版（Eager Singleton）：指单例实例在程序运行时被立即执行初始化
// ============================================================================
class EagerSingleton
{
    private:
        static EagerSingleton instance;
    private:
        EagerSingleton();
        ~EagerSingleton();
        EagerSingleton(const EagerSingleton&);
        EagerSingleton& operator=(const EagerSingleton&);
    public:
        static EagerSingleton& getInstance() {
            return instance;
        }
}

// initialize defaultly
EagerSingleton EagerSingleton::instance;

// 由于在main函数之前初始化，所以没有线程安全的问题。但是潜在问题在于
// no-local static对象（函数外的static对象）在不同编译单元中的初始化顺序是
// 未定义的。也即，static Singleton instance;和static Singleton& getInstance()
// 二者的初始化顺序不确定，如果在初始化完成之前调用 getInstance()
// 方法会返回一个未定义的实例。


// 补充：C++中static对象的初始化

// non-local static对象（函数外）
// C++规定，non-local static 对象的初始化发生在main函数执行之前，也即main函数之
// 前的单线程启动阶段，所以不存在线程安全问题。但C++没有规定多个non-local static
// 对象的初始化顺序，尤其是来自多个编译单元的non-local static对象
// 他们的初始化顺序是随机的。

// local static 对象（函数内）
// 对于local static 对象，其初始化发生在控制流第一次执行到该对象的初始化语句时.
// 多个线程的控制流可能同时到达其初始化语句。

// 在C++11之前，在多线程环境下local static对象的初始化并不是线程安全的。
// 具体表现就是：如果一个线程正在执行local static对象的初始化语句但
// 还没有完成初始化，此时若其它线程也执行到该语句，那么这个线程会认为自己是
// 第一次执行该语句并进入该local static对象的构造函数中。
// 这会造成这个local static对象的重复构造，进而产生内存泄露问题。所以
// local static对象在多线程环境下的重复构造问题是需要解决的。

// 而C++11则在语言规范中解决了这个问题。C++11规定，在一个线程开始local static
// 对象的初始化后到完成初始化前，其他线程执行到这个local static对象的初始化语句
// 就会等待，直到该local static 对象初始化完成。




