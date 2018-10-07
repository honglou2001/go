基本使用
select是Go中的一个控制结构，类似于switch语句，用于处理异步IO操作。
select会监听case语句中channel的读写操作，当case中channel读写操作为非阻塞状态（即能读写）时，将会触发相应的动作。
  select中的case语句必须是一个channel操作
  select中的default子句总是可运行的。
如果有多个case都可以运行，select会随机公平地选出一个执行，其他不会执行。
如果没有可运行的case语句，且有default语句，那么就会执行default的动作。
如果没有可运行的case语句，且没有default语句，select将阻塞，直到某个case通信可以运行