# Redis常用数据类型数据结构

## 数据类型
字符串、列表、字典、集合、有序集合

## 列表
两种实现方法
1. 压缩列表（ziplist）
    * 列表中保存的单个数据（有可能是字符串类型的）小于 64 字节；
    * 列表中数据个数少于 512 个。 
    
   相较数组压缩列表这种存储结构，**一方面比较节省内存，另一方面可以支持不同类型数据的存储**。
   而且，因为数据存储在一片连续的内存空间，通过键来获取值为列表类型的数据，读取的效率也非常高。

   压缩列表不支持随机访问，但是比较省存储空间啊。Redis一般都是通过key获取整个value的值，也就是整个压缩列表的数据，并不需要随机访问。
   有利于CPU缓存

2. 双向循环链表

## 字典（hash）
1. 压缩列表
   * 字典中保存的键和值的大小都要小于 64 字节；
   * 字典中键值对的个数要小于 512 个。
2. 散列表
   
   Redis 就使用散列表来实现字典类型。Redis 使用MurmurHash2这种运行速度快、随机性好的哈希算法作为哈希函数。
   对于哈希冲突问题，Redis 使用链表法来解决。除此之外，Redis 还支持散列表的动态扩容、缩容。

   当数据动态增加之后，散列表的装载因子会不停地变大。为了避免散列表性能的下降，当装载因子大于 1 的时候，
   Redis 会触发扩容，将散列表扩大为原来大小的 2 倍左右

   当数据动态减少之后，为了节省内存，当装载因子小于 0.1 的时候，Redis 就会触发缩容，缩小为字典中数据个数的大约 2 倍大小

   扩容缩容要做大量的数据搬移和哈希值的重新计算，所以比较耗时。
   Redis 使用渐进式扩容缩容策略，将数据的搬移分批进行，避免了大量数据一次性搬移导致的服务停顿。
   >Go Map扩缩容也是渐进式

## 集合（set）
1. 有序数组
   * 存储的数据都是整数；
   * 存储的数据元素个数不超过 512 个。
2. 散列表

## 有序集合（sortedset）
1. 压缩列表
   * 所有数据的大小都要小于 64 字节；
   * 元素个数要小于 128 个。
2. 跳表 
   
   跳表是一种动态数据结构，支持快速地插入、删除、查找操作，时间复杂度都是 O(logn)。跳表的空间复杂度是 O(n)

   跳表是通过随机函数来维护“平衡性”。（红黑树、AVL 树这样平衡二叉树，是通过左右旋的方式保持左右子树的大小平衡）
   
   >**为什么 Redis 要用跳表来实现有序集合，而不是红黑树？**
   
   Redis 中的有序集合支持的核心操作主要有下面这几个：
   * 插入一个数据；
   * 删除一个数据；
   * 查找一个数据；
   * 按照区间查找数据（比如查找值在[100, 356]之间的数据）；
   * 迭代输出有序序列。
   
   1. 其中，插入、删除、查找以及迭代输出有序序列这几个操作，红黑树也可以完成，时间复杂度跟跳表是一样的。
   但是，按照区间来查找数据这个操作，红黑树的效率没有跳表高。
   对于按照区间查找数据这个操作，跳表可以做到 O(logn) 的时间复杂度定位区间的起点，然后在原始链表中顺序往后遍历就可以了。这样做非常高效。
   2. 跳表更容易代码实现。 虽然跳表的实现也不简单，但比起红黑树来说还是好懂、好写多了，而简单就意味着可读性好，不容易出错。
   3. 跳表更加灵活，它可以通过改变索引构建策略，有效平衡执行效率和内存消耗。
   
   跳表也不能完全替代红黑树。因为红黑树比跳表的出现要早一些，很多编程语言中的 Map 类型都是通过红黑树来实现的。但是跳表并没有一个现成的实现，要自己实现。

## 参考
* [**数据结构与算法之美**](http://gk.link/a/10p9l)