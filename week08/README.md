#Go进阶训练营第8周作业


##1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
###命令:
```
redis-benchmark -d 10 -t get,set
```
###SET
|-|执行次数和耗时|每秒请求次数|
|----|----|----|
|10|100000 requests completed in 3.79 seconds|26420.08 requests per second|
|20|100000 requests completed in 3.78 seconds|26455.03 requests per second|
|50|100000 requests completed in 3.91 seconds|25549.31 requests per second|
|100|100000 requests completed in 3.98 seconds|25138.26 requests per second|
|200|100000 requests completed in 3.80 seconds|26288.12 requests per second|
|1k|100000 requests completed in 3.99 seconds|25056.38 requests per second|
|5k|100000 requests completed in 4.94 seconds|20230.63 requests per second|
###GET
|-|执行次数和耗时|每秒请求次数|
|----|----|----|
|10|100000 requests completed in 3.74 seconds|26737.97 requests per second|
|20|100000 requests completed in 3.98 seconds|25125.63 requests per second|
|50|100000 requests completed in 3.87 seconds|25866.53 requests per second|
|100|100000 requests completed in 3.89 seconds|25706.94 requests per second|
|200|100000 requests completed in 3.88 seconds|25799.79 requests per second|
|1k|100000 requests completed in 4.00 seconds|25025.03 requests per second|
|5k|100000 requests completed in 4.79 seconds|20859.41 requests per second|

### 100000 requests completed in 3.79 seconds
### 26420.08 requests per second

##2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。
###代码链接[github](https://github.com/87170360/redis-benchmark)
###代码工作原理，写入不同数量不同长度的value, 分析内存占用, 导出结果到csv文件
###结论 相同长度的value在写入数量越多情况下，平均每个value占用内存更多
