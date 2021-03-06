# 数组

## 定义
数组（Array）是一种线性表数据结构。它用一组连续的内存空间，来存储一组具有相同类型的数据。
### 线性表
线性表就是数据排成像一条线一样的结构。每个线性表上的数据最多只有前和后两个方向。除了数组，链表、队列、栈等也是线性表结构。
### 连续的内存空间和相同类型的数据
支持**随机访问**，不过如果想在数组中删除、插入一个数据，为了保证连续性，就需要做大量的数据搬移工作。**平均情况时间复杂度为 O(n)**

a[i]_address = base_address + i * data_type_size

>q:为什么大多数编程语言中，数组要从 0 开始编号，而不是从 1 开始呢？
> 
>a:从数组存储的内存模型上来看，“下标”最确切的定义应该是“偏移（offset）”。
> 从 1 开始编号，每次随机访问数组元素都多了一次减法运算，对于 CPU 来说，就是多了一次减法指令。
> 所以为了减少一次减法操作，数组选择了从 0 开始编号，而不是从 1 开始。

访问越界
```
int main(int argc, char* argv[]){
    int i = 0;
    int arr[3] = {0};
    for(; i<=3; i++){
        arr[i] = 0;
        printf("hello world\n");
    }
    return 0;
}
```
无限打印“hello world”,
在 C 语言中，只要不是访问受限的内存，所有的内存空间都是可以自由访问的。
根据我们前面讲的数组寻址公式，a[3]也会被定位到某块不属于数组的内存地址上，
而这个地址正好是存储变量 i 的内存地址，那么 a[3]=0 就相当于 i=0，
所以就会导致代码无限循环。

函数体内的局部变量存在栈上，且是连续压栈。在Linux进程的内存布局中，
栈区在高地址空间，从高向低增长。变量i和arr在相邻地址，且i比arr的地址大，
所以arr越界正好访问到i。当然，前提是i和arr元素同类型，否则那段代码仍是未决行为。

## 参考
* [**数据结构与算法之美**](http://gk.link/a/10p9l)