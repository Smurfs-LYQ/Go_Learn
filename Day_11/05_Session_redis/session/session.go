package session

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const session_time int = 20

// Session 表示提个具体的用户session数据
type Session struct {
	ID   string
	Data map[string]interface{}
	lock sync.RWMutex
	// // 过期时间
	// MaxAge int
}

// InitSession 初始化Session类型
func InitSession() *Session {
	rand.Seed(time.Now().UnixNano())
	return &Session{
		ID:   fmt.Sprintf("%d%d", time.Now().UnixNano(), rand.Intn(10000)),
		Data: make(map[string]interface{}, 2),
	}
}

func (s *Session) Add(key string, val interface{}) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	s.Data[key] = val
}

func Save(data *Session) error {
	id := data.ID
	data, err := json.Marshal(data.Data)
	if err != nil {
		return err
	}

	fmt.Println(id, data)
}
