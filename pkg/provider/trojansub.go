package provider

import (
	"github.com/jth445600/proxypool/pkg/tool"
	"strings"
)

type TrojanSub struct {
	Base
}

func (sub TrojanSub) Provide() string {
	sub.Types = "trojan"
	sub.preFilter()
	var resultBuilder strings.Builder
	for _, p := range *sub.Proxies {
		resultBuilder.WriteString(p.Link() + "\n")
	}
	return tool.Base64EncodeString(resultBuilder.String(), false)
}
