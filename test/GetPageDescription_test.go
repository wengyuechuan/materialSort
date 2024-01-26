package test

import (
	"materialsSort/utils"
	"testing"
)

func TestGetPageDescription(t *testing.T) {
	var tests = []struct {
		url         string
		description string
	}{
		{"https://www.baidu.com", ""},
		{"https://www.bilibili.com/video/BV18e411n7ka", "不会还有人不知道长按大拇指三连关注会成为b站lv6用户然后身体健康发大财吧哎嘿\\(^o^)/~, 视频播放量 869468、弹幕量 5654、点赞数 37411、投硬币枚数 27635、收藏人数 12542、转发人数 2379, 视频作者 棉白有点笨, 作者简介 商业合作等等东西加vx：mib820我想150万粉丝呜呜呜！你们能不能进来看看我！ ，相关视频：我把DIO做成了奶冻！快进来打一下！哈哈哈哈哈哈，永久封号的三大网红，手感很好，可惜是男性，“嫩是真的嫩！”，哪个人才发明的实用性用品，20个监控下的尴尬瞬间，这个女人的行为把我看懵了！，日本地雷系整形博主的医美记录～，“有些人骨子里的骚…是瞒不住的！”，真实的人眼标本，【飞碟社】没见过这么离谱的原神动画OP ！"},
	}
	for _, test := range tests {
		if got, _ := utils.GetPageDescription(test.url); got != test.description {
			t.Errorf("GetPageDescription(’%s‘) = %s", test.url, got)
		} else {
			t.Logf("GetPageDescription(’%s‘) = %s", test.url, got)
		}
	}
}
