package main

import (
	"strings"
)

func FaBing(name string, num int) string {
	var fadian = `{name}……🤤
	嘿嘿………🤤,
	……好可爱……嘿嘿……{name}🤤,
	……{name}……我的🤤,
	……嘿嘿……🤤,
	………亲爱的……赶紧让我抱一抱……啊啊啊{name}软软的脸蛋🤤还有软软的小手手……🤤,
	…{name}……不会有人来伤害你的…🤤,
	你就让我保护你吧嘿嘿嘿嘿嘿嘿嘿嘿🤤,
	……太可爱了……🤤,
	……美丽可爱的{name}……像珍珠一样……🤤,
	嘿嘿……{name}……🤤,
	嘿嘿……🤤,
	……好想一口吞掉……🤤……但是舍不得啊……我的{name}🤤,
	……嘿嘿……🤤,
	我的宝贝……我最可爱的{name}……🤤,
	没有{name}……我就要死掉了呢……🤤,
	我的……🤤,
	嘿嘿……可爱的{name}……嘿嘿🤤,
	……可爱的{name}……嘿嘿🤤🤤,
	……可爱的{name}……🤤,
	……嘿嘿🤤,
	……可爱的{name}…（吸）身上的味道……好好闻～🤤,
	…嘿嘿🤤,
	……摸摸～……可爱的{name}……再贴近我一点嘛……（蹭蹭）嘿嘿🤤,
	……可爱的{name}……嘿嘿🤤,
	……～亲一口～……可爱的{name}……嘿嘿🤤,
	……抱抱你～可爱的{name}～（舔）喜欢～真的好喜欢～……（蹭蹭）🤤,
	脑袋要融化了呢～已经……除了{name}以外～什么都不会想了呢～🤤,
	嘿嘿🤤,
	……可爱的{name}……嘿嘿🤤,
	……可爱的{name}……我的～……嘿嘿🤤`
	name1 := strings.ReplaceAll(fadian, "{name}", name)
	if string(num) == "" {
		return name1
	} else {
		buf := make([]byte, 0, num*len(name1))
		for i := 0; i < num; i++ {
			buf = append(buf, name1...)
		}
		return string(buf)
	}
}