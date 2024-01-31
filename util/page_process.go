package util

import (
	"golang.org/x/net/html"
	"io"
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

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

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

// GetPageDescription 获取网页的描述
func GetPageDescription(url string) (string, error) {
	// 发送 http get 请求
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// 解析HTML
	tokenizer := html.NewTokenizer(resp.Body)
	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			return "", tokenizer.Err()
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()

			if token.Data == "meta" {
				// 读取<meta>标签中的内容
				for _, attr := range token.Attr {
					if attr.Key == "name" && attr.Val == "description" {
						for _, attr := range token.Attr {
							if attr.Key == "content" {
								return strings.TrimSpace(attr.Val), nil
							}
						}
					}
				}
			}
		}
	}
}
