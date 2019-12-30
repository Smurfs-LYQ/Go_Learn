#### <center>NSQ</center>

1. > NSQ介绍

    NSQ是Go语言编写的一个开源的实时分布式内存消息队列，性能十分优异。NSQ有以下优势:

    - NSQ提倡分布式和分散的拓扑，没有单点故障，支持容错和高可用性，并提供可靠的消息交付保证。
    - NSQ支持横向扩展，没有任何集中式代理。
    - NSQ易于配置和部署，并且内置了管理界面。

2. > NSQ的应用场景

    通常来说，消协队列都使用一下场景:

    - 异步队列

       参照下图利用消息队列把业务流程中的非关键流程异步化，从而显著降级业务请求的响应时间。

       ![avatar](https://www.liwenzhou.com/images/Go/nsq/nsq1.png)

    - 应用解耦

       通过使用消息队列将不同的业务逻辑解耦，降低系统间的耦合，提高系统的健壮性。后续有其他业务要使用订单数据可直接订阅消息队列，提高系统的灵活性。

       ![avatar](https://www.liwenzhou.com/images/Go/nsq/nsq2.png)

    - 流量削峰

       类似秒杀(大秒)等场景下，某一时间可能会产生大量的请求，使用消息队列能够为后续处理请求提供一定的缓冲区，保证后端服务的稳定性。

       ![avatar](https://www.liwenzhou.com/images/Go/nsq/nsq3.png)

3. > NSQ组件

    - **nsqd**

        nsqd是一个守护进程，它接收、排队并向客户端发送消息。
        启动 `nsqd`，指定 `-broadcast-address=IP` 来配置广播地址

        ```shell
        ./nsqd -broadcast-address=127.0.0.1
        ```

        如果再搭配 `nsqlookupd` 使用的模式下还需要指定 `nsqlookupd` 地址:

        ```shell
        ./nsqd -broadcast-address=127.0.0.1 -lookupd-tcp-address=127.0.0.1:4160
        ```

        如果是部署了多个 `nsqlookup` 节点的集群，那还可以指定多个 `-lookupd-tcp-address`。

        `nsqdq` 相关配置项如下:

        ```conf
        -auth-http-address value
            <addr>:<port> to query auth server (may be given multiple times)
        -broadcast-address string
            address that will be registered with lookupd (defaults to the OS hostname) (default "PROSNAKES.local")
        -config string
            path to config file
        -data-path string
            path to store disk-backed messages
        -deflate
            enable deflate feature negotiation (client compression) (default true)
        -e2e-processing-latency-percentile value
            message processing time percentiles (as float (0, 1.0]) to track (can be specified multiple times or comma separated '1.0,0.99,0.95', default none)
        -e2e-processing-latency-window-time duration
            calculate end to end latency quantiles for this duration of time (ie: 60s would only show quantile calculations from the past 60 seconds) (default 10m0s)
        -http-address string
            <addr>:<port> to listen on for HTTP clients (default "0.0.0.0:4151")
        -http-client-connect-timeout duration
            timeout for HTTP connect (default 2s)
        -http-client-request-timeout duration
            timeout for HTTP request (default 5s)
        -https-address string
            <addr>:<port> to listen on for HTTPS clients (default "0.0.0.0:4152")
        -log-prefix string
            log message prefix (default "[nsqd] ")
        -lookupd-tcp-address value
            lookupd TCP address (may be given multiple times)
        -max-body-size int
            maximum size of a single command body (default 5242880)
        -max-bytes-per-file int
            number of bytes per diskqueue file before rolling (default 104857600)
        -max-deflate-level int
            max deflate compression level a client can negotiate (> values == > nsqd CPU usage) (default 6)
        -max-heartbeat-interval duration
            maximum client configurable duration of time between client heartbeats (default 1m0s)
        -max-msg-size int
            maximum size of a single message in bytes (default 1048576)
        -max-msg-timeout duration
            maximum duration before a message will timeout (default 15m0s)
        -max-output-buffer-size int
            maximum client configurable size (in bytes) for a client output buffer (default 65536)
        -max-output-buffer-timeout duration
            maximum client configurable duration of time between flushing to a client (default 1s)
        -max-rdy-count int
            maximum RDY count for a client (default 2500)
        -max-req-timeout duration
            maximum requeuing timeout for a message (default 1h0m0s)
        -mem-queue-size int
            number of messages to keep in memory (per topic/channel) (default 10000)
        -msg-timeout string
            duration to wait before auto-requeing a message (default "1m0s")
        -node-id int
            unique part for message IDs, (int) in range [0,1024) (default is hash of hostname) (default 616)
        -snappy
            enable snappy feature negotiation (client compression) (default true)
        -statsd-address string
            UDP <addr>:<port> of a statsd daemon for pushing stats
        -statsd-interval string
            duration between pushing to statsd (default "1m0s")
        -statsd-mem-stats
            toggle sending memory and GC stats to statsd (default true)
        -statsd-prefix string
            prefix used for keys sent to statsd (%s for host replacement) (default "nsq.%s")
        -sync-every int
            number of messages per diskqueue fsync (default 2500)
        -sync-timeout duration
            duration of time per diskqueue fsync (default 2s)
        -tcp-address string
            <addr>:<port> to listen on for TCP clients (default "0.0.0.0:4150")
        -tls-cert string
            path to certificate file
        -tls-client-auth-policy string
            client certificate auth policy ('require' or 'require-verify')
        -tls-key string
            path to key file
        -tls-min-version value
            minimum SSL/TLS version acceptable ('ssl3.0', 'tls1.0', 'tls1.1', or 'tls1.2') (default 769)
        -tls-required
            require TLS for client connections (true, false, tcp-https)
        -tls-root-ca-file string
            path to certificate authority file
        -verbose
            enable verbose logging
        -version
            print version string
        -worker-id
            do NOT use this, use --node-id
        ```

    - **nsqlookupd**

        nsqlookupd是维护所有nsqd状态、提供服务发现的守护进程。它能为消费者查找特定 `topic` 下的nsqd，提供了运行时的自动发现服务。它不维持持久状态，也不需要与任何其他nsqlookupd实例协调以满足查询。因此根据你系统的冗余要求尽可能多的部署 `nsqlookupd` 节点。它们消耗的资源很少，可以与其他服务共存。建议为每个数据中心运行只要3个集群。

        ```conf
        -broadcast-address string
            address of this lookupd node, (default to the OS hostname) (default "PROSNAKES.local")
        -config string
            path to config file
        -http-address string
            <addr>:<port> to listen on for HTTP clients (default "0.0.0.0:4161")
        -inactive-producer-timeout duration
            duration of time a producer will remain in the active list since its last ping (default 5m0s)
        -log-prefix string
            log message prefix (default "[nsqlookupd] ")
        -tcp-address string
            <addr>:<port> to listen on for TCP clients (default "0.0.0.0:4160")
        -tombstone-lifetime duration
            duration of time a producer will remain tombstoned if registration remains (default 45s)
        -verbose
            enable verbose logging
        -version
            print version string
        ```

    - **nsqadmin**

        一个实时监控集群状态、执行各种管理任务的Web管理平台。启动 `nsqadmin`，指定 `nsqlookupd` 地址:

        ```shell
        ./nsqadmin -lookupd-http-address=127.0.0.1:4167
        ```

        使用浏览器打开 `http://127.0.0.1:4171/` 访问管理界面。

        `nsqadmin` 相关的配置项如下:

        ```conf
        -allow-config-from-cidr string
            A CIDR from which to allow HTTP requests to the /config endpoint (default "127.0.0.1/8")
        -config string
            path to config file
        -graphite-url string
            graphite HTTP address
        -http-address string
            <addr>:<port> to listen on for HTTP clients (default "0.0.0.0:4171")
        -http-client-connect-timeout duration
            timeout for HTTP connect (default 2s)
        -http-client-request-timeout duration
            timeout for HTTP request (default 5s)
        -http-client-tls-cert string
            path to certificate file for the HTTP client
        -http-client-tls-insecure-skip-verify
            configure the HTTP client to skip verification of TLS certificates
        -http-client-tls-key string
            path to key file for the HTTP client
        -http-client-tls-root-ca-file string
            path to CA file for the HTTP client
        -log-prefix string
            log message prefix (default "[nsqadmin] ")
        -lookupd-http-address value
            lookupd HTTP address (may be given multiple times)
        -notification-http-endpoint string
            HTTP endpoint (fully qualified) to which POST notifications of admin actions will be sent
        -nsqd-http-address value
            nsqd HTTP address (may be given multiple times)
        -proxy-graphite
            proxy HTTP requests to graphite
        -statsd-counter-format string
            The counter stats key formatting applied by the implementation of statsd. If no formatting is desired, set this to an empty string. (default "stats.counters.%s.count")
        -statsd-gauge-format string
            The gauge stats key formatting applied by the implementation of statsd. If no formatting is desired, set this to an empty string. (default "stats.gauges.%s")
        -statsd-interval duration
            time interval nsqd is configured to push to statsd (must match nsqd) (default 1m0s)
        -statsd-prefix string
            prefix used for keys sent to statsd (%s for host replacement, must match nsqd) (default "nsq.%s")
        -version
            print version string
        ```

4. > NSQ架构

    - **NSQ工作模式**

        ![avatar](https://www.liwenzhou.com/images/Go/nsq/nsq4.png)

    - **Topic 和 Channel**

        每个nsqd实例只在一次处理多个数据流。这些数据流称为 `topics`，一个 `topic` 具有1个或多个 `channels`。每个 `channel` 都会收到 `topic` 所有消息的副本，实际上下游的服务器是通过对应的 `channel` 来消费 `topic` 消息。

        `topic` 和 `channel` 不是预先配置的。`topic` 在首次使用时创建，方法是将其发布到指定 `topic`，或者订阅指定 `topic` 上的 `channel`。`channel` 是通过订阅指定的 `channel` 在第一次使用时创建的。

        `topic` 和 `channel` 都相互独立的缓冲数据，防止缓慢的消费者导致其它 `channel` 的挤压(同样使用于 `topic` 级别)。

        `channel` 可以并且通常会连接多个客户端。假设所有连接的客户端都处于准备接受消息的状态，则每条消息将被传递到随机客户端。例如:

        ![avatar](https://www.liwenzhou.com/images/Go/nsq/nsq5.gif)
        
        总而言之，消息是从 `topic -> channel` (每个channel接受该topic的所有消息的副本) 多播的，但是从 `channel -> consumers` 均匀分布 (每个消费者都接收channel的一部分消息)。

    - **NSQ接受和发送消息流程**

        ![avatar](https://www.liwenzhou.com/images/Go/nsq/nsq6.png)

    - **NSQ特性**

        - 消息默认不持久化，可以配置成持久化模式。nsq采用的方式是 `内存+硬盘` 模式，当内存到达一定程度时就会将数据持久化到硬盘。
            - 如果将 `--mem-queue-size` 设置为0，所有的消息将会存储到磁盘。
            - 服务器重启时也会将当时在内存中的消息持久化。
        - 每条消息至少传递一次。
        - 消息不保证有序。

5. > Go操作NSQ

    - 官方提供了Go语言版的客户端: [go-nsq](https://github.com/nsqio/go-nsq)，更多客户端支持请查看 [CLIENT LIBRARIES](https://nsq.io/clients/client_libraries.html)。

    - 安装

        ```go
        go get -u github.com/nsqio/go-nsq
        ```

    - 生产者

        生产者示例 [demo](https://github.com/Smurfs-LYQ/Go_Learn/blob/master/Day_09/17_NSQ/producer/main.go)

    - 消费者

        消费者示例 [demo](https://github.com/Smurfs-LYQ/Go_Learn/blob/master/Day_09/17_NSQ/consumer/main.go)