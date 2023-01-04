package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func matchNum(key string, exp float64, data map[string]interface{}) bool {
	if v, ok := data[key]; ok {
		if val, ok := v.(float64); ok && val == exp {
			return true
		}
	}

	return false
}

func matchString(key string, exp string, data map[string]interface{}) bool {
	if v, ok := data[key]; ok {
		if val, ok := v.(string); ok && strings.EqualFold(val, exp) {
			return true
		}
	}

	return false
}

func contains(known, unknown map[string]interface{}) error {
	for k, v := range known {
		switch x := v.(type) {
		case float64:
			if !matchNum(k, x, unknown) {
				return fmt.Errorf("%s unmatched (%d)", k, int(x))
			}
		case string:
			if !matchString(k, x, unknown) {
				return fmt.Errorf("%s unmatched (%s)", k, x)
			}
		case map[string]interface{}:
			if val, ok := unknown[k]; !ok {
				return fmt.Errorf("%s missing in resp", k)
			} else if unk, ok := val.(map[string]interface{}); ok {
				if err := contains(x, unk); err != nil {
					return fmt.Errorf("%s unmatched (%+v): %s", k, x, err)
				}
			} else {
				return fmt.Errorf("%s wrong in resp (%#v)", k, val)
			}
		}

	}

	return nil
}

func CheckData(known, unknown []byte) error {
	var k, u map[string]interface{}

	if err := json.Unmarshal(known, &k); err != nil {
		return err
	}

	if err := json.Unmarshal(unknown, &u); err != nil {
		return err
	}

	return contains(k, u)
}
