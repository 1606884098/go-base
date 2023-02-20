package demo

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

//1.功能测试
func TestGetArea(t *testing.T) { //函数方法相当于是测试用例
	t.Helper() //报错时将输出帮助函数调用者的信息,调用 t.Helper() 让报错信息更准确，有助于定位
	area := GetArea(40, 50)
	//d := 1 / 0
	if area != 2000 {
		t.Error("测试失败")
	}
}

//2.压力测试
func BenchmarkGetArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetArea(40, 50)
	}
	//并发测试
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			// 所有 goroutine 一起，循环一共执行 b.N 次
			buf.Reset()
		}
	})
}

//3.子测试
func TestMul(t *testing.T) {
	t.Run("pos", func(t *testing.T) {
		if Mul(2, 3) != 6 {
			t.Fatal("fail")
		}

	})
	t.Run("neg", func(t *testing.T) {
		if Mul(2, -3) != -6 {
			t.Fatal("fail")
		}
	})
}

//4.测试组
func TestMulGroup(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Expected int
	}{
		{"pos", 2, 3, 6},
		{"neg", 2, -3, -6},
		{"zero", 2, 0, 0},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := Mul(c.A, c.B); ans != c.Expected {
				t.Fatalf("%d * %d expected %d, but %d got",
					c.A, c.B, c.Expected, ans)
			}
		})
	}
}

//5.网络测试
func TestConn(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	handleError(t, err)
	defer ln.Close()

	http.HandleFunc("/hello", helloHandler)
	go http.Serve(ln, nil)
	//尽量不对 http 和 net 库使用 mock，这样可以覆盖较为真实的场景
	resp, err := http.Get("http://" + ln.Addr().String() + "/hello")
	handleError(t, err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	handleError(t, err)

	if string(body) != "hello world" {
		t.Fatal("expected hello world, but got", string(body))
	}
}
func handleError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("failed", err)
	}
}

//针对 http 开发的场景，使用标准库 net/http/httptest 进行测试更为高效
func TestConnfHttp(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	helloHandler(w, req)
	bytes, _ := ioutil.ReadAll(w.Result().Body)

	if string(bytes) != "hello world" {
		t.Fatal("expected hello world, but got", string(bytes))
	}
}
