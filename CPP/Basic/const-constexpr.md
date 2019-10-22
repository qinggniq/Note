# `const` 和 `constexpr`
## 语法
>  **const** literal-type identifier = constant-expression ; **const** literal-type identifier { constant-expression } ; **const** literal-type identifier ( params ) ; **const** ctor ( params ) ; 
> **constexpr** literal-type identifier = constant-expression ; **constexpr** literal-type identifier { constant-expression } ; **constexpr** literal-type identifier ( params ) ; **constexpr** ctor ( params ) ; 
## 语义
- 相同的是`const`和`constexpr`会让编译器在代码试图修改被修饰变量时报错。
- 不同是
  - 语义：`constexpr`指被变量、返回值是常量，并且如果可能的话，尽量在编译期间就确定。
  - 修饰对象：`constexpr`还可以修饰类的构造函数。
  - 确定时期：`const`修饰的变量可以在运行时期确定，而`constexpr`必须在编译期间确定。

