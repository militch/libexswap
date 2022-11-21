package common

import (
	"encoding/json"
)

func SimpleObjCopy(src interface{}, dest interface{}){
    srcJsonRaw, _ := json.Marshal(src)
    _ = json.Unmarshal(srcJsonRaw, dest)
}
