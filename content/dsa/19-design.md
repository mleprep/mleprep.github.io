# Design

**Pattern:** combine basic structures so that every operation hits its target complexity.

| Need | Combination |
|---|---|
| O(1) get + O(1) recency update | hash map → node in a **doubly linked list** (LRU) |
| O(1) insert / delete / random | array + value→index map, **swap with last** on delete |
| Frequency buckets | map key→freq, map freq→ordered set, track `minFreq` (LFU) |
| Versioned values | per-key list of `(version, value)` + binary search |
| FIFO with stacks | two stacks, move lazily (amortized O(1)) |

```viz
LRU:  map {k -> node}
      head <-> [most recent] <-> ... <-> [least recent] <-> tail
      get/put: move node to the front; evict tail.prev when over capacity
```

Other design problems elsewhere: Min Stack (155), Time Map (981), Twitter (355), Median Finder (295), Trie (208), BST Iterator (173), Serialize Tree (297).

## 146. LRU Cache | Medium
related: 460, 588
tags: blind75, hash-map, linked-list

Hash map from key to node of a doubly linked list ordered by recency. `get` moves the node to the front; `put` inserts at the front and evicts the node before the tail when over capacity. Sentinels (`head`, `tail`) remove edge cases. Python's `OrderedDict` packages exactly this.

**Complexity:** O(1) per operation.

```python
class LRUCache:
    def __init__(self, capacity: int):
        self.cap = capacity
        self.od = OrderedDict()      # hash map + doubly linked list in one

    def get(self, key: int) -> int:
        if key not in self.od:
            return -1
        self.od.move_to_end(key)     # most recent at the end
        return self.od[key]

    def put(self, key: int, value: int) -> None:
        self.od[key] = value
        self.od.move_to_end(key)
        if len(self.od) > self.cap:
            self.od.popitem(last=False)   # evict least recent
# test
c = LRUCache(2)
c.put(1, 1); c.put(2, 2)
assert c.get(1) == 1
c.put(3, 3)                  # evicts 2
assert c.get(2) == -1
c.put(4, 4)                  # evicts 1
assert c.get(1) == -1 and c.get(3) == 3 and c.get(4) == 4
```

```java
class LRUCache {
    private static class Node { int key, val; Node prev, next; Node(int k, int v) { key = k; val = v; } }

    private final int cap;
    private final Map<Integer, Node> map = new HashMap<>();
    private final Node head = new Node(0, 0), tail = new Node(0, 0); // head.next = most recent

    public LRUCache(int capacity) {
        cap = capacity;
        head.next = tail;
        tail.prev = head;
    }

    private void unlink(Node n) { n.prev.next = n.next; n.next.prev = n.prev; }

    private void addFront(Node n) {
        n.next = head.next; n.prev = head;
        head.next.prev = n; head.next = n;
    }

    public int get(int key) {
        Node n = map.get(key);
        if (n == null) return -1;
        unlink(n);
        addFront(n);
        return n.val;
    }

    public void put(int key, int value) {
        Node n = map.get(key);
        if (n != null) { n.val = value; unlink(n); addFront(n); return; }
        n = new Node(key, value);
        map.put(key, n);
        addFront(n);
        if (map.size() > cap) {
            Node lru = tail.prev;
            unlink(lru);
            map.remove(lru.key);
        }
    }
}
// test
LRUCache c = new LRUCache(2);
c.put(1, 1); c.put(2, 2);
check(c.get(1), 1);
c.put(3, 3);
check(c.get(2), -1);
c.put(4, 4);
check(c.get(1), -1);
check(c.get(3), 3);
check(c.get(4), 4);
```

## 460. LFU Cache | Hard
related: 146
tags: hash-map, frequency-buckets

`key → (value, freq)`, `freq → OrderedDict of keys` (LRU order within a frequency), plus `minFreq`. On access, move the key from bucket `f` to `f+1` (bump `minFreq` if bucket `f` emptied and was the min). On insert, evict the oldest key in `minFreq`'s bucket, then set `minFreq = 1`.

