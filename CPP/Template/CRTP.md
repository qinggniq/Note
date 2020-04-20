# CRTP （奇异递归模版模式）

一个设计模式。

> 继承一个由自己作为模版参数的类。

## 例子

### 静态多态

```c++
template<typename T>
class base {
  public:
  	void interface() {
      std::static_cast<T>(this)->implemenation();
    }
};

class drive : public base<drive> {
 public:
 	void implemenation() {
    //...
  } 
}
```

这样父类类型的对象就会指向子类的实现了，并且不需要虚表。

**更多的细节在[这里](https://www.cnblogs.com/kesalin/archive/2010/03/25/CRTP.html)**

### 对象计数

```c++
template <typename T>
struct counter
{
    static int objects_created;
    static int objects_alive;

    counter()
    {
        ++objects_created;
        ++objects_alive;
    }
    
    counter(const counter&)
    {
        ++objects_created;
        ++objects_alive;
    }
protected:
    ~counter() // objects should never be removed through pointers of this type
    {
        --objects_alive;
    }
};
template <typename T> int counter<T>::objects_created( 0 );
template <typename T> int counter<T>::objects_alive( 0 );

class X : counter<X>
{
    // ...
};

class Y : counter<Y>
{
    // ...
};
```

和普通的用一个基类计数的区别在于可以区分不同类型的对象。

