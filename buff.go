package tft

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type buff struct {
	attrs
	name   string
	g      *ground
	reduce action
	remain int
}

func newB(remain int, a ...*attrs) *buff {
	bf := &buff{remain: remain}
	for _, attr := range a {
		bf.attrs.Add(attr)
	}
	return bf
}

func (b *buff) key() string {
	segment1, _ := json.Marshal(b.attrs)
	segment2, _ := json.Marshal(b.reduce)
	// 创建 SHA256 哈希对象
	h := sha256.New()
	// 写入 JSON 数据
	h.Write(segment1)
	h.Write(segment2)
	// 获取哈希值
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (b *buff) valid() bool {
	return b.remain > 0
}

func (b *buff) body() *attrs {
	return &b.attrs
}

func (b *buff) process(e event) {
	if e.match(b.reduce) {
		b.remain--
	}
}
