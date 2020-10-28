//bolt 操作

package blockChain

import (
	"fmt"
	"github.com/bolt"
)

/**
 * blockChain 是区块链，链式结构的管理者（包括对区块的增，删，改，查）
 * 增：将区块添加到链上去
 * 删，改：区块链不可篡改，不支持
 * 查：[
 *  1.由hash值去查询某个特定的区块
 *  2.由区块的高度查询某个特定的区块
 *  3.获取区块链上所有的数据（遍历区块链）
 * ]
 */
type BlockChain struct {
	LastKey []byte  //表中最新的的一条数据的key（持有最后一个区块的数据）
	BoltDB *bolt.DB //DB对象
}



//打开 bolt 文件库(DbName 为新建bolt文件库的名字)
func Open(DbName string) error {
	db, err := bolt.Open(DbName, 0600, nil)
	if err != nil {
		return err
	}
	BoltDB = db
	return nil
}

//写操作（参数为name, key, value [数据表的名字, 存入时的关键字， 存入时的值{block序列化后的信息}]）
func AddDate(tableName []byte, key []byte, value []byte) error {
	err := BoltDB.Update(func(tx *bolt.Tx) error {
		// 1.创建一个数据桶， 或判断一个数据桶是否存在
		bucket, err := tx.CreateBucketIfNotExists(tableName)
		if err != nil {
			fmt.Println("创建一个数据桶出现错误请重试")
			return err
		}
		/**
		 * 2.存入数据（数据从参数获取）
		 * 2.1 判断本次存入的数据是否存在，不存在则存入数据，存在则直接退出
		 * 2.2 数据不存在时添加该条数据，并刷新lastValue的值
		 * 2.3 数据存在
		 */
		thisTimeValue := bucket.Get(key)
		if thisTimeValue == nil { //该条数据不存在
			_ = bucket.Put(key, value)
			LastKey = key
			_ = bucket.Put(LastKey, value)
		}
		return nil
	})
	return err
}

//读操作
func Quary() error {
	err := BoltDB.View(func(tx *bolt.Tx) error {
		// 1.拿到一个数据桶
		bucket := tx.Bucket([]byte("blocks"))
		// 2.判断该桶是否存在，存在则拿到桶中的所有数据，无则返回nil
		if bucket != nil {
			//那到桶中的数据
			blockByte := bucket.Get(block.Hash)
			//将数据反序列化
			block, _ := deSerialize(blockByte, block)
			bk := *block
			fmt.Printf("%x", bk.Hash)

		}
		return nil
	})
	return err
}
