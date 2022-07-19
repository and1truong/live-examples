package counter

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	
	"github.com/jfyne/live"
)

func Test_Handler(t *testing.T) {
	store := live.NewCookieStore("session-name", []byte("weak-secret"))
	engine := live.NewHttpHandler(store, NewHandler())
	// http.HandleFunc("/test", engine.ServeHTTP)
	
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	engine.ServeHTTP(rr, req)
	
	res := rr.Result()
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	
	fmt.Println(string(data))
}
