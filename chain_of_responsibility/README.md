### 责任链模式

在这个示例中，`Handler` 是一个接口，定义了一个名为 `SetNextHandler` 的方法和一个名为 `HandleRequest` 的方法，用于设置下一个处理者和处理请求。

`ConcreteHandlerA`、`ConcreteHandlerB` 和 `ConcreteHandlerC` 是具体的处理者类，实现了 `Handler` 接口。

在 `Run` 函数中，创建了三个处理者对象，并设置它们的下一个处理者。然后，循环处理一些请求，每个请求都会被第一个处理者处理，如果第一个处理者无法处理该请求，则会传递给下一个处理者，直到有一个处理者能够处理该请求。