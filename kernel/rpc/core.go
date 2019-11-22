package rpc

import (
	"errors"
	"fmt"
	"math/rand"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strings"
	"time"

	"webserver/kernel/common"
)

const (
	maxLink   = 10 // 最大連線數
	maxReCall = 3  // call 最大重試次數
)

// Link connection
type Link struct {
	Client *rpc.Client
	Lock   chan bool // 操作鎖
}

var (
	linkPool      map[string][maxLink]*Link
	serviceConfig []common.Service
)

func init() {
	linkPool = make(map[string][maxLink]*Link)
	serviceConfig = common.LoadConfig().Service
	for _, service := range serviceConfig {
		var tmp [maxLink]*Link
		for i := range tmp {
			tmp[i] = &Link{
				Lock: make(chan bool, 1),
			}
		}
		linkPool[service.Name] = tmp
	}
}

// FetchConn find connection
func FetchConn(serviceName string) (*Link, error) {
	if _, ok := linkPool[serviceName]; !ok {
		return nil, fmt.Errorf("Connection of service '%s' not exist", serviceName)
	}

	rand.Seed(time.Now().UnixNano())
	var retryTime time.Duration = 10
	if conn, err := linkPool[serviceName][rand.Intn(maxLink)].getConn(serviceName); err == nil {
		return conn, nil
	}

	// 嘗試取得其他連線
	for i := 0; i < 10; i++ {
		time.Sleep(retryTime + time.Duration(rand.Intn(500))*time.Millisecond)
		if conn, err := linkPool[serviceName][rand.Intn(maxLink)].getConn(serviceName); err == nil {
			return conn, nil
		}
	}

	return nil, errors.New("all occupied")
}

// GetConn get connection
func (c *Link) getConn(serviceName string) (*Link, error) {
	select {
	case c.Lock <- true:
		if c.Client == nil {
			for _, service := range serviceConfig {
				if service.Name == serviceName {
					c.Client, _ = jsonrpc.Dial("tcp", service.IP+":"+service.Port)
				}
			}
		}
		return c, nil
	default:
		return nil, errors.New("occupied")
	}
}

// PutBack put back connection
func (c *Link) PutBack() error {
	if c == nil {
		return nil
	}

	select {
	case <-c.Lock:
		return nil
	default:
		return errors.New("not in used")
	}
}

// Call 執行呼叫
func (c *Link) Call(serviceMethod string, args interface{}, reply interface{}) error {
	// call發生錯誤 嘗試重新連線再call
	err := c.Client.Call(serviceMethod, args, reply)
	if err != nil {
		for i := 0; i < maxReCall; i++ {
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			c.repairLink(strings.Split(serviceMethod, ".")[0])
			err = c.Client.Call(serviceMethod, args, reply)
			if err == nil {
				break
			}
		}
	}

	return err
}

// RepairLink 重新連線
func (c *Link) repairLink(serviceName string) (err error) {
	fmt.Println("Repair Link...")
	for _, service := range serviceConfig {
		if service.Name == serviceName {
			c.Client, _ = jsonrpc.Dial("tcp", service.IP+":"+service.Port)
		}
	}
	return
}
