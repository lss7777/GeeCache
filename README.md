# GeeCache

GeeCache是一个模仿groupcache实现的分布式缓存系统


花费7天时间分别解决内存如何淘汰(LRU置换)，内存并发读写如何安全，如何使用一致性哈希算法寻找访问节点(防止缓存雪崩)，分布式节点的访问，防止缓存击穿，

使用protobuf通信提高二进制传输效率等问题。