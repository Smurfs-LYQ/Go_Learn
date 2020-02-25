#### <center>Day_13</center>

1. etcd put和get示例
2. etcd watch示例

#### <center>笔记</center>

1. > etcd 配置通过IP访问

    ```shell
    --listen-client-urls http://0.0.0.0:2379 --advertise-client-urls http://0.0.0.0:2379 --listen-peer-urls http://0.0.0.0:2380
    ```

    其中:

    - `--listen-client-urls` : 监听本机的URL，监听本机的哪个网卡，哪个端口

    - `--advertise-client-urls` : 告知客户端URL，也就是可以跟etcd服务交互的URL

    - `--listen-peer-urls` : 监听URL，用于与其他节点通讯
