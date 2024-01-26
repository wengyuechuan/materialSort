package test

import (
	"materialsSort/utils"
	"testing"
)

func TestGetPageTitle(t *testing.T) {
	var tests = []struct {
		url   string
		title string
	}{
		{"https://www.baidu.com", "百度一下，你就知道"},
		{"https://www.bilibili.com/video/BV18e411n7ka", "恶心吐了！全网最恶心的擦边球！【土味视频】_哔哩哔哩_bilibili"},
	}
	for _, test := range tests {
		if got, _ := utils.GetPageTitle(test.url); got != test.title {
			t.Errorf("GetPageTitle(’%s‘) = %s", test.url, got)
		} else {
			t.Logf("GetPageTitle(’%s‘) = %s", test.url, got)
		}
	}
}
