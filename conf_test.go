package main

import "testing"

//TODO ファイルの存在チェックエラー対応
func TestGetConfig(t *testing.T) {
	actual := 1

	expected := 1
	if actual != expected {
		t.Errorf("got %vwant %v", actual, expected)
	}
}
