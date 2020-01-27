package session

import (
	"Go_Learn/Day_11/05_Session_redis/db"
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Session_time 设置session生效时长
const Session_time int = 5

// Session 表示提个具体的用户session数据
type Session struct {
	ID   string
	Data map[string]interface{}
	lock sync.RWMutex
	// // 过期时间
	// MaxAge int
}

// InitSession 初始化Session对象
func InitSession() *Session {
	rand.Seed(time.Now().UnixNano())
	return &Session{
		ID:   fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Intn(1000)),
		Data: make(map[string]interface{}, 2),
	}
}

// Sel_Session 从Redis中查询UUID
func Sel_Session(id string) (*Session, error) {
	res, err := db.RedisDB.Get(id).Result()
	if err != nil {
		fmt.Println("Sel_Session err:", err)
		return nil, fmt.Errorf("reids not found")
	}

	sess := &Session{}

	// 将返回值反序列化
	err = json.Unmarshal([]byte(res), sess)
	if err != nil {
		fmt.Println("json Unmarshal err:", err)
		return nil, fmt.Errorf("json Unmarshal err")
	}

	return sess, err
}

// Add 增加信息
func (s *Session) Add(key string, val interface{}) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	s.Data[key] = val
}

// Del 删除信息
func (s *Session) Del(key string) (err error) {
	s.lock.Lock()
	defer s.lock.RUnlock()

	delete(s.Data, key)

	return
}

// Sel 查询信息
func (s *Session) Sel(key string) (interface{}, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	res, ok := s.Data[key]
	if !ok {
		return nil, fmt.Errorf("没找到指定的信息")
	}
	return res, nil
}

// Save 保存信息
func (s *Session) Save() (id string, err error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	json, err := json.Marshal(s)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	id = fmt.Sprintf("%d-%d", time.Now().UnixNano(), rand.Intn(10000))

	err = db.RedisDB.Set(id, string(json), 0).Err()
	if err != nil {
		return
	}

	return
}
