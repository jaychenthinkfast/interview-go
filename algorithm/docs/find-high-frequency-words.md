# 如何从大量数据中找出高频词

## 描述
有一个1G大小的文件，文件里面每一行是一个词，每个词的大小不超过16个字节，内存大小限制是1M，要求返回频数最高的100个词。

## 解答
### 分治法
步骤
1. 将1G文件拆分，确保拆分后的文件不大于内存限制1M
2. 统计单个文件中的词频，可用map存储计数
3. 利用步骤2建立[**topk最大堆**](https://leetcode-cn.com/problems/kth-largest-element-in-an-array/solution/shu-zu-zhong-de-di-kge-zui-da-yuan-su-by-leetcode-/) 通过不断遍历文件的词频更新topk最大堆

## 参考
* Go程序员面试算法宝典