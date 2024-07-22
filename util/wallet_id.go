package util

import (
	"regexp"
	"sync"
	"time"
)

// SnowflakeIDGenerator 结构体定义
type SnowflakeIDGenerator struct {
	mu           sync.Mutex
	epoch        int64 // 起始时间戳，用于计算时间戳部分
	nodeID       int64 // 节点ID
	sequence     int64 // 序列号
	lastGenTime  int64 // 上次生成ID的时间戳
	sequenceBits int64 // 序列号所占的位数
	nodeIDBits   int64 // 节点ID所占的位数
}

// NewSnowflakeIDGenerator 创建一个新的SnowflakeIDGenerator
func NewSnowflakeIDGenerator(epoch, nodeID, sequenceBits, nodeIDBits int64) *SnowflakeIDGenerator {
	gen := &SnowflakeIDGenerator{
		epoch:        epoch,
		nodeID:       nodeID,
		sequenceBits: sequenceBits,
		nodeIDBits:   nodeIDBits,
	}
	if gen.sequenceBits <= 0 {
		gen.sequenceBits = 12
	}
	if gen.nodeIDBits <= 0 {
		gen.nodeIDBits = 10
	}
	return gen
}

// Generate 生成一个雪花ID
func (gen *SnowflakeIDGenerator) Generate() int64 {
	gen.mu.Lock()
	defer gen.mu.Unlock()

	currentTime := time.Now().UnixNano() / 1e6
	if currentTime < gen.lastGenTime {
		panic("Invalid system clock!")
	}

	if currentTime == gen.lastGenTime {
		gen.sequence = (gen.sequence + 1) & ((1 << gen.sequenceBits) - 1)
		if gen.sequence == 0 {
			// 当序列号溢出时，等待到下一毫秒
			for currentTime <= gen.lastGenTime {
				currentTime = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		gen.sequence = 0
	}

	gen.lastGenTime = currentTime

	id := ((currentTime - gen.epoch) & ((1 << (64 - gen.sequenceBits - gen.nodeIDBits)) - 1)) << (gen.sequenceBits + gen.nodeIDBits)
	id |= gen.nodeID << gen.sequenceBits
	id |= gen.sequence

	return id
}

func ValidatePassword(password string) bool {
	// 使用正则表达式匹配密码规则
	match, _ := regexp.MatchString("(^(?=.*[a-zA-Z])(?=.*\\d).{8,16}$)|(^\\d{8,16}$)", password)
	return match
}
