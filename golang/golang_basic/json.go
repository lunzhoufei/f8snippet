package golang_basic

import (
	"encoding/json"
	"git.code.oa.com/going/srf/srfs"
)

func prettyPrint(s interface{}) (s string) {
	if str, err := json.MarshalIndent(s, "", "  "); err == nil {
		s = string(str)
	} else {
		s = "json parse failed! raw:"
	}
}

func print(s interface{}) (s string) {
	if str, err := json.Marshal(s); err == nil {
		s = string(str)
	} else {
		s = "json parse failed! raw:"
	}
}
