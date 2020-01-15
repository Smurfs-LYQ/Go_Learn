package session

import (
	"fmt"
)

// AddSessionData 新增session数据
func (s *SessionData) AddSessionData(key string, val interface{}) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	s.Data[key] = val
}

// DelSessionData 删除session数据
func (s *SessionData) DelSessionData(key string) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	delete(s.Data, key)
}

// SetSessionData 修改session数据
func (s *SessionData) SetSessionData(key string, val interface{}) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	s.Data[key] = val
}

// GetSessionData 获取session数据
func (s *SessionData) GetSessionData(key string) (val interface{}, err error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	val, ok := s.Data[key]
	if !ok {
		return nil, fmt.Errorf("无法获取到 \"%s\" 对应的值", key)
	}

	return
}
