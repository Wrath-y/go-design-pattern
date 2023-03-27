```
+------------------+          +------------------+
|  AbstractFactory |          |    ProductA      |
+------------------+          +------------------+
| CreateProductA() |<-------->|  OperationA()    |
| CreateProductB() |          +------------------+
+------------------+          +------------------+
                              |    ProductB      |
                              +------------------+
                              |  OperationB()    |
                              +------------------+

+------------------+          +------------------+
| ConcreteFactory1 |          | ConcreteProductA1 |
+------------------+          +------------------+
| CreateProductA() |<-------->|  OperationA()    |
| CreateProductB() |          +------------------+
+------------------+          +------------------+
                              | ConcreteProductB1 |
                              +------------------+
                              |  OperationB()    |
                              +------------------+

+------------------+          +------------------+
| ConcreteFactory2 |          | ConcreteProductA2 |
+------------------+          +------------------+
| CreateProductA() |<-------->|  OperationA()    |
| CreateProductB() |          +------------------+
+------------------+          +------------------+
                              | ConcreteProductB2 |
                              +------------------+
                              |  OperationB()    |
                              +------------------+
```

是一个抽象工厂模式的UML类图。在这个模式中，有一个名为 `AbstractFactory` 的抽象工厂接口，它定义了两个方法： `CreateProductA()` 和 `CreateProductB()`，用于创建不同种类的产品。

`ConcreteFactory1` 和 `ConcreteFactory2` 是具体工厂类，它们实现了 `AbstractFactory` 接口，分别用于创建不同种类的产品。

`ProductA` 和` ProductB` 是抽象产品类，它们定义了各自的操作方法。ConcreteProductA1、ConcreteProductA2、ConcreteProductB1 和 ConcreteProductB2 是具体产品类，它们实现了 `ProductA` 和 `ProductB` 接口，分别对应不同种类的产品。