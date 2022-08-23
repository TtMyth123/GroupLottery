package Game28ResultServer

import (
	"fmt"
	"github.com/TtMyth123/kit/stringKit"
	"testing"
)

func TestMap2AAA(t *testing.T) {
	resurl := `{
    "gameCode": "jndpc28",
    "preIssue": "2697269",
    "openNum": [
        7,
        3,
        9
    ],
    "dragonTigerArr": [],
    "sumArr": [
        19,
        1,
        1,
        0
    ],
    "issue": "2697270",
    "currentOpenDateTime": 1617814500000,
    "openDateTime": 1617814710000,
    "serverTime": 1617814521950,
    "openedCount": 287,
    "dailyTotal": 396,
    "formArr": [],
    "mimcryArr": [],
    "zodiacArr": [],
    "compareArr": [],
    "sumType": null,
    "wuxing": null
}`
	r, e := kjw6589Result2Game28Result(resurl)

	fmt.Println(stringKit.GetJsonStr(r), e)
}
