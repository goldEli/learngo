package geecache

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const defaultPath = "/_geecache/"

type HTTPPool struct {
	self     string
	basePath string
}

func NewHTTPPool(self string) *HTTPPool {
	return &HTTPPool{
		self:     self,
		basePath: defaultPath,
	}
}

func (p *HTTPPool) log(format string, v ...interface{}) {
	log.Printf("[Server %s] %s", p.self, fmt.Sprintf(format, v...))

}

func (p *HTTPPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 判断url 是否以 defaultPath
	if !strings.HasPrefix(r.URL.Path, p.basePath) {
		panic("HTTPPool serving unexpected path:" + r.URL.Path)
	}
	p.log("%s %s", r.Method, r.URL.Path)
	parts := strings.SplitN(r.URL.Path[len(p.basePath):], "/", 2)

	if len(parts) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	groupName := parts[0]
	key := parts[1]

	group := GetGroup(groupName)

	if group == nil {
		http.Error(w, "no such group", http.StatusNotFound)
		return
	}

	view, err := group.Get(key)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(view.ByteSlice())
}
