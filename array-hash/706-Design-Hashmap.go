type Pair struct {
	Key   int
	Value int
}

type MyHashMap struct {
	buckets    [][]Pair
	bucketSize int
}

func Constructor() MyHashMap {
	return MyHashMap{
		buckets:    make([][]Pair, 1000),
		bucketSize: 1000,
	}
}

func (this *MyHashMap) hash(key int) int {
	return key % this.bucketSize
}

func (this *MyHashMap) Put(key int, value int) {
	index := this.hash(key)
	bucket := this.buckets[index]
	for i, pair := range bucket {
		if pair.Key == key {
			bucket[i].Value = value
			return
		}
	}
	this.buckets[index] = append(bucket, Pair{Key: key, Value: value})
}

func (this *MyHashMap) Get(key int) int {
	index := this.hash(key)
	for _, pair := range this.buckets[index] {
		if pair.Key == key {
			return pair.Value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	index := this.hash(key)
	bucket := this.buckets[index]
	for i, pair := range bucket {
		if pair.Key == key {
			this.buckets[index] = append(bucket[:i], bucket[i+1:]...)
			return
		}
	}
}

/**
* Your MyHashMap object will be instantiated and called as such:
* obj := Constructor();
* obj.Put(key,value);
* param_2 := obj.Get(key);
* obj.Remove(key);
 */