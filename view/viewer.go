package view

import (
	"fuckdb/bases/model"
	"io"
	"sync"
)

var pool = make(map[string]Viewer)
var mux sync.Mutex

// Register 注册视图
func Register(name string, viewer Viewer) {
	mux.Lock()
	defer mux.Unlock()
	if viewer == nil {
		panic("Register viewer is nil")
	}
	if _, dup := pool[name]; dup {
		panic("Register called twice for viewer " + name)
	}
	pool[name] = viewer
}

// Registered 返回已注册的视图名称列表
func Registered() (names []string) {
	mux.Lock()
	defer mux.Unlock()

	for key := range pool {
		names = append(names, key)
	}
	return names
}

// SelectViewer 根据视图名称从已注册的视图中进行查找，若指定视图不存在，则返回nil。
func SelectViewer(name string) Viewer {
	mux.Lock()
	defer mux.Unlock()

	for key := range pool {
		if key == name {
			return pool[key]
		}
	}
	return nil
}

// Viewer 数据视图
type Viewer interface {
	Do(items []model.DB, out io.Writer) error
}
