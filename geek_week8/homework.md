###### 1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。
*   10 
```shell
root@286f5dffec27:/data# redis-benchmark -t get,set -d 10  
====== SET ======                                                     
  100000 requests completed in 0.68 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
Summary:
  throughput summary: 146842.88 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.181     0.080     0.167     0.287     0.495     1.831

====== GET ======                                                     
  100000 requests completed in 0.65 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Summary:
  throughput summary: 153846.16 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.172     0.040     0.167     0.247     0.327     0.487
```
*   20
```
root@286f5dffec27:/data# redis-benchmark -t get,set -d 20
====== SET ======                                                     
  100000 requests completed in 0.69 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
Summary:
  throughput summary: 145560.41 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.182     0.040     0.167     0.295     0.519     2.159

====== GET ======                                                     
  100000 requests completed in 0.64 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Summary:
  throughput summary: 156006.25 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.170     0.048     0.167     0.247     0.327     0.783
```
*   50
```shell
====== SET ======                                                     
  100000 requests completed in 0.69 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
Summary:
  throughput summary: 145772.59 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.182     0.048     0.167     0.295     0.487     2.615

====== GET ======                                                     
  100000 requests completed in 0.65 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
Summary:
  throughput summary: 153609.83 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.173     0.080     0.167     0.247     0.335     0.839
```
*   100
```
====== SET ======                                                     
  100000 requests completed in 0.68 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
Summary:
  throughput summary: 147929.00 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.180     0.032     0.167     0.287     0.479     2.015
====== GET ======                                                     
  100000 requests completed in 0.66 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
Summary:
  throughput summary: 151975.69 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.175     0.064     0.167     0.247     0.319     0.543
```
*   200
```
====== SET ======                                                     
  100000 requests completed in 0.68 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
Summary:
  throughput summary: 148148.14 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.179     0.056     0.167     0.287     0.479     1.479
====== GET ======                                                     
  100000 requests completed in 0.65 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
Summary:
  throughput summary: 154798.75 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.171     0.072     0.167     0.247     0.343     0.655
```
*   1k
```
root@286f5dffec27:/data# redis-benchmark -t get,set -d 1000
====== SET ======                                                     
  100000 requests completed in 0.69 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
Summary:
  throughput summary: 145137.88 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.183     0.080     0.167     0.287     0.527     2.031
====== GET ======                                                     
  100000 requests completed in 0.65 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
Summary:
  throughput summary: 153846.16 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.173     0.064     0.167     0.255     0.327     0.863
```
*   5k
```
root@286f5dffec27:/data# redis-benchmark -t get,set -d 5000
====== SET ======                                                     
  100000 requests completed in 0.71 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
Summary:
  throughput summary: 141043.72 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.187     0.040     0.175     0.295     0.527     3.127
====== GET ======                                                     
  100000 requests completed in 0.69 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no
Summary:
  throughput summary: 143884.89 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.185     0.080     0.183     0.255     0.343     0.591
```

*   测试了几次发现每次的benchmark结果都不一样....,但是整体来说这个大小的value存取效率差距不大。


###### 2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。
redis-benchmark -t set -d 10000 -n 100000
* 1W
```
插入数据前
used_memory:22943392
used_memory_human:21.88M
插入数据后
used_memory:26039768
used_memory_human:24.83M
约(26039768-22943392)/100000=30
```
* 10W
  插入数据前
root@286f5dffec27:/data# redis-cli info memory
# Memory
used_memory:23047840
used_memory_human:21.98M
插入数据后
# Memory
used_memory:29832800
used_memory_human:28.45M
约(29832800-23047840)/100000=67
* 50W
```
插入前:
used_memory:23457440
used_memory_human:22.37M
插入数据后
used_memory:50723032
used_memory_human:48.37M
约(50723032-23457440)/100000=272

```
