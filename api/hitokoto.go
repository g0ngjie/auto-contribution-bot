package api

import (
	"log"

	"github.com/go-resty/resty/v2"
)

// 一言

type HitokotoBody struct {
	Hitokoto string `json:"hitokoto"`
	From     string `json:"from"`
	Fromwho  string `json:"from_who"`
	Creator  string `json:"creator"`
}

// 获取一言
func GetHitokoto() *HitokotoBody {
	client := resty.New()
	hitokoto := &HitokotoBody{}
	// v1.hitokoto.cn
	_, err := client.R().SetResult(hitokoto).Get("https://v1.hitokoto.cn")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return hitokoto
}
