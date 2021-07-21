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

2. > etcd 基本操作

    **增**

    ```shell
    etcd --endpoints=http://服务地址:端口 put "值"
    etcd put 键 "值"
    ```

    **删**

    ```shell
    etcd --endpoints=http://服务地址:端口 del 键
    etcd del 键
    ```

    **改**

    ```txt
    `改` 的操作和 `增` 是一样的
    ```

    **查**

    ```shell
    etcd --endpoints=http://服务地址:端口 get 键
    etcd get 键

    # 参数
        --prefix : 这个参数代表查询所有以前面的键开始的信息
    ```

    **注意:**

    - 如果就是操作本机的ETCD服务，并且端口是默认的，那么 --endpoints 参数可以不加