```viz
freq 1: [k3]            minFreq = 1  -> evict k3 first
freq 2: [k1, k4]        (k1 older than k4)
freq 5: [k2]
```

**Complexity:** O(1) per operation.

```python
class LFUCache:
    def __init__(self, capacity: int):
        self.cap = capacity
        self.kv = {}                                 # key -> [value, freq]
        self.buckets = defaultdict(OrderedDict)      # freq -> keys in LRU order
        self.min_f = 0

    def _touch(self, key):
        val, f = self.kv[key]
        del self.buckets[f][key]
        if not self.buckets[f]:
            del self.buckets[f]
            if self.min_f == f:
                self.min_f += 1
        self.kv[key][1] = f + 1
        self.buckets[f + 1][key] = None

    def get(self, key: int) -> int:
        if key not in self.kv:
            return -1
        self._touch(key)
        return self.kv[key][0]

    def put(self, key: int, value: int) -> None:
        if self.cap == 0:
            return
        if key in self.kv:
            self.kv[key][0] = value
            self._touch(key)
            return
        if len(self.kv) == self.cap:
            old, _ = self.buckets[self.min_f].popitem(last=False)
            del self.kv[old]
        self.kv[key] = [value, 1]
        self.buckets[1][key] = None
        self.min_f = 1
# test
c = LFUCache(2)
c.put(1, 1); c.put(2, 2)
assert c.get(1) == 1
c.put(3, 3)                       # evicts 2 (freq 1)
assert c.get(2) == -1 and c.get(3) == 3
c.put(4, 4)                       # 1 and 3 both freq 2 -> evict LRU among them = 1
assert c.get(1) == -1 and c.get(3) == 3 and c.get(4) == 4
```

```java
class LFUCache {
    private final int cap;
    private int minF = 0;
    private final Map<Integer, int[]> kv = new HashMap<>();                    // key -> {value, freq}
    private final Map<Integer, LinkedHashSet<Integer>> buckets = new HashMap<>(); // freq -> keys (LRU order)

    public LFUCache(int capacity) { cap = capacity; }

    private void touch(int key) {
        int[] e = kv.get(key);
        LinkedHashSet<Integer> b = buckets.get(e[1]);
        b.remove(key);
        if (b.isEmpty()) { buckets.remove(e[1]); if (minF == e[1]) minF++; }
        e[1]++;
        buckets.computeIfAbsent(e[1], k -> new LinkedHashSet<>()).add(key);
    }

    public int get(int key) {
        if (!kv.containsKey(key)) return -1;
        touch(key);
        return kv.get(key)[0];
    }

    public void put(int key, int value) {
        if (cap == 0) return;
        if (kv.containsKey(key)) { kv.get(key)[0] = value; touch(key); return; }
        if (kv.size() == cap) {
            LinkedHashSet<Integer> b = buckets.get(minF);
            int old = b.iterator().next();
            b.remove(old);
            if (b.isEmpty()) buckets.remove(minF);
            kv.remove(old);
        }
        kv.put(key, new int[]{value, 1});
        buckets.computeIfAbsent(1, k -> new LinkedHashSet<>()).add(key);
        minF = 1;
    }
}
// test
LFUCache c = new LFUCache(2);
c.put(1, 1); c.put(2, 2);
check(c.get(1), 1);
c.put(3, 3);
check(c.get(2), -1);
check(c.get(3), 3);
c.put(4, 4);
check(c.get(1), -1);
check(c.get(3), 3);
check(c.get(4), 4);
```

## 380. Insert Delete GetRandom O(1) | Medium
related: 381
tags: array-plus-map

Values live in an array (O(1) random pick); a map stores each value's index. To delete, **move the last element into the hole** and pop the end.

```viz
vals [5, 8, 2, 9]   idx {5:0, 8:1, 2:2, 9:3}
remove 8: move 9 into slot 1 -> [5, 9, 2]  idx {5:0, 9:1, 2:2}
```

**Complexity:** O(1) average per operation.

