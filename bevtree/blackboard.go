package bevtree

import (
	"errors"
)

var (
	ErrInvalidKey  = errors.New("invalid key")
	ErrInvalidType = errors.New("invalid type")
)

// 简单单例模式
var bb *BlackBoard

func GetBlackboard() *BlackBoard {
	if bb == nil {
		bb = newBlackboard()
	}
	return bb
}

/*
	黑板结构体
	提供基础类型的Get Set方法用于节点间共享数据
 */
type BlackBoard struct {
	board map[int]interface{}
}

func newBlackboard() *BlackBoard {
	mBoard := make(map[int]interface{}, 1000)
	return &BlackBoard{board: mBoard}
}

func (b *BlackBoard) GetValueAsBool(key int) (bool, error) {
	if v, ok := b.board[key]; ok {
		if i, ok := v.(bool); ok {
			return i, nil
		}
		return false, ErrInvalidType
	}
	return false, ErrInvalidKey
}

func (b *BlackBoard) SetValueAsBool(key int, v bool) {
	b.board[key] = v
}

func (b *BlackBoard) GetValueAsInt(key int) (int, error) {
	if v, ok := b.board[key]; ok {
		if i, ok := v.(int); ok {
			return i, nil
		}
		return 0, ErrInvalidType
	}
	return 0, ErrInvalidKey
}

func (b *BlackBoard) SetValueAsInt(key int, v int) {
	b.board[key] = v
}

func (b *BlackBoard) GetValueAsFloat32(key int) (float32, error) {
	if v, ok := b.board[key]; ok {
		if i, ok := v.(float32); ok {
			return i, nil
		}
		return 0, ErrInvalidType
	}
	return 0, ErrInvalidKey
}

func (b *BlackBoard) SetValueAsFloat32(key int, v float32) {
	b.board[key] = v
}

func (b *BlackBoard) GetValueAsFloat64(key int) (float64, error) {
	if v, ok := b.board[key]; ok {
		if i, ok := v.(float64); ok {
			return i, nil
		}
		return 0, ErrInvalidType
	}
	return 0, ErrInvalidKey
}

func (b *BlackBoard) SetValueAsFloat64(key int, v float64) {
	b.board[key] = v
}

func (b *BlackBoard) GetValueAsString(key int) (string, error) {
	if v, ok := b.board[key]; ok {
		if i, ok := v.(string); ok {
			return i, nil
		}
		return "", ErrInvalidType
	}
	return "", ErrInvalidKey
}

func (b *BlackBoard) SetValueAsString(key int, v string) {
	b.board[key] = v
}

func (b *BlackBoard) GetValueAsInterface(key int) (interface{}, error) {
	if v, ok := b.board[key]; ok {
		return v, nil
	}
	return nil, ErrInvalidKey
}

func (b *BlackBoard) SetValueAsInterface(key int, v interface{}) {
	b.board[key] = v
}
