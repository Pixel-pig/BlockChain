package errorPk

import "errors"

func ALREADYEXISTS() error {
	return errors.New("该区块在区块链上已存在!")
}

func SERIALIZATIONFAILED() error {
	return errors.New("数据序列化失败!")
}

func ISEMPTY() error {
	return errors.New("该数据表已存在创世区块，请新建立一个数据表!")
}