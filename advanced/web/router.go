package web

import (
	"fmt"
	"strings"
)

type router struct {
	trees map[string]*node
}

func newRouter() router {
	return router{
		trees: map[string]*node{},
	}
}

func (r *router) addRoute(method, path string, handler HandleFunc) {
	r.validatePath(path)
	root, ok := r.trees[method]
	if !ok {
		root = &node{
			path: "/",
		}
		r.trees[method] = root
	}
	if path == "/" {
		if root.handler != nil {
			panic("web: 路由冲突 ")
		}
		root.handler = handler
		return
	}
	segs := strings.Split(path[1:], "/")
	n := root
	for _, seg := range segs {
		if seg == "" {
			panic(fmt.Sprintf("web: 非法路由。不允许使用 //a/b, /a//b 之类的路由, [%s]", path))
		}
		n = n.getOrCreateChild(seg)
	}
	if n.handler != nil {
		panic("web: 路由冲突")
	}
	n.handler = handler
}

func (r *router) validatePath(path string) {
	if path == "" {
		panic("web: 路由是空字符串")
	}
	if path[0] != '/' {
		panic("web: 路由必须以 / 开头")
	}
	if path != "/" && path[len(path)-1] == '/' {
		panic("web: 路由不能以 / 结尾")
	}
}

func (r *router) findRoute(method, path string) (*node, bool) {
	root, ok := r.trees[method]
	if !ok {
		return nil, false
	}
	if path == "/" {
		return root, true
	}
	segs := strings.Split(strings.Trim(path, "/"), "/")
	n := root
	for _, seg := range segs {
		n, ok = n.getChild(seg)
		if !ok {
			return nil, false
		}
	}
	return n, true
}
