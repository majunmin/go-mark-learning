# G P M

G: 一个 goroutine                     -- 携带任务
P: 一个装满G的队列,用于维护一些任务      -- 分配任务
M: 一个调度器，用于讲一个G搬到线程上执行  -- 寻找任务

## go 调度的基本流程

1. 创建一个 G 对象
2. 将 G 保存至 P中
3. P 去唤醒(notify)一个 M,然后继续执行它的执行序(分配下一个 G)
4. M 寻找空闲的 P, 读取该 P 要分配的 G
5. 接下来 M 执行一个调度循环,调用 G → 执行 → 清理线程 → 继续找新的 G 执行


### 各自携带的信息

1. G
  - 需执行函数的指令(指针)
  - 线程上下文的信息(goroutine切换时,用于保存 g 的上下文,例如，变量、相关信息等)
  - 现场保护和现场恢复(用于全局队列执行时的保护)
  - 所属的函数栈
  - 当前执行的 m
  - 被阻塞的时间
2. P
> P/M需要进行绑定，构成一个执行单元。P决定了同时可以并发任务的数量，可通过`GOMAXPROCS`限制同时执行用户级任务的操作系统线程.
> 可以通过`runtime.GOMAXPROCS()`进行指定。

  - 状态(空闲 运行...)
  - 关联的 m
  - 可运行的 goroutine 的队列
  - 下一个 g

3. M
> 所有M是有线程栈的. 如果不对该线程栈提供内存的话,系统会给该线程栈提供内存(不同操作系统提供的线程栈大小不同)

  - 所属的调度栈
  - 当前运行的 g
  - 关联的 p
  - 状态

