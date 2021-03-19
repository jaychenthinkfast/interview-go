# 队列
## 定义
先进者先出，这就是典型的“队列”。

## 应用
循环队列、阻塞队列、并发队列。

它们在很多偏底层系统、框架、中间件的开发中，起着关键性的作用。
比如高性能队列 Disruptor、Linux 环形缓存，都用到了循环并发队列；
Java concurrent 并发包利用 ArrayBlockingQueue 来实现公平锁等。

队列可以应用在任何有限资源池中，用于排队请求，比如数据库连接池等。实际上，对于大部分资源有限的场景，
当没有空闲资源时，基本上都可以通过“队列”这种数据结构来实现请求排队。

## 循环队列
```

public class CircularQueue {
  // 数组：items，数组大小：n
  private String[] items;
  private int n = 0;
  // head表示队头下标，tail表示队尾下标
  private int head = 0;
  private int tail = 0;

  // 申请一个大小为capacity的数组
  public CircularQueue(int capacity) {
    items = new String[capacity];
    n = capacity;
  }

  // 入队
  public boolean enqueue(String item) {
    // 队列满了
    if ((tail + 1) % n == head) return false;
    items[tail] = item;
    tail = (tail + 1) % n;
    return true;
  }

  // 出队
  public String dequeue() {
    // 如果head == tail 表示队列为空
    if (head == tail) return null;
    String ret = items[head];
    head = (head + 1) % n;
    return ret;
  }
}
```

## 阻塞队列和并发队列
阻塞队列其实就是在队列基础上增加了阻塞操作。

并发队列。最简单直接的实现方式是直接在 enqueue()、dequeue() 方法上加锁，但是锁粒度大并发度会比较低，同一时刻仅允许一个存或者取操作。
实际上，基于数组的循环队列，利用 CAS 原子操作，可以实现非常高效的并发队列。这也是循环队列比链式队列应用更加广泛的原因。

## 参考
* [**数据结构与算法之美**](http://gk.link/a/10p9l)



