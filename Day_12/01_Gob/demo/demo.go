package demo

import (
	"Go_Learn/Day_12/01_Gob/model"
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
)

// JSONDemo json序列化和反序列化
func JSONDemo() {
	var s1 = model.S{
		Data: make(map[string]interface{}, 8),
	}

	s1.Data["count"] = 1 // int
	// 序列化
	ret, err := json.Marshal(s1.Data)
	if err != nil {
		fmt.Println("marshal failed err:", err)
	}
	fmt.Println(ret)
	fmt.Println(string(ret))

	var s2 = model.S{
		Data: make(map[string]interface{}, 8),
	}
	// 反序列化
	err = json.Unmarshal(ret, &s2.Data)
	if err != nil {
		fmt.Println("unmarshal failed err:", err)
	}

	// 循环打印数据信息和类型
	for _, v := range s2.Data {
		fmt.Printf("value: %v, type: %T\n", v, v)
	}
}

// GobDemo gob序列化和反序列化
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
