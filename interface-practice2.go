package main

func main() {
	for _, usemock := range []bool{true, false} {
		db := NewDB(usemock)
		db.Get("foo")
		db.Set("foo", "bar")
	}
}

type DB interface {
	Get(string) string
	Set(string, string)
}

type MockDB struct{}

func (m *MockDB) Get(key string) string {
	return "hoge"
}

func (m *MockDB) Set(key string, val string) {
}

type RealDB struct{}

func (m *RealDB) Get(key string) string {
	return "REALDATA"
}

func (m *RealDB) Set(key string, val string) {
}

func NewDB(usemock bool) DB {
	if usemock {
		return &MockDB{}
	}
	return &RealDB{}
}
