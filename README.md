# 支持优先级协程调度，基于go1.16.7改造

### 编译
`git clone https://git.code.tencent.com/jemuelmiao/go.git`

`cd go/src`

`./all.bash`

编译后的go在目录go/bin

### 关键字说明
增加了gorder关键字，使用gorder调用的函数第一个参数需要为整型优先级，值越大优先级越高。

### 使用方式

```golang
package main
import (
        "fmt"
        "sync"
)
func main() {
        var wg sync.WaitGroup
        //=====任务A=====
        wg.Add(1)
        gorder func(priority int) {
                defer wg.Done()
                fmt.Println("任务A")
        }(1)
        //=====任务B=====
        wg.Add(1)
        gorder func(priority int) {
                defer wg.Done()
                fmt.Println("任务B")
        }(100)
        //=====任务C=====
        wg.Add(1)
        gorder func(priority int) {
                defer wg.Done()
                fmt.Println("任务C")
        }(10)
        wg.Wait()
}
```
