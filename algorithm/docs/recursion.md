# 递归
## 三个条件
* 一个问题的解可以分解为几个子问题的解
* 这个问题与分解之后的子问题，除了数据规模不同，求解思路完全一样
* 存在递归终止条件

写出递推公式，找到终止条件

## 注意
递归代码要警惕堆栈溢出 ,可以通过在代码中限制递归调用的最大深度的方式来解决这个问题。

递归代码要警惕重复计算,我们可以通过一个数据结构（比如散列表）来保存已经求解过的。

空间复杂度并不是 O(1)，而是 O(n)，可使用迭代循环。

## 参考
* [**数据结构与算法之美**](http://gk.link/a/10p9l)
