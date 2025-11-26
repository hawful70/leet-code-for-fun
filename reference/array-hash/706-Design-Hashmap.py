class Pair:
    def __init__(self, key: int, value: int):
        self.key = key
        self.value = value

class MyHashMap:
    def __init__(self):
        self.bucket_size = 1000
        self.buckets = [[] for _ in range(self.bucket_size)]

    def hash(self, key: int) -> int:
        return key % self.bucket_size

    def put(self, key: int, value: int) -> None:
        index = self.hash(key)
        bucket = self.buckets[index]
        for pair in bucket:
            if pair.key == key:
                pair.value = value
                return
        bucket.append(Pair(key, value))

    def get(self, key: int) -> int:
        index = self.hash(key)
        bucket = self.buckets[index]
        for pair in bucket:
            if pair.key == key:
                return pair.value
        return -1

    def remove(self, key: int) -> None:
        index = self.hash(key)
        bucket = self.buckets[index]
        for i, pair in enumerate(bucket):
            if pair.key == key:
                bucket.pop(i)
                return
        


# Your MyHashMap object will be instantiated and called as such:
# obj = MyHashMap()
# obj.put(key,value)
# param_2 = obj.get(key)
# obj.remove(key)