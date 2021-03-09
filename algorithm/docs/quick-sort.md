# 快速排序

## 实现
* [quick-sort.go](../quick-sort.go)
* [quick-sort_test.go](../quick-sort_test.go)
  
  package sort 中使用的是快排

## 原理
>分治法：将问题分解为若干个规模更小但结构与原问题相似的子问题。递归地解这些子问题，然后将这些子问题的解组合为原问题的解。

分治法，选择基准值，
从左右两边交替遍历数组（或者单边），
左侧找到大于基准的和右边小于基准的进行交换，
依次迭代可适用递归。

## 复杂度
* 平均时间复杂度为O(nlogn)，最坏情况为O(n^2)
* [空间复杂度 O(logn)](https://zh.wikipedia.org/wiki/%E5%BF%AB%E9%80%9F%E6%8E%92%E5%BA%8F#%E7%A9%BA%E9%96%93%E8%A4%87%E9%9B%9C%E5%BA%A6)
* 不稳定排序

## 参考
* https://colobu.com/2018/06/26/implement-quick-sort-in-golang/
* http://data.biancheng.net/view/117.html
* https://juejin.cn/post/6844903910772047886
* https://visualgo.net/zh/sorting