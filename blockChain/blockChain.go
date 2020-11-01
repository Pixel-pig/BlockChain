//bolt 操作

package blockChain

import (
	"blockChainProject/errorPk"
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
	LastHash []byte   //表中最新的的一条数据的hash值
	BoltDB   *bolt.DB //DB对象
}

const (
	BOLTDBFILENAME = "blockchain.db"
	LASTKEY        = "lastkey"
)

var BUCKETNAME = []byte("blockchain")

//实例化一条区块链
func NewBlockChain() (*BlockChain, error) {
	var bc BlockChain
	//1.拿到BlotDB链接
	db, err := bolt.Open(BOLTDBFILENAME, 0600, nil)
	if err != nil {
		panic("bolt数据库创建失败！")
	}

	//2.判断boltDB文件中是否存在数据表
	err = db.Update(func(tx *bolt.Tx) error {
		//3.1 判断DB文件中是否存在BUCKETNAME文件，不存在则创建该问件
		bucket := tx.Bucket(BUCKETNAME)
		if bucket == nil {
			//DB文件中不存在BUCKETNAME数据表
			_, err := tx.CreateBucket(BUCKETNAME)
			if err != nil {
				return err
			}
			//创建创世区块
			genesisBlock := NewGenesisBlock()
			//区块上链
			bc.SaveBlock(genesisBlock)
		} else {
			//DB文件中存在BUCKETNAME数据表,查看数据表是否存在创世区块，不存在则添加创世区块，存在测退出
			//利用最后一个lastkey去判断数据表中是否存在数据
			lastHash := bucket.Get([]byte(LASTKEY))
			if lastHash == nil {
				//不存在创世区块,创建创世区块并存入到区块链上
				genesisBlock := NewGenesisBlock()
				_, err := bc.SaveBlock(genesisBlock)
				if err != nil {
					return err
				}
			} else{
				//存在创世,拿到创世区块的值
				return errorPk.ISEMPTY()
			}
		}
		return nil
	})
	//4. 实例化一条区块链
	bc = BlockChain{
		BoltDB:   db,
	}

	return &bc, err
}

//存储block
func (bc BlockChain) SaveBlock(block Block) (*Block, error) {
	//1.拿到bolt链接
	db := bc.BoltDB
	err := db.Update(func(tx *bolt.Tx) error {
		//2.拿到数据表
		bucket := tx.Bucket(BUCKETNAME)
		//判断数据表中是否存在该区块
		thisBlockByte := bucket.Get(block.Hash)
		if thisBlockByte != nil {
			return errorPk.ALREADYEXISTS()
		}
		//3.将block数据序列化
		blockByte, err := block.Serialize()
		if err != nil {
			return errorPk.SERIALIZATIONFAILED()
		}
		//4.将数据添加block链上，并更新lastkey中的数据
		err = bucket.Put(block.Hash,blockByte)
		if err != nil {
			fmt.Println("数据上链失败")
			return err
		}
		_ = bucket.Put([]byte(LASTKEY),block.Hash)
		bc.LastHash = block.Hash
		return nil
	})
	return &block, err
}

//根据区块高度查询区块
func (bc BlockChain) QuaryBlockByHeight(height int64) *Block {
	//拿到bolt链接
	db := bc.BoltDB
	var block1 *Block
	_ = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(BUCKETNAME)
		lastHash := bc.LastHash
		for {
			blockByte := bucket.Get(lastHash)
			//反序列化
			block1, _ = DeSerialize(blockByte)
			lastHash = block1.PrevHash
			if block1.Height == height {
				break
			}
		}
		return nil
	})
	return block1
}

//查询链上所有的区块
func (bc BlockChain) QuaryAllBlock() []*Block {
	//拿到bolt链接
	db := bc.BoltDB
	//新建一个大整数用于剔除lastkey这个特殊的键值对
	//var bigInt = new(big.Int)
	//存储所遍历到的所有区块
	blocks := make([]*Block, 0)
	_ = db.View(func(tx *bolt.Tx) error {
		//拿到存储数据的数据表
		bucket := tx.Bucket(BUCKETNAME)
		//对数据表进行遍历
		_ = bucket.ForEach(func(k, v []byte) error {
			fmt.Println("k=",string(k),"v=",string(v))
			//key := bigInt.SetBytes(k)
			//lastKey := bigInt.SetBytes([]byte(LASTKEY))
			//if key.Cmp(lastKey) != 0 {
			//	//反序列化
			//	block1, _ := DeSerialize(v)
			//	fmt.Println("k=",k,"v=",block1)
			//	blocks = append(blocks, block1)
			//}

			return nil
		})
		return nil
	})
	return blocks
}

