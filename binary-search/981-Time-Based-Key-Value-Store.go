type TimeMap struct {
    dict map[string][][]interface{}
}

func Constructor() TimeMap {
    return TimeMap{
        dict: make(map[string][][]interface{}),
    }
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
    if _, ok := tm.dict[key]; !ok {
        tm.dict[key] = [][]interface{}{}
    }
    tm.dict[key] = append(tm.dict[key], []interface{}{value, timestamp})
}

func (tm *TimeMap) Get(key string, timestamp int) string {
    ans := ""
    values, ok := tm.dict[key]
    if !ok {
        return ans
    }
    left, right := 0, len(values)-1
    for left <= right {
        mid := (left + right) / 2
        if values[mid][1].(int) <= timestamp {
            left = mid + 1
            ans = values[mid][0].(string)
        } else {
            right = mid - 1
        }
    }
    return ans
}