```python
class RandomizedSet:
    def __init__(self):
        self.vals, self.idx = [], {}

    def insert(self, val: int) -> bool:
        if val in self.idx:
            return False
        self.idx[val] = len(self.vals)
        self.vals.append(val)
        return True

    def remove(self, val: int) -> bool:
        if val not in self.idx:
            return False
        i, last = self.idx.pop(val), self.vals.pop()
        if i < len(self.vals):            # removed element wasn't the last one
            self.vals[i] = last
            self.idx[last] = i
        return True

    def getRandom(self) -> int:
        return random.choice(self.vals)
# test
rs = RandomizedSet()
assert rs.insert(1) and not rs.remove(2) and rs.insert(2)
assert rs.getRandom() in (1, 2)
assert rs.remove(1) and not rs.insert(2) and rs.getRandom() == 2
```

```java
class RandomizedSet {
    private final List<Integer> vals = new ArrayList<>();
    private final Map<Integer, Integer> idx = new HashMap<>();
    private final Random rnd = new Random();

    public boolean insert(int val) {
        if (idx.containsKey(val)) return false;
        idx.put(val, vals.size());
        vals.add(val);
        return true;
    }

    public boolean remove(int val) {
        Integer i = idx.remove(val);
        if (i == null) return false;
        int last = vals.remove(vals.size() - 1);
        if (i < vals.size()) { vals.set(i, last); idx.put(last, i); }
        return true;
    }

    public int getRandom() { return vals.get(rnd.nextInt(vals.size())); }
}
// test
RandomizedSet rs = new RandomizedSet();
check(rs.insert(1) && !rs.remove(2) && rs.insert(2), true);
check(rs.remove(1) && !rs.insert(2), true);
check(rs.getRandom(), 2);
```

## 232. Implement Queue using Stacks | Easy
related: 225, 155
tags: two-stacks

Push onto `in`. For `pop/peek`, if `out` is empty, pour all of `in` into `out` (reversing order). Each element moves at most once → amortized O(1).

**Complexity:** amortized O(1).

```python
class MyQueue:
    def __init__(self):
        self.inp, self.out = [], []

    def push(self, x: int) -> None:
        self.inp.append(x)

    def _move(self):
        if not self.out:
            while self.inp:
                self.out.append(self.inp.pop())

    def pop(self) -> int:
        self._move()
        return self.out.pop()

    def peek(self) -> int:
        self._move()
        return self.out[-1]

    def empty(self) -> bool:
        return not self.inp and not self.out
# test
q = MyQueue()
q.push(1); q.push(2)
assert q.peek() == 1 and q.pop() == 1 and not q.empty()
```

```java
class MyQueue {
    private final Deque<Integer> in = new ArrayDeque<>(), out = new ArrayDeque<>();

    public void push(int x) { in.push(x); }

    private void move() { if (out.isEmpty()) while (!in.isEmpty()) out.push(in.pop()); }

    public int pop() { move(); return out.pop(); }
    public int peek() { move(); return out.peek(); }
    public boolean empty() { return in.isEmpty() && out.isEmpty(); }
}
// test
MyQueue q = new MyQueue();
q.push(1); q.push(2);
check(q.peek(), 1);
check(q.pop(), 1);
check(q.empty(), false);
```

## 622. Design Circular Queue | Medium
related: 641
tags: ring-buffer

Fixed array + `head` index + `count`. Tail index = `(head + count − 1) % k`. No need to waste a slot when you track `count`.

**Complexity:** O(1) per operation.

```python
class MyCircularQueue:
    def __init__(self, k: int):
        self.a, self.head, self.n = [0] * k, 0, 0

    def enQueue(self, value: int) -> bool:
        if self.isFull():
            return False
        self.a[(self.head + self.n) % len(self.a)] = value
        self.n += 1
        return True

    def deQueue(self) -> bool:
        if self.isEmpty():
            return False
        self.head = (self.head + 1) % len(self.a)
        self.n -= 1
        return True

    def Front(self) -> int:
        return -1 if self.isEmpty() else self.a[self.head]

    def Rear(self) -> int:
        return -1 if self.isEmpty() else self.a[(self.head + self.n - 1) % len(self.a)]

    def isEmpty(self) -> bool:
        return self.n == 0

    def isFull(self) -> bool:
        return self.n == len(self.a)
# test
cq = MyCircularQueue(3)
assert cq.enQueue(1) and cq.enQueue(2) and cq.enQueue(3) and not cq.enQueue(4)
assert cq.Rear() == 3 and cq.isFull() and cq.deQueue() and cq.enQueue(4) and cq.Rear() == 4
```

