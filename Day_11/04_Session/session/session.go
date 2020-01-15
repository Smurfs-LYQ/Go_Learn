package session

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// SessionData 表示提个具体的用户session数据
type SessionData struct {
	ID   string
	Data map[string]interface{}
	lock sync.RWMutex
	// // 过期时间
	// MaxAge int
}

// SessionMgr 全局的session管理
type SessionMgr struct {
	Session map[string]*SessionData
	rwlock  sync.RWMutex
}

// NewSessionData SessionData构造函数
func NewSessionData(key string) *SessionData {
	return &SessionData{
		ID:   key,
		Data: make(map[string]interface{}, 8),
	}
}

// NewSessionMgr SessionMgr构造函数
func NewSessionMgr() *SessionMgr {
	return &SessionMgr{
		Session: make(map[string]*SessionData, 2),
	}
}

// AddSession 添加session信息
func (s *SessionMgr) AddSession() (key string) {
	s.rwlock.RLock()
	defer s.rwlock.Unlock()

	rand.Seed(time.Now().UnixNano())
	key = fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Intn(10000))

	s.Session[key] = NewSessionData(key)

	return
}

// DelSession 删除session信息
func (s *SessionMgr) DelSession(key string) {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	delete(s.Session, key)
}

// GetSession 获取session信息
func (s *SessionMgr) GetSession(key string) (sd *SessionData, err error) {
	s.rwlock.RLock()
	defer s.rwlock.RUnlock()

	sd, ok := s.Session[key]
	if !ok {
		err = fmt.Errorf("没有找到对应的信息%s", key)
		return
	}
	return
}
