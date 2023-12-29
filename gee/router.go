package gee

import "strings"

type router struct {
	routers  map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		routers:  make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	parts := strings.Split(pattern, "/")
	res := make([]string, 0)
	for _, item := range parts {
		if item == "" {
			continue
		}
		res = append(res, item)
		if item[0] == '*' {
			break
		}
	}
	return res
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	if _, ok := r.routers[method]; !ok {
		r.routers[method] = &node{}
	}
	r.routers[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.routers[method]
	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}

			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(404, "404 NOT FOUND:%s\n", c.Path)
	}
}
