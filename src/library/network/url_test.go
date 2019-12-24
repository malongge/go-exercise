package network

import (
	"fmt"
	"net/url"
	"testing"
)

func TestUrl(t *testing.T) {
	var testUrl = "100%e8%a5%bf%e5%ae%89%e4%b8%ad%e5%b0%8f%e5%ad%a62019%e5%8d%87%e5%ad%a6%e6%96%b0%e6%94%bf%e6%9d%a5%e4%ba%86%ef%bc%8c%e6%b6%88%e6%81%af%e4%b8%80%e5%87%ba%ef%bc%8c%e4%bc%a0%e9%81%8d%e5%ae%b6%e9%95%bf%e4%bb%ac%e6%9c%8b%e5%8f%8b%e5%9c%88.mp4"

	escapeUrl := url.QueryEscape("你好")
	fmt.Println(escapeUrl)
	resUri, err := url.QueryUnescape(testUrl)
	if err == nil {
		fmt.Println(resUri)
	}



}