```java
class MyCircularQueue {
    private final int[] a;
    private int head = 0, n = 0;

    public MyCircularQueue(int k) { a = new int[k]; }

    public boolean enQueue(int value) {
        if (isFull()) return false;
        a[(head + n++) % a.length] = value;
        return true;
    }

    public boolean deQueue() {
        if (isEmpty()) return false;
        head = (head + 1) % a.length;
        n--;
        return true;
    }

    public int Front() { return isEmpty() ? -1 : a[head]; }
    public int Rear() { return isEmpty() ? -1 : a[(head + n - 1) % a.length]; }
    public boolean isEmpty() { return n == 0; }
    public boolean isFull() { return n == a.length; }
}
// test
MyCircularQueue cq = new MyCircularQueue(3);
check(cq.enQueue(1) && cq.enQueue(2) && cq.enQueue(3) && !cq.enQueue(4), true);
check(cq.Rear(), 3);
check(cq.deQueue() && cq.enQueue(4), true);
check(cq.Rear(), 4);
check(cq.Front(), 2);
```

## 1146. Snapshot Array | Medium
related: 981
tags: versioning, binary-search

Don't copy the array on `snap`. Each index keeps a list of `(snap_id, value)` changes; `get(i, s)` = binary search for the last change with `snap_id ≤ s`.

**Complexity:** set O(1), snap O(1), get O(log changes).

```python
class SnapshotArray:
    def __init__(self, length: int):
        self.hist = [[(-1, 0)] for _ in range(length)]
        self.snap_id = 0

    def set(self, index: int, val: int) -> None:
        h = self.hist[index]
        if h[-1][0] == self.snap_id:
            h[-1] = (self.snap_id, val)
        else:
            h.append((self.snap_id, val))

    def snap(self) -> int:
        self.snap_id += 1
        return self.snap_id - 1

    def get(self, index: int, snap_id: int) -> int:
        h = self.hist[index]
        i = bisect.bisect_right(h, (snap_id, float("inf"))) - 1
        return h[i][1]
# test
sa = SnapshotArray(3)
sa.set(0, 5)
assert sa.snap() == 0
sa.set(0, 6)
assert sa.get(0, 0) == 5 and sa.get(0, 1) == 6 and sa.get(1, 0) == 0
```

```java
class SnapshotArray {
    private final List<TreeMap<Integer, Integer>> hist = new ArrayList<>();
    private int snapId = 0;

    public SnapshotArray(int length) {
        for (int i = 0; i < length; i++) {
            TreeMap<Integer, Integer> t = new TreeMap<>();
            t.put(-1, 0);
            hist.add(t);
        }
    }

    public void set(int index, int val) { hist.get(index).put(snapId, val); }
    public int snap() { return snapId++; }
    public int get(int index, int snap_id) { return hist.get(index).floorEntry(snap_id).getValue(); }
}
// test
SnapshotArray sa = new SnapshotArray(3);
sa.set(0, 5);
check(sa.snap(), 0);
sa.set(0, 6);
check(sa.get(0, 0), 5);
check(sa.get(1, 0), 0);
```

## 535. Encode and Decode TinyURL | Medium
related: 271
tags: design, base62

Assign each new long URL an incrementing id, encode the id in **base62** (`[0-9a-zA-Z]`) for a short code; keep both maps so repeated URLs reuse the same code. (See the URL shortener system design article for scaling this.)

**Complexity:** O(L) per call.

