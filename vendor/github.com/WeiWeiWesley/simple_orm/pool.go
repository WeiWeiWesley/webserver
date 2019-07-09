package orm

import (
	"sync"
	"errors"

	"github.com/jinzhu/gorm"
)

type pool struct {
	Pool map[string]*DB
	mu *sync.Mutex
}

//DB db conn
type DB struct {
	Conn *gorm.DB
}

var (
	//DbPool 連線池
	DbPool pool
)

func init() {
	DbPool = pool{
		Pool: make(map[string]*DB),
		mu: &sync.Mutex{},
	}
}

//AddPool add connection to pool map
func (p *pool) AddPool(connName string, conn *gorm.DB) {
	p.mu.Lock()
	defer p.mu.Unlock()

	DbPool.Pool[connName] = &DB{
		Conn: conn,
	}
}

//GetDbConn Fetch db conncetion
func GetDbConn(connName string) (*DB, error) {
	if db, ok := DbPool.Pool[connName]; ok {
		return db, nil
	}

	return nil, errors.New("DB conn not exists")
}
