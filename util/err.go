package util

import "errors"

func ALREADYEXISTS() error {
	return errors.New("该区块在区块链上已存在")
}
