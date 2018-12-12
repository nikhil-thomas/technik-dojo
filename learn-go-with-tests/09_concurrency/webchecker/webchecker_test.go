package webchecker

import (
	"reflect"
	"testing"
	"time"
)

func wcMock(url string) bool {
	if url == "hhhh://test.com" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"https://www.google.com",
		"hhhh://test.com",
		"https://www.github.com",
	}

	want := map[string]bool{
		"https://www.google.com": true,
		"hhhh://test.com":        false,
		"https://www.github.com": true,
	}

	got := CheckWebsites(wcMock, websites)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func slowWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "url"
	}
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowWebsiteChecker, urls)
	}
}

func TestCheckWebsitesWithChan(t *testing.T) {
	websites := []string{
		"https://www.google.com",
		"hhhh://test.com",
		"https://www.github.com",
	}

	want := map[string]bool{
		"https://www.google.com": true,
		"hhhh://test.com":        false,
		"https://www.github.com": true,
	}

	got := CheckWebsitesWithChan(wcMock, websites)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func BenchmarkCheckWebsitesWithChan(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "url"
	}
	for i := 0; i < b.N; i++ {
		CheckWebsitesWithChan(slowWebsiteChecker, urls)
	}
}
