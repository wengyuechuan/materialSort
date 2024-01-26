package utils

import (
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

// GetPageTitle 获取网页的标题
func GetPageTitle(url string) (string, error) {
	// 发送 http get 请求
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 解析HTML
	// 解析HTML
	tokenizer := html.NewTokenizer(resp.Body)
	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			return "", tokenizer.Err()
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()

			if token.Data == "title" {
				// 读取<title>标签中的内容
				tokenType = tokenizer.Next()
				if tokenType == html.TextToken {
					return strings.TrimSpace(tokenizer.Token().Data), nil
				}
			}
		}
	}
}
