package tft

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

type buff struct {
	attrs
	name   string
	g      *ground
	reduce action
	remain int
	end    int
	unique bool // 镜像，即可重复存在但持续时间单独计算
}

func newB(remain int, a ...*attrs) *buff {
	bf := &buff{remain: remain}
	for _, attr := range a {
		bf.attrs.Add(attr)
	}
	return bf
}

func (b *buff) key() string {
	if b.unique {
		return uuid.New().String()
	}
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
	if b.reduce == TimeGoA {
		return b.end >= b.g.ticks
	}
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
