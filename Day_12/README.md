#### <center>Day_12</center>

1. gob序列化示例

#### <center>笔记</center>

1. > gob序列化示例

    - 标准库gob是golang提供的 "私有" 的编码方式，它的效率会比json、xml等更高，特别适合在Go语言程序间传递数据。

        ```go
        func GobDemo() {
            var s1 = model.S{
                Data: make(map[string]interface{}, 8),
            }
            s1.Data["count"] = 1

            // encode 编码
            buf := new(bytes.Buffer)   // 创建一个字节类型的缓冲区
            enc := gob.NewEncoder(buf) // 创建一个新的编码器对象
            err := enc.Encode(s1.Data) // 开始编码 编码后的数据会放在"创建编码器对象"时传入的缓冲区
            if err != nil {
                fmt.Println("gob encode failed, err:", err)
                return
            }

            b := buf.Bytes() // 拿到编码后的字节数据
            fmt.Println(b)

            var s2 = model.S{
                Data: make(map[string]interface{}, 8),
            }
            // decode 解码
            dec := gob.NewDecoder(bytes.NewBuffer(b)) // 创建一个新的解码器对象 并传入一个字节类型的缓冲区(里边保存需要解码的字节数据)
            err = dec.Decode(&s2.Data)                // 开始解码
            if err != nil {
                fmt.Println("gob decode failed, err:", err)
                return
            }
            fmt.Println(s2.Data)
            for _, v := range s2.Data {
                fmt.Printf("value: %v, type: %T\n", v, v)
            }
        }
        ```

2. > sd