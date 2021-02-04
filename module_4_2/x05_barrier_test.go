package barrier

import (
	"strings"
	"testing"
)

//go test x05_barrier* -run=TestBarrier/Correct_endpoints -v
//go test x05_barrier* -run=TestBarrier/One_endpoint_incorrect -v
// go test x05_barrier* -run=TestBarrier/Very_short_timeout -v

func TestBarrier(t *testing.T) {
	t.Run("Correct endpoints", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/User-Agent"}

	result := captureBarrierOutput(endpoints...)
	if !strings.Contains(result, "Accept-Encoding") || !strings.Contains(result, "User-Agent") {
		t.Fail()
	}
	t.Log(result)

	})

	t.Run("One endpoint incorrect", func(t *testing.T) {
		endpoints := []string{"http://incorrect-url", "http://httpbin.org/User-Agent"}

	result := captureBarrierOutput(endpoints...)
	if !strings.Contains(result, "ERROR") {
		t.Fail() }
	t.Log(result)
	})

	t.Run("Very short timeout", func(t *testing.T) {
		endpoints := []string{"http://httpbin.org/headers", "http://httpbin.org/User-Agent"}
		timeoutMilliseconds = 1
		result := captureBarrierOutput(endpoints...)
		if !strings.Contains(result, "Timeout") {
			t.Fail() }
		t.Log(result)
	})
}