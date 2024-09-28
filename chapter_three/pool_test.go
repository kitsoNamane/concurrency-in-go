package chapter_three

import (
	"io"
	"net"
	"testing"
)

func init() {
	daemonStarted := StartNetworkDaemon()
	daemonStarted.Wait()
}

func BenchmarkNetworkRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:9090")
		if err != nil {
			b.Fatalf("cannot dial host: %v", err)
		}

		if _, err := io.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}
		conn.Close()
	}
}
