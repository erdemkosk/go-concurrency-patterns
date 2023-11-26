
# Go Concurrency Patterns 

Concurrency Patterns about go are collected in this repo. For **learning purposes**


## Appendix

Concurrency refers to the ability of a software system to handle **multiple tasks** or **processes at the same time**. It allows programs to execute several computations simultaneously, making efficient use of available resources.

In software, concurrency doesn’t necessarily mean true simultaneous execution on multiple cores. Instead, it involves managing multiple tasks by **interleaving their execution**, giving the **illusion of parallelism**. This is typically achieved through mechanisms like threading, asynchronous programming, or the use of concurrent constructs like **goroutines** in Go or **async/await** in languages like JavaScript.

Concurrency **enables applications to remain responsive by managing various operations concurrently**, such as handling user input, processing data, or performing I/O operations, without waiting for one task to finish before starting another. It's essential in designing systems that can efficiently utilize available resources and improve overall performance.


## Lessons Learned

#### Goroutine 
A goroutine is a lightweight thread of execution in Go. It's a concurrent unit of work that can be executed independently within a Go program. **Goroutines are managed by the Go runtime** and allow multiple functions to run concurrently. They're different from threads in other programming languages in that they're managed at a higher level by the Go runtime, enabling efficient scheduling and utilization of resources. Goroutines are relatively cheap to create and are fundamental to building **concurrent and scalable applications** in Go.


#### Select
The select statement is used to listen for data coming from multiple channels concurrently, but the crucial aspect is the concurrent (or asynchronous) nature. Especially when multiple tasks are being performed, using the select statement allows simultaneous access to data from channels and optimizes the processing time of the program.

For instance, when waiting for data from the network or different threads of execution, it enables the main thread to continue doing other tasks while waiting for data from channels. **Instead of waiting idly until data arrives** from any channel, **select concurrently monitors multiple channels** and acts accordingly when data arrives.

**The main thread waits until data is received from** any of the channels being listened to with the select statement. However, it can also continue with other tasks, execute other threads of execution, or perform different operations.

This feature allows programs to monitor and manage various tasks concurrently, preventing the program from blocking in unexpected situations. The use of select is highly beneficial in scenarios involving asynchronous operations or listening for data from various sources.

#### Buffered vs Unbuffered Channels
Unbuffered Channels: These channels facilitate synchronous communication between the sender and receiver. When the sender sends data, the receiver waits until it receives the data. This ensures **synchronization between the sender and receiver**, allowing for instantaneous data transfer. For instance, ch := make(chan int) creates an unbuffered channel.

Buffered Channels: These channels have a specific capacity and allow data to be sent up to that capacity. The sender can send data as long as the channel is not full, and the receiver doesn't need to wait for data until it's ready to receive. This provides more flexibility in data transfer and is suitable for situations that don't require strict synchronization. For example, ch := make(chan int, 5) creates a buffered channel that **can hold up to 5 data items**.

#### The Done Channel
Another example of using channels is done. In this approach, a very long goroutine **can be stopped**. Similar to go context logic, a long goroutine can be used as a **breaker at any time**.

#### Pipeline




