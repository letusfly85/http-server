package main

import "testing"

// ファイルの存在チェックエラー対応
func TestGetConfig(t *testing.T) {
	cfg, err := GetConfig("dummy.gcfg")

	if err == nil {
		t.Errorf("expected error, but got %v", cfg.ToString())
	}
}
