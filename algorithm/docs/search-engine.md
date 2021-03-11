# 搜索引擎数据结构算法

## 流程概览
* 搜集 
  
    爬虫爬取网页
* 分析
  
    网页内容抽取、分词，构建临时索引，计算 PageRank 值
* 索引

  通过分析阶段得到的临时索引，构建倒排索引
* 查询

  响应用户的请求，根据倒排索引获取相关网页，计算网页排名，返回查询结果给用户

## 搜集
### 爬取链接
**图广度优先搜索**遍历爬取页面链接（**字符串匹配算法**匹配<link>）存到links.bin中

### 网页判重
使用**布隆过滤器（位图：比较“特殊”的散列表）** 对网页判重，可以定期持久化到磁盘中，存储到bloom_filter.bin文件中，
如果机器宕机重启后可以重新加载到内存中（可能会丢失部分数据，不过对于搜索引擎而言可以容忍)

布隆过滤器非常适合这种不需要 100% 准确的、允许存在小概率误判的大规模判重场景。除了爬虫网页去重这个例子，
还有比如统计一个大型网站的每天的 UV 数，也就是每天有多少用户访问了网站，
我们就可以使用布隆过滤器，对重复访问的用户进行去重。
>利用布隆过滤器，在执行效率方面，是否比散列表更加高效呢？

> 布隆过滤器用多个哈希函数对同一个网页链接进行处理，CPU 只需要将网页链接从内存中读取一次，
> 进行多次哈希计算，理论上讲这组操作是 CPU 密集型的。而在散列表的处理方式中，
> 需要读取散列值相同（散列冲突）的多个网页链接，分别跟待判重的网页链接，进行字符串匹配。
> 这个操作涉及很多内存数据的读取，所以是内存密集型的。我们知道 CPU 计算可能是要比内存访问更快速的，
> 所以，理论上讲，布隆过滤器的判重方式，更加快速。

### 网页存储
doc_raw.bin

每个网页都存储为一个独立的文件，那磁盘中的文件就会非常多，数量可能会有几千万，甚至上亿。
常用的文件系统显然不适合存储如此多的文件。所以，我们可以把多个网页存储在一个文件中。
每个网页之间，通过一定的标识进行分隔，方便后续读取。
>例：网页编号doc1_id\t网页大小doc1_size\t网页\r\n\r\n

文件系统对文件的大小也有一定的限制,超出后新建新的文件存储

网页编号doc1_id 由中心计数器分配，doc_id.bin文件存储

## 分析
### 抽取网页信息
**字符串匹配算法** 删除标签

### 分词并创建临时索引
Trie树匹配分词，单词与网页之间对应关系写入临时索引文件tmp_Index.bin
>例：单词编号term1_id\t网页编号doc_id\r\n

给单词编号的方式，跟给网页编号类似。维护一个计数器。
在这个过程中，我们还需要使用**散列表**，记录已经编过号的单词。
在对网页文本信息分词的过程中，我们拿分割出来的单词，先到散列表中查找，
如果找到，那就直接使用已有的编号；如果没有找到，我们再去计数器中拿号码，
并且将这个新单词以及编号添加到散列表中。
单词编号文件：term_id.bin

## 索引
索引阶段主要负责将分析阶段产生的临时索引，构建成倒排索引。
倒排索引（ Inverted index）中记录了每个单词以及包含它的网页列表。
>例：单词编号tid1\t 包含单词的网页编号列表did1,did2。。。didx\r\n

>如何通过临时索引文件，构建出倒排索引文件呢？

我们先对临时索引文件，按照单词编号的大小进行排序。
因为临时索引很大，所以一般基于内存的排序算法就没法处理这个问题了。
我们可以用之前讲到的 **归并排序** 的处理思想，将其分割成多个小文件，先对每个小文件独立排序，
最后再合并在一起。当然，实际的软件开发中，我们其实可以直接利用 MapReduce 来处理。

除了倒排文件之外，我们还需要一个文件，来记录每个单词编号在倒排索引文件中的偏移位置。
我们把这个文件命名为 term_offset.bin。
这个文件的作用是，帮助我们快速地查找某个单词编号在倒排索引中存储的位置，
进而快速地从倒排索引中读取单词编号对应的网页编号列表。
>例：单词编号tid1\t偏移位置offset1分隔符\r\n

## 查询
* doc_id.bin：记录网页链接和编号之间的对应关系。
* term_id.bin：记录单词和编号之间的对应关系。
* index.bin：倒排索引文件，记录每个单词编号以及对应包含它的网页编号列表。
* term_offset.bin：记录每个单词编号在倒排索引文件中的偏移位置。
  
这四个文件中，除了倒排索引文件（index.bin）比较大之外，其他的都比较小。
为了方便快速查找数据，我们将其他三个文件都加载到内存中，并且组织成**散列表**这种数据结构。

1. 查询时，输入文本，分词处理后得到k个单词，
2. 去term_id.bin 找到编号
3. 去term_offset.bin 找到偏移位置
4. 去index.bin倒排索引拿到网页编号列表（根据一定规则排序后）
5. 去doc_id.bin拿到网页链接

## 涉及到的算法
图、散列表、Trie 树、布隆过滤器、单模式字符串匹配算法、AC 自动机、广度优先遍历、归并排序等。

## 参考
* [**数据结构与算法之美**](http://gk.link/a/10p9l)