# Go 程序逻辑学习

## 程序架构流程图

## 代码清单——程序的项目结构

### 知识点补充

1. 要让 Go 语言对包做初始化操作`init`，但是并不使用包里的标识符，可以在`import`的包前加`_`。

    例如：

    ```go
    import(
        _ "github.com/luckyOrangeGo/studtLog/Log/20180419-程序/sample/matchers"
    )
    ```

2. 在 Go 语言中，所有变量都被初始化为其零值。

    对于数值类型，零值是 `0` ；对于字符串类型，零值是空字符串 `""` ；对于布尔类型，零值是 `false` ；对于指针，零值是 `nil` 。

    对于引用类型来说，所引用的底层数据结构会被初始化为对应的零值。但是被声明为其零值的引用类型的变量，会返回nil作为其值。

3. `var matchers = make(map[string]Matcher)`

    包级变量 `matchers` ,使用 `var` 声明为 `Matcher` 类型的映射（ `map` ），这个映射以 `string` 类型值作为键， `Matcher` 类型值作为映射后的值。

    其他包可以间接访问不公开的标识符。 例如，一个函数可以返回一个未公开类型的值，那么这个函数的任何调用者，哪怕调用者不是在这个包里声明的，都可以访问这个值。

4. 根据经验，如果需要声明初始值为零值的变量，应该使用 `var` 关键字声明变量； 如果提供确切的非零值初始化变量或者使用函数返回值创建变量，应该使用简化变量声明运算符 `:=` 。

5. `results := make(chan *Result)`

    使用内置的 `make` 函数创建了一个无缓冲的通道。

    在 Go 语言中，通道（channel）和映射（map）与切片（slice）一样，也是引用类型，不过，通道本身实现的是一组带类型的值，这组值用于在 `goroutine` 之间传递数据。通道内置同步机制， 从而保证通信安全。

6. 防止程序在全部搜索执行完之前终止
    ```go
    // 构造一个wait group，以便处理所有的数据源
    var waitGroup sync.WaitGroup

    // 设置需要等待处理
    // 每个数据源的goroutine的数量
    waitGroup.Add(len(feeds)
    ```

    使用 `sync` 包的 `WaitGroup` 跟踪所有启动的 `goroutine` 。非常推荐使用 `WaitGroup` 来 跟踪 `goroutine` 的工作是否完成。

    `WaitGroup` 是一个计数信号量，可以利用它来统计所有的 `goroutine` 是不是都完成了工作。

7. 为每个数据源启动 goroutine 的代码

    ```go
    // 为每个数据源启动一个goroutine来查找结果
    for _, feed := range feeds {
        // 获取一个匹配器用于查找
        matcher, exists := matchers[feed.Type]
        if !exists {
            matcher = matchers["default"]
        }

        // 启动一个个匿名函数作为 goroutine来执行搜索
        go func(matcher Matcher, feed *Feed) {
            Match(matcher, feed, searchTerm, results)
            waitGroup.Done()
        }(matcher, feed)
    }
    ```

    匿名函数是指没有明确声明名字的函数。在 for range 循环里，为每个数据源，以 goroutine 的方式启动了一个匿名函数。这样可以并发地独立处理每个数据源的数据。

    matcher和feed两个变量的值被传入匿名函数。