```python
class Codec:
    ALPHA = string.digits + string.ascii_letters

    def __init__(self):
        self.long_to_short, self.short_to_long = {}, {}

    def _base62(self, n):
        s = ""
        while True:
            n, r = divmod(n, 62)
            s = self.ALPHA[r] + s
            if n == 0:
                return s

    def encode(self, longUrl: str) -> str:
        if longUrl not in self.long_to_short:
            code = self._base62(len(self.short_to_long) + 1)
            self.long_to_short[longUrl] = code
            self.short_to_long[code] = longUrl
        return "http://tinyurl.com/" + self.long_to_short[longUrl]

    def decode(self, shortUrl: str) -> str:
        return self.short_to_long[shortUrl.rsplit("/", 1)[1]]
# test
c = Codec()
u = "https://leetcode.com/problems/design-tinyurl"
assert c.decode(c.encode(u)) == u and c.encode(u) == c.encode(u)
```

```java
class Codec {
    private static final String ALPHA = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
    private final Map<String, String> longToShort = new HashMap<>(), shortToLong = new HashMap<>();

    public String encode(String longUrl) {
        String code = longToShort.get(longUrl);
        if (code == null) {
            long n = shortToLong.size() + 1;
            StringBuilder sb = new StringBuilder();
            do { sb.append(ALPHA.charAt((int) (n % 62))); n /= 62; } while (n > 0);
            code = sb.reverse().toString();
            longToShort.put(longUrl, code);
            shortToLong.put(code, longUrl);
        }
        return "http://tinyurl.com/" + code;
    }

    public String decode(String shortUrl) {
        return shortToLong.get(shortUrl.substring(shortUrl.lastIndexOf('/') + 1));
    }
}
// test
Codec c = new Codec();
String u = "https://leetcode.com/problems/design-tinyurl";
check(c.decode(c.encode(u)), u);
```

## 706. Design HashMap | Easy
related: 705
tags: hashing, chaining

Array of buckets, each a list of `[key, value]` pairs (**separate chaining**). Index = `key % buckets`. Real implementations also resize when the load factor gets high.

**Complexity:** O(1) average, O(n/buckets) worst per bucket.

```python
class MyHashMap:
    def __init__(self):
        self.size = 1009                       # prime number of buckets
        self.b = [[] for _ in range(self.size)]

    def put(self, key: int, value: int) -> None:
        bucket = self.b[key % self.size]
        for pair in bucket:
            if pair[0] == key:
                pair[1] = value
                return
        bucket.append([key, value])

    def get(self, key: int) -> int:
        for k, v in self.b[key % self.size]:
            if k == key:
                return v
        return -1

    def remove(self, key: int) -> None:
        bucket = self.b[key % self.size]
        for i, (k, _) in enumerate(bucket):
            if k == key:
                bucket.pop(i)
                return
# test
hm = MyHashMap()
hm.put(1, 1); hm.put(2, 2)
assert hm.get(1) == 1 and hm.get(3) == -1
hm.put(2, 1); assert hm.get(2) == 1
hm.remove(2); assert hm.get(2) == -1
hm.put(1009 + 1, 7); assert hm.get(1) == 1 and hm.get(1010) == 7   # same bucket
```

```java
class MyHashMap {
    private static final int SIZE = 1009;
    private final List<int[]>[] b;

    @SuppressWarnings("unchecked")
    public MyHashMap() {
        b = new List[SIZE];
        for (int i = 0; i < SIZE; i++) b[i] = new LinkedList<>();
    }

    public void put(int key, int value) {
        for (int[] p : b[key % SIZE]) if (p[0] == key) { p[1] = value; return; }
        b[key % SIZE].add(new int[]{key, value});
    }

    public int get(int key) {
        for (int[] p : b[key % SIZE]) if (p[0] == key) return p[1];
        return -1;
    }

    public void remove(int key) { b[key % SIZE].removeIf(p -> p[0] == key); }
}
// test
MyHashMap hm = new MyHashMap();
hm.put(1, 1); hm.put(2, 2); hm.put(1010, 7);
check(hm.get(1), 1);
check(hm.get(1010), 7);
hm.remove(2);
check(hm.get(2), -1);
```
