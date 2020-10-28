package marshal

import (
	"fmt"
	"bytes"
	"encoding/gob"
	"BoltProject/block"
)


//gob 序列化
func Serialize(b block.Block) ([]byte, error) {
	//1. 开辟缓冲区
	var buf bytes.Buffer
	//2. 创建编码器
	encoder := gob.NewEncoder(&buf)
	//3. 编码
	err := encoder.Encode(b)
	//4. 序列化后的值
	structInfo := buf.Bytes()
	if err != nil {
		fmt.Println("gob序列化错误")
		return nil, err
	}
	return structInfo, nil
}

/**
 * gob 反序列化(b,)
 * 1. 获取序列化后的值
 * 2. 创建解码器
 * 3. 解码
 */
func DeSerialize(structInfo []byte, b block.Block) (*block.Block, error) {
	//创建解码器
	decoder := gob.NewDecoder(bytes.NewReader(structInfo))
	//解码
	err := decoder.Decode(&b)
	if err != nil {
		fmt.Println("gob反序列化错误")
		return nil, err
	}
	return &b, nil
}
