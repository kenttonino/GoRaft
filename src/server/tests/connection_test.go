package server_test

import (
	"GoRaft/src/server"
	"GoRaft/src/store"
	"bufio"
	"fmt"
	"net"
	"path/filepath"
	"testing"
)

func newTestServer(t *testing.T) *server.Server {
	t.Helper()
	s, err := store.New(filepath.Join(t.TempDir(), "wal.log"))
	if err != nil {
		t.Fatalf("store.New: %v", err)
	}
	t.Cleanup(func() { s.Close() })
	return server.New(":0", s)
}

// start runs the server in a background goroutine and returns its
// actual bound address. Registers t.Cleanup to stop it.
func start(t *testing.T, srv *server.Server) string {
	t.Helper()
	ready := make(chan string, 1)
	go func() {
		addr, err := srv.StartOnRandom(ready)
		if err != nil && addr == "" {
			t.Errorf("server exited with error: %v", err)
		}
	}()
	addr := <-ready
	t.Cleanup(func() { srv.Stop() })
	return addr
}

func dial(t *testing.T, addr string) (net.Conn, *bufio.Scanner) {
	t.Helper()
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	t.Cleanup(func() { conn.Close() })
	return conn, bufio.NewScanner(conn)
}

func send(t *testing.T, conn net.Conn, sc *bufio.Scanner, cmd string) string {
	t.Helper()
	fmt.Fprintln(conn, cmd)
	if !sc.Scan() {
		t.Fatalf("scanner closed after sending %q: %v", cmd, sc.Err())
	}
	return sc.Text()
}

func TestSetAndGet(t *testing.T) {
	conn, sc := dial(t, start(t, newTestServer(t)))

	if got := send(t, conn, sc, "SET name goraft"); got != "OK" {
		t.Errorf("SET: got %q, want OK", got)
	}
	if got := send(t, conn, sc, "GET name"); got != "goraft" {
		t.Errorf("GET: got %q, want goraft", got)
	}
}

func TestGetMissingKey(t *testing.T) {
	conn, sc := dial(t, start(t, newTestServer(t)))

	if got := send(t, conn, sc, "GET ghost"); got != "NULL" {
		t.Errorf("GET missing: got %q, want NULL", got)
	}
}

func TestDelete(t *testing.T) {
	conn, sc := dial(t, start(t, newTestServer(t)))

	send(t, conn, sc, "SET x 1")
	if got := send(t, conn, sc, "DEL x"); got != "OK" {
		t.Errorf("DEL: got %q, want OK", got)
	}
	if got := send(t, conn, sc, "GET x"); got != "NULL" {
		t.Errorf("GET after DEL: got %q, want NULL", got)
	}
}

func TestDeleteMissingKey(t *testing.T) {
	conn, sc := dial(t, start(t, newTestServer(t)))

	if got := send(t, conn, sc, "DEL ghost"); got != "OK" {
		t.Errorf("DEL missing key: got %q, want OK", got)
	}
}

func TestCaseInsensitiveCommands(t *testing.T) {
	conn, sc := dial(t, start(t, newTestServer(t)))

	if got := send(t, conn, sc, "set key val"); got != "OK" {
		t.Errorf("lowercase set: got %q, want OK", got)
	}
	if got := send(t, conn, sc, "get key"); got != "val" {
		t.Errorf("lowercase get: got %q, want val", got)
	}
}

func TestSetInvalidArity(t *testing.T) {
	conn, sc := dial(t, start(t, newTestServer(t)))

	if got := send(t, conn, sc, "SET onlykey"); got != "ERR usage: SET key value" {
		t.Errorf("SET arity: got %q", got)
	}
}

func TestGetInvalidArity(t *testing.T) {
	conn, sc := dial(t, start(t, newTestServer(t)))

	if got := send(t, conn, sc, "GET"); got != "ERR usage: GET key" {
		t.Errorf("GET arity: got %q", got)
	}
}

func TestDelInvalidArity(t *testing.T) {
	conn, sc := dial(t, start(t, newTestServer(t)))

	if got := send(t, conn, sc, "DEL"); got != "ERR usage: DEL key" {
		t.Errorf("DEL arity: got %q", got)
	}
}

func TestUnknownCommand(t *testing.T) {
	conn, sc := dial(t, start(t, newTestServer(t)))

	if got := send(t, conn, sc, "FLUSHALL"); got != "ERR unknown command: FLUSHALL" {
		t.Errorf("unknown command: got %q", got)
	}
}

func TestEmptyLinesIgnored(t *testing.T) {
	conn, sc := dial(t, start(t, newTestServer(t)))

	fmt.Fprintln(conn, "")
	fmt.Fprintln(conn, "   ")
	if got := send(t, conn, sc, "SET ping pong"); got != "OK" {
		t.Errorf("after empty lines: got %q, want OK", got)
	}
}

func TestOverwriteKey(t *testing.T) {
	conn, sc := dial(t, start(t, newTestServer(t)))

	send(t, conn, sc, "SET k v1")
	send(t, conn, sc, "SET k v2")
	if got := send(t, conn, sc, "GET k"); got != "v2" {
		t.Errorf("overwrite: got %q, want v2", got)
	}
}

func TestConcurrentClients(t *testing.T) {
	addr := start(t, newTestServer(t))
	const clients = 20

	done := make(chan struct{}, clients)
	for i := range clients {
		go func(i int) {
			defer func() { done <- struct{}{} }()
			conn, sc := dial(t, addr)
			key := fmt.Sprintf("key%d", i)
			val := fmt.Sprintf("val%d", i)
			if got := send(t, conn, sc, fmt.Sprintf("SET %s %s", key, val)); got != "OK" {
				t.Errorf("concurrent SET %s: got %q", key, got)
			}
			if got := send(t, conn, sc, fmt.Sprintf("GET %s", key)); got != val {
				t.Errorf("concurrent GET %s: got %q, want %q", key, got, val)
			}
		}(i)
	}
	for range clients {
		<-done
	}
}
