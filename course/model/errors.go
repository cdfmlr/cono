package model

import "errors"

/*
 * 这个文件中定义 model 中一些通用的有关错误
 */

var (
	// ErrNotExist 表示记录不存在于数据库的错误
	// TODO: Unused, to be deleted
	ErrNotExist = errors.New("record not exist")
)
