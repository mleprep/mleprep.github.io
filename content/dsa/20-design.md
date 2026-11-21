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
tags: hash-map, linked-list

```problem
Design `LRUCache(capacity)`: `get(key)` returns the value or −1 if absent; `put(key, value)` inserts or updates a key, and if that pushes the size over `capacity`, it evicts the least recently used key. Both must run in O(1) average time.

**Example:** capacity 2: put(1,1), put(2,2), get(1) → 1, put(3,3) evicts key 2, get(2) → −1

**Constraints:** up to 2·10⁵ calls.
```

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

```go
type lruNode struct {
	key, val   int
	prev, next *lruNode
}

type LRUCache struct {
	cap        int
	nodes      map[int]*lruNode
	head, tail *lruNode // sentinels: head.next is the most recent, tail.prev the least
}

func Constructor(capacity int) LRUCache {
	head, tail := &lruNode{}, &lruNode{}
	head.next, tail.prev = tail, head
	return LRUCache{cap: capacity, nodes: map[int]*lruNode{}, head: head, tail: tail}
}

func (this *LRUCache) unlink(n *lruNode) {
	n.prev.next, n.next.prev = n.next, n.prev
}

func (this *LRUCache) pushFront(n *lruNode) {
	n.prev, n.next = this.head, this.head.next
	this.head.next.prev = n
	this.head.next = n
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.nodes[key]
	if !ok {
		return -1
	}
	this.unlink(n) // used: move to the front
	this.pushFront(n)
	return n.val
}

func (this *LRUCache) Put(key int, value int) {
	if n, ok := this.nodes[key]; ok {
		n.val = value
		this.unlink(n)
		this.pushFront(n)
		return
	}
	n := &lruNode{key: key, val: value}
	this.nodes[key] = n
	this.pushFront(n)
	if len(this.nodes) > this.cap { // evict the least recently used
		lru := this.tail.prev
		this.unlink(lru)
		delete(this.nodes, lru.key)
	}
}
// test
c := Constructor(2)
c.Put(1, 1)
c.Put(2, 2)
assertEq(c.Get(1), 1)
c.Put(3, 3) // evicts 2
assertEq(c.Get(2), -1)
c.Put(4, 4) // evicts 1
assertEq([]int{c.Get(1), c.Get(3), c.Get(4)}, []int{-1, 3, 4})
```

## 460. LFU Cache | Hard
related: 146
tags: hash-map, bucket-sort

```problem
Design `LFUCache(capacity)` with `get(key)` and `put(key, value)` in O(1) average time. When full, `put` evicts the least frequently used key (counting both gets and puts), breaking ties by evicting the least recently used among them.

**Example:** capacity 2: put(1,1), put(2,2), get(1) → 1, put(3,3) evicts key 2 (used once), get(2) → −1
```

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

```go
type lfuEntry struct {
	key, val, freq int
}

type LFUCache struct {
	cap, minFreq int
	items        map[int]*list.Element // key -> its element in lists[freq]
	lists        map[int]*list.List    // freq -> entries, most recently used at the front
}

func Constructor(capacity int) LFUCache {
	return LFUCache{cap: capacity, items: map[int]*list.Element{}, lists: map[int]*list.List{}}
}

func (this *LFUCache) touch(el *list.Element) { // one more use: move to the next frequency list
	e := el.Value.(*lfuEntry)
	old := this.lists[e.freq]
	old.Remove(el)
	if old.Len() == 0 {
		delete(this.lists, e.freq)
		if this.minFreq == e.freq {
			this.minFreq++
		}
	}
	e.freq++
	if this.lists[e.freq] == nil {
		this.lists[e.freq] = list.New()
	}
	this.items[e.key] = this.lists[e.freq].PushFront(e)
}

func (this *LFUCache) Get(key int) int {
	el, ok := this.items[key]
	if !ok {
		return -1
	}
	e := el.Value.(*lfuEntry)
	this.touch(el)
	return e.val
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}
	if el, ok := this.items[key]; ok {
		e := el.Value.(*lfuEntry)
		e.val = value
		this.touch(el)
		return
	}
	if len(this.items) == this.cap { // evict the LRU entry among the least frequent
		l := this.lists[this.minFreq]
		victim := l.Back()
		l.Remove(victim)
		if l.Len() == 0 {
			delete(this.lists, this.minFreq)
		}
		delete(this.items, victim.Value.(*lfuEntry).key)
	}
	if this.lists[1] == nil {
		this.lists[1] = list.New()
	}
	this.items[key] = this.lists[1].PushFront(&lfuEntry{key: key, val: value, freq: 1})
	this.minFreq = 1
}
// test
c := Constructor(2)
c.Put(1, 1)
c.Put(2, 2)
assertEq(c.Get(1), 1)
c.Put(3, 3) // evicts 2 (freq 1)
assertEq([]int{c.Get(2), c.Get(3)}, []int{-1, 3})
c.Put(4, 4) // 1 and 3 both have freq 2: evict the least recent of them, 1
assertEq([]int{c.Get(1), c.Get(3), c.Get(4)}, []int{-1, 3, 4})
```

## 380. Insert Delete GetRandom O(1) | Medium
related: 381
tags: array-plus-map

```problem
Design `RandomizedSet` with `insert(val)` and `remove(val)` (each returns whether the set changed) and `getRandom()` (a uniformly random current element, never called on an empty set), all in O(1) average time.

**Example:** insert(1) → true, remove(2) → false, insert(2) → true, getRandom() → 1 or 2, remove(1) → true, getRandom() → 2
```

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

```go
type RandomizedSet struct {
	vals []int
	idx  map[int]int // value -> position in vals
}

func Constructor() RandomizedSet {
	return RandomizedSet{idx: map[int]int{}}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.idx[val]; ok {
		return false
	}
	this.idx[val] = len(this.vals)
	this.vals = append(this.vals, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	i, ok := this.idx[val]
	if !ok {
		return false
	}
	last := this.vals[len(this.vals)-1] // move the last value into the hole: O(1)
	this.vals[i] = last
	this.idx[last] = i
	this.vals = this.vals[:len(this.vals)-1]
	delete(this.idx, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.vals[rand.Intn(len(this.vals))]
}
// test
rs := Constructor()
assertEq([]bool{rs.Insert(1), rs.Remove(2), rs.Insert(2)}, []bool{true, false, true})
r := rs.GetRandom()
assertEq(r == 1 || r == 2, true)
assertEq([]bool{rs.Remove(1), rs.Insert(2)}, []bool{true, false})
assertEq(rs.GetRandom(), 2)
```

## 232. Implement Queue using Stacks | Easy
related: 225, 155
tags: two-stacks

```problem
Implement a FIFO queue `MyQueue` (`push`, `pop`, `peek`, `empty`) using only two stacks and standard stack operations.

**Example:** push 1, push 2 → peek() = 1, pop() = 1, empty() = false

**Constraints:** follow-up: amortised O(1) per operation.
```

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
q.push(3)
assert q.pop() == 2 and q.peek() == 3 and q.pop() == 3 and q.empty()
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

```go
type MyQueue struct {
	in, out []int // push onto in; pop from out, refilled (reversed) only when empty
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) move() {
	if len(this.out) == 0 {
		for len(this.in) > 0 {
			this.out = append(this.out, this.in[len(this.in)-1])
			this.in = this.in[:len(this.in)-1]
		}
	}
}

func (this *MyQueue) Pop() int {
	this.move()
	x := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return x
}

func (this *MyQueue) Peek() int {
	this.move()
	return this.out[len(this.out)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}
// test
q := Constructor()
q.Push(1)
q.Push(2)
assertEq([]any{q.Peek(), q.Pop(), q.Empty()}, []any{1, 1, false})
q.Push(3)
assertEq([]any{q.Pop(), q.Peek(), q.Pop(), q.Empty()}, []any{2, 3, 3, true})
```

## 622. Design Circular Queue | Medium
related: 641
tags: ring-buffer

```problem
Design `MyCircularQueue(k)`, a fixed-size ring buffer with `enQueue(value)` and `deQueue()` (each returns whether it succeeded), `Front()` and `Rear()` (−1 when empty), `isEmpty()` and `isFull()`. Don't use a built-in queue.

**Example:** k = 3: enQueue 1, 2, 3 → true, enQueue 4 → false, Rear() = 3, deQueue() → true, enQueue 4 → true, Rear() = 4
```

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

```go
type MyCircularQueue struct {
	a       []int
	head, n int // index of the front, number of items
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{a: make([]int, k)}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.a[(this.head+this.n)%len(this.a)] = value
	this.n++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % len(this.a)
	this.n--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.a[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.a[(this.head+this.n-1)%len(this.a)]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.n == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.n == len(this.a)
}
// test
cq := Constructor(3)
assertEq([]bool{cq.EnQueue(1), cq.EnQueue(2), cq.EnQueue(3), cq.EnQueue(4)}, []bool{true, true, true, false})
assertEq([]any{cq.Rear(), cq.IsFull(), cq.DeQueue(), cq.EnQueue(4), cq.Rear(), cq.Front()}, []any{3, true, true, true, 4, 2})
```

## 1146. Snapshot Array | Medium
related: 981
tags: versioning, binary-search

```problem
Design `SnapshotArray(length)`, initially all zeros, with `set(index, val)`, `snap()` (takes a snapshot and returns its id: the number of snaps before it), and `get(index, snap_id)` (the value at `index` when snapshot `snap_id` was taken).

**Example:** `SnapshotArray(3)`: set(0,5), snap() → 0, set(0,6), get(0,0) → 5

**Constraints:** length and number of calls ≤ 5·10⁴.
```

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

```go
type SnapshotArray struct {
	hist   [][][2]int // per index: (snap id, value) changes, in snap order
	snapID int
}

func Constructor(length int) SnapshotArray {
	hist := make([][][2]int, length)
	for i := range hist {
		hist[i] = [][2]int{{-1, 0}} // everything starts at 0
	}
	return SnapshotArray{hist: hist}
}

func (this *SnapshotArray) Set(index int, val int) {
	h := this.hist[index]
	if h[len(h)-1][0] == this.snapID { // overwrite within the same snapshot
		h[len(h)-1][1] = val
	} else {
		this.hist[index] = append(h, [2]int{this.snapID, val})
	}
}

func (this *SnapshotArray) Snap() int {
	this.snapID++
	return this.snapID - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	h := this.hist[index]
	i := sort.Search(len(h), func(k int) bool { return h[k][0] > snap_id }) - 1 // last change at or before snap_id
	return h[i][1]
}
// test
sa := Constructor(3)
sa.Set(0, 5)
assertEq(sa.Snap(), 0)
sa.Set(0, 6)
assertEq([]int{sa.Get(0, 0), sa.Get(0, 1), sa.Get(1, 0)}, []int{5, 6, 0})
```

## 535. Encode and Decode TinyURL | Medium
related: 271
tags: design, base62

```problem
Design a URL shortener `Codec`: `encode(longUrl)` returns a short URL, and `decode(shortUrl)` returns the original long URL. Any scheme works as long as decoding an encoded URL gives back the original.

**Example:** `decode(encode("https://example.com/some/long/path"))` → `"https://example.com/some/long/path"`
```

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
for u in ("https://example.com/a?b=c#d", "x"):
    assert c.decode(c.encode(u)) == u
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

```go
type Codec struct {
	longToShort, shortToLong map[string]string
}

func Constructor() Codec {
	return Codec{longToShort: map[string]string{}, shortToLong: map[string]string{}}
}

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func base62(n int) string {
	s := ""
	for {
		s = string(alphabet[n%62]) + s
		n /= 62
		if n == 0 {
			return s
		}
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	code, ok := this.longToShort[longUrl]
	if !ok { // a counter in base 62: short, unique, no collisions
		code = base62(len(this.shortToLong) + 1)
		this.longToShort[longUrl] = code
		this.shortToLong[code] = longUrl
	}
	return "http://tinyurl.com/" + code
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.shortToLong[shortUrl[strings.LastIndex(shortUrl, "/")+1:]]
}
// test
c := Constructor()
u := "https://leetcode.com/problems/design-tinyurl"
assertEq(c.decode(c.encode(u)), u)
assertEq(c.encode(u) == c.encode(u), true)
for _, u := range []string{"https://example.com/a?b=c#d", "x"} {
	assertEq(c.decode(c.encode(u)), u)
}
```

## 706. Design HashMap | Easy
related: 705
tags: hashing, chaining

```problem
Design `MyHashMap` without built-in hash tables: `put(key, value)` inserts or updates, `get(key)` returns the value or −1, and `remove(key)` deletes the key if present.

**Example:** put(1,1), put(2,2), get(1) → 1, get(3) → −1, put(2,1), get(2) → 1, remove(2), get(2) → −1

**Constraints:** keys and values ≤ 10⁶; up to 10⁴ calls.
```

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

```go
type MyHashMap struct {
	b [][][2]int // separate chaining: buckets of (key, value)
}

const buckets = 1009 // a prime number of buckets

func Constructor() MyHashMap {
	return MyHashMap{b: make([][][2]int, buckets)}
}

func (this *MyHashMap) Put(key int, value int) {
	b := this.b[key%buckets]
	for i := range b {
		if b[i][0] == key {
			b[i][1] = value
			return
		}
	}
	this.b[key%buckets] = append(b, [2]int{key, value})
}

func (this *MyHashMap) Get(key int) int {
	for _, kv := range this.b[key%buckets] {
		if kv[0] == key {
			return kv[1]
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	b := this.b[key%buckets]
	for i := range b {
		if b[i][0] == key {
			this.b[key%buckets] = append(b[:i], b[i+1:]...)
			return
		}
	}
}
// test
hm := Constructor()
hm.Put(1, 1)
hm.Put(2, 2)
assertEq([]int{hm.Get(1), hm.Get(3)}, []int{1, -1})
hm.Put(2, 1)
assertEq(hm.Get(2), 1)
hm.Remove(2)
assertEq(hm.Get(2), -1)
hm.Put(1009+1, 7) // same bucket as 1
assertEq([]int{hm.Get(1), hm.Get(1010)}, []int{1, 7})
```


## 1472. Design Browser History | Medium
related: 155
tags: array, pointer

```problem
Design `BrowserHistory(homepage)` for a single tab: `visit(url)` opens url and clears all forward history, `back(steps)` and `forward(steps)` move at most `steps` pages (stopping at either end) and return the current URL.

**Example:** home "a.com", visit b, c, d → back(1) = "c", back(1) = "b", forward(1) = "c", visit e → forward(2) = "e", back(2) = "b", back(7) = "a.com"
```

Array of pages + `cur` index + `last` (the end of the valid forward history). `visit` overwrites at `cur+1` and truncates forward history by setting `last = cur`.

**Complexity:** O(1) per operation.

```python
class BrowserHistory:
    def __init__(self, homepage: str):
        self.pages, self.cur, self.last = [homepage], 0, 0

    def visit(self, url: str) -> None:
        self.cur += 1
        if self.cur == len(self.pages):
            self.pages.append(url)
        else:
            self.pages[self.cur] = url
        self.last = self.cur

    def back(self, steps: int) -> str:
        self.cur = max(0, self.cur - steps)
        return self.pages[self.cur]

    def forward(self, steps: int) -> str:
        self.cur = min(self.last, self.cur + steps)
        return self.pages[self.cur]
# test
b = BrowserHistory("leetcode.com")
for u in ("google.com", "facebook.com", "youtube.com"):
    b.visit(u)
assert b.back(1) == "facebook.com" and b.back(1) == "google.com" and b.forward(1) == "facebook.com"
b.visit("linkedin.com")
assert b.forward(2) == "linkedin.com" and b.back(2) == "google.com" and b.back(7) == "leetcode.com"
```

```java
class BrowserHistory {
    private final List<String> pages = new ArrayList<>();
    private int cur = 0, last = 0;

    public BrowserHistory(String homepage) { pages.add(homepage); }

    public void visit(String url) {
        cur++;
        if (cur == pages.size()) pages.add(url); else pages.set(cur, url);
        last = cur;
    }

    public String back(int steps) { cur = Math.max(0, cur - steps); return pages.get(cur); }
    public String forward(int steps) { cur = Math.min(last, cur + steps); return pages.get(cur); }
}
// test
BrowserHistory b = new BrowserHistory("leetcode.com");
b.visit("google.com"); b.visit("facebook.com"); b.visit("youtube.com");
check(b.back(2), "google.com");
b.visit("linkedin.com");
check(b.forward(2), "linkedin.com");
check(b.back(7), "leetcode.com");
```

```go
type BrowserHistory struct {
	pages     []string
	cur, last int // current page; last valid page (later ones are stale)
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{pages: []string{homepage}}
}

func (this *BrowserHistory) Visit(url string) {
	this.cur++
	if this.cur == len(this.pages) {
		this.pages = append(this.pages, url)
	} else {
		this.pages[this.cur] = url // overwrite instead of truncating
	}
	this.last = this.cur // forward history is gone
}

func (this *BrowserHistory) Back(steps int) string {
	this.cur = max(0, this.cur-steps)
	return this.pages[this.cur]
}

func (this *BrowserHistory) Forward(steps int) string {
	this.cur = min(this.last, this.cur+steps)
	return this.pages[this.cur]
}
// test
b := Constructor("leetcode.com")
for _, u := range []string{"google.com", "facebook.com", "youtube.com"} {
	b.Visit(u)
}
assertEq([]string{b.Back(1), b.Back(1), b.Forward(1)}, []string{"facebook.com", "google.com", "facebook.com"})
b.Visit("linkedin.com")
assertEq([]string{b.Forward(2), b.Back(2), b.Back(7)}, []string{"linkedin.com", "google.com", "leetcode.com"})
```

## 1396. Design Underground System | Medium
related: 146
tags: hash-map

```problem
Design `UndergroundSystem`: `checkIn(id, station, t)` and `checkOut(id, station, t)` record a traveller's trip, and `getAverageTime(start, end)` returns the average duration of all completed trips that went directly from `start` to `end` (at least one exists). A traveller is checked into at most one station at a time.

**Example:** trips Leyton → Waterloo taking 12 and 10 → `getAverageTime("Leyton","Waterloo")` = 11.0
```

`checkIn` stores `id → (station, time)`; `checkOut` pops it and accumulates `(start, end) → [total time, trips]`.

**Complexity:** O(1) per operation.

```python
class UndergroundSystem:
    def __init__(self):
        self.active, self.stats = {}, defaultdict(lambda: [0, 0])

    def checkIn(self, id: int, stationName: str, t: int) -> None:
        self.active[id] = (stationName, t)

    def checkOut(self, id: int, stationName: str, t: int) -> None:
        start, t0 = self.active.pop(id)
        s = self.stats[(start, stationName)]
        s[0] += t - t0
        s[1] += 1

    def getAverageTime(self, startStation: str, endStation: str) -> float:
        total, n = self.stats[(startStation, endStation)]
        return total / n
# test
u = UndergroundSystem()
u.checkIn(45, "Leyton", 3); u.checkIn(32, "Paradise", 8); u.checkIn(27, "Leyton", 10)
u.checkOut(45, "Waterloo", 15); u.checkOut(27, "Waterloo", 20); u.checkOut(32, "Cambridge", 22)
assert u.getAverageTime("Paradise", "Cambridge") == 14.0 and u.getAverageTime("Leyton", "Waterloo") == 11.0
u2 = UndergroundSystem()
u2.checkIn(10, "A", 1); u2.checkOut(10, "B", 5); u2.checkIn(11, "A", 2); u2.checkOut(11, "B", 4)
assert u2.getAverageTime("A", "B") == 3.0
```

```java
class UndergroundSystem {
    private final Map<Integer, Object[]> active = new HashMap<>();   // id -> {station, time}
    private final Map<String, long[]> stats = new HashMap<>();       // "a->b" -> {total, trips}

    public void checkIn(int id, String stationName, int t) { active.put(id, new Object[]{stationName, t}); }

    public void checkOut(int id, String stationName, int t) {
        Object[] in = active.remove(id);
        long[] s = stats.computeIfAbsent(in[0] + "->" + stationName, k -> new long[2]);
        s[0] += t - (int) in[1];
        s[1]++;
    }

    public double getAverageTime(String startStation, String endStation) {
        long[] s = stats.get(startStation + "->" + endStation);
        return (double) s[0] / s[1];
    }
}
// test
UndergroundSystem u = new UndergroundSystem();
u.checkIn(45, "Leyton", 3); u.checkIn(27, "Leyton", 10);
u.checkOut(45, "Waterloo", 15); u.checkOut(27, "Waterloo", 20);
check(u.getAverageTime("Leyton", "Waterloo"), 11.0);
```

```go
type checkIn struct {
	station string
	t       int
}

type UndergroundSystem struct {
	active map[int]checkIn      // id -> where and when they entered
	stats  map[[2]string][2]int // (from, to) -> (total time, trips)
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{active: map[int]checkIn{}, stats: map[[2]string][2]int{}}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.active[id] = checkIn{stationName, t}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	in := this.active[id]
	delete(this.active, id)
	key := [2]string{in.station, stationName}
	s := this.stats[key]
	s[0] += t - in.t
	s[1]++
	this.stats[key] = s
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	s := this.stats[[2]string{startStation, endStation}]
	return float64(s[0]) / float64(s[1])
}
// test
u := Constructor()
u.CheckIn(45, "Leyton", 3)
u.CheckIn(32, "Paradise", 8)
u.CheckIn(27, "Leyton", 10)
u.CheckOut(45, "Waterloo", 15)
u.CheckOut(27, "Waterloo", 20)
u.CheckOut(32, "Cambridge", 22)
assertEq([]float64{u.GetAverageTime("Paradise", "Cambridge"), u.GetAverageTime("Leyton", "Waterloo")}, []float64{14, 11})
u2 := Constructor()
u2.CheckIn(10, "A", 1)
u2.CheckOut(10, "B", 5)
u2.CheckIn(11, "A", 2)
u2.CheckOut(11, "B", 4)
assertEq(u2.GetAverageTime("A", "B"), 3.0)
```

## 1352. Product of the Last K Numbers | Medium
related: 238, 303
tags: prefix-product

```problem
Design `ProductOfNumbers` for a stream of non-negative integers: `add(num)` appends a number and `getProduct(k)` returns the product of the last `k` numbers (k never exceeds the count so far; products fit in 32 bits).

**Example:** add 3, 0, 2, 5, 4 → getProduct(2) = 20, getProduct(3) = 40, getProduct(4) = 0; add 8 → getProduct(2) = 32

**Constraints:** follow-up: O(1) per call.
```

Keep prefix products; a 0 resets the list (any window containing it is 0). `getProduct(k)` = `pre[-1] / pre[-1-k]`, or 0 if k reaches past the last reset.

**Complexity:** O(1) per operation.

```python
class ProductOfNumbers:
    def __init__(self):
        self.pre = [1]

    def add(self, num: int) -> None:
        if num == 0:
            self.pre = [1]
        else:
            self.pre.append(self.pre[-1] * num)

    def getProduct(self, k: int) -> int:
        if k >= len(self.pre):
            return 0
        return self.pre[-1] // self.pre[-1 - k]
# test
p = ProductOfNumbers()
for x in (3, 0, 2, 5, 4):
    p.add(x)
assert [p.getProduct(2), p.getProduct(3), p.getProduct(4)] == [20, 40, 0]
p.add(8)
assert p.getProduct(2) == 32
```

```java
class ProductOfNumbers {
    private final List<Integer> pre = new ArrayList<>(List.of(1));

    public void add(int num) {
        if (num == 0) { pre.clear(); pre.add(1); }
        else pre.add(pre.get(pre.size() - 1) * num);
    }

    public int getProduct(int k) {
        int n = pre.size();
        return k >= n ? 0 : pre.get(n - 1) / pre.get(n - 1 - k);
    }
}
// test
ProductOfNumbers p = new ProductOfNumbers();
for (int x : new int[]{3, 0, 2, 5, 4}) p.add(x);
check(p.getProduct(3), 40);
check(p.getProduct(4), 0);
```

```go
type ProductOfNumbers struct {
	pre []int // prefix products since the last zero
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{pre: []int{1}}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.pre = []int{1} // any window containing the zero has product 0
	} else {
		this.pre = append(this.pre, this.pre[len(this.pre)-1]*num)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	n := len(this.pre)
	if k >= n {
		return 0 // the window reaches back past a zero
	}
	return this.pre[n-1] / this.pre[n-1-k]
}
// test
p := Constructor()
for _, x := range []int{3, 0, 2, 5, 4} {
	p.Add(x)
}
assertEq([]int{p.GetProduct(2), p.GetProduct(3), p.GetProduct(4)}, []int{20, 40, 0})
p.Add(8)
assertEq(p.GetProduct(2), 32)
```

## 284. Peeking Iterator | Medium
related: 341, 173
tags: iterator, design

```problem
Wrap an iterator (with `next()` and `hasNext()`) in a `PeekingIterator` that also supports `peek()`: return the next element without advancing.

**Example:** over `[1,2,3]`: next() = 1, peek() = 2, next() = 2, next() = 3, hasNext() = false

**Constraints:** follow-up: make it generic over element types.
```

Cache the next element: `peek` returns the cache; `next` returns it and refills from the underlying iterator.

**Complexity:** O(1) per call.

```python
class PeekingIterator:
    def __init__(self, iterator):
        self.it = iterator
        self.nxt = iterator.next() if iterator.hasNext() else None
        self.has = self.nxt is not None or iterator.hasNext()

    def peek(self):
        return self.nxt

    def next(self):
        val = self.nxt
        self.has = self.it.hasNext()
        self.nxt = self.it.next() if self.has else None
        return val

    def hasNext(self):
        return self.has
# test
class Iterator:   # LeetCode's interface, minimal stand-in
    def __init__(self, nums): self.nums, self.i = nums, 0
    def hasNext(self): return self.i < len(self.nums)
    def next(self): self.i += 1; return self.nums[self.i - 1]
pi = PeekingIterator(Iterator([1, 2, 3]))
assert [pi.next(), pi.peek(), pi.next(), pi.next(), pi.hasNext()] == [1, 2, 2, 3, False]
```

```java
class PeekingIterator implements Iterator<Integer> {
    private final Iterator<Integer> it;
    private Integer nxt;

    public PeekingIterator(Iterator<Integer> iterator) {
        it = iterator;
        nxt = it.hasNext() ? it.next() : null;
    }

    public Integer peek() { return nxt; }

    @Override
    public Integer next() {
        Integer val = nxt;
        nxt = it.hasNext() ? it.next() : null;
        return val;
    }

    @Override
    public boolean hasNext() { return nxt != null; }
}
// test
PeekingIterator pi = new PeekingIterator(List.of(1, 2, 3).iterator());
check(pi.next(), 1);
check(pi.peek(), 2);
check(pi.next(), 2);
check(pi.next(), 3);
check(pi.hasNext(), false);
```

```go
type Iterator struct { // provided by LeetCode (a minimal stand-in)
	nums []int
	i    int
}

func (this *Iterator) hasNext() bool {
	return this.i < len(this.nums)
}

func (this *Iterator) next() int {
	this.i++
	return this.nums[this.i-1]
}

type PeekingIterator struct {
	iter    *Iterator
	has     bool
	nextVal int // cached next element
}

func Constructor(iter *Iterator) *PeekingIterator {
	p := &PeekingIterator{iter: iter}
	p.advance()
	return p
}

func (this *PeekingIterator) advance() {
	this.has = this.iter.hasNext()
	if this.has {
		this.nextVal = this.iter.next()
	}
}

func (this *PeekingIterator) hasNext() bool {
	return this.has
}

func (this *PeekingIterator) next() int {
	v := this.nextVal
	this.advance()
	return v
}

func (this *PeekingIterator) peek() int {
	return this.nextVal
}
// test
pi := Constructor(&Iterator{nums: []int{1, 2, 3}})
assertEq([]any{pi.next(), pi.peek(), pi.next(), pi.next(), pi.hasNext()}, []any{1, 2, 2, 3, false})
```

## 341. Flatten Nested List Iterator | Medium
related: 284, 385
tags: stack, iterator

```problem
Design `NestedIterator` over a nested list in which each element is an integer or another nested list. `next()` returns the integers one by one in left-to-right order, flattening the nesting, and `hasNext()` says whether any remain.

**Example:** `[[1,1],2,[1,1]]` → `1,1,2,1,1`; `[1,[4,[6]]]` → `1,4,6`
```

Keep a stack of lists in **reverse** order. `hasNext` does the real work: while the top is a list, pop it and push its children reversed; stop when the top is an integer.

**Complexity:** amortized O(1) per element.

```python
class NestedInteger:  # LeetCode's interface (minimal stand-in)
    def __init__(self, v): self.v = v
    def isInteger(self): return isinstance(self.v, int)
    def getInteger(self): return self.v
    def getList(self): return self.v

class NestedIterator:
    def __init__(self, nestedList: List["NestedInteger"]):
        self.st = nestedList[::-1]

    def next(self) -> int:
        return self.st.pop().getInteger()      # hasNext() guarantees an integer on top

    def hasNext(self) -> bool:
        while self.st and not self.st[-1].isInteger():
            self.st.extend(reversed(self.st.pop().getList()))
        return bool(self.st)
# test
def wrap(x): return NestedInteger(x if isinstance(x, int) else [wrap(y) for y in x])
it, out = NestedIterator([wrap(x) for x in [[1, 1], 2, [1, 1]]]), []
while it.hasNext():
    out.append(it.next())
assert out == [1, 1, 2, 1, 1]
it, out = NestedIterator([wrap(x) for x in [1, [4, [6]], [[]]]]), []
while it.hasNext():
    out.append(it.next())
assert out == [1, 4, 6]
```

```java
interface NestedInteger { // provided by LeetCode
    boolean isInteger();
    Integer getInteger();
    List<NestedInteger> getList();
}

class NI implements NestedInteger { // local stand-in for tests
    private final Integer v; private final List<NestedInteger> l;
    NI(int v) { this.v = v; this.l = null; }
    NI(NestedInteger... xs) { this.v = null; this.l = List.of(xs); }
    public boolean isInteger() { return v != null; }
    public Integer getInteger() { return v; }
    public List<NestedInteger> getList() { return l; }
}

class NestedIterator implements Iterator<Integer> {
    private final Deque<NestedInteger> st = new ArrayDeque<>();

    public NestedIterator(List<NestedInteger> nestedList) {
        for (int i = nestedList.size() - 1; i >= 0; i--) st.push(nestedList.get(i));
    }

    @Override
    public Integer next() { hasNext(); return st.pop().getInteger(); }

    @Override
    public boolean hasNext() {
        while (!st.isEmpty() && !st.peek().isInteger()) {
            List<NestedInteger> l = st.pop().getList();
            for (int i = l.size() - 1; i >= 0; i--) st.push(l.get(i));
        }
        return !st.isEmpty();
    }
}
// test
NestedIterator it = new NestedIterator(List.of(new NI(new NI(1), new NI(1)), new NI(2), new NI(new NI(1), new NI(1))));
List<Integer> out = new ArrayList<>();
while (it.hasNext()) out.add(it.next());
check(out, "[1, 1, 2, 1, 1]");
```

```go
type NestedInteger struct { // provided by LeetCode (a minimal stand-in)
	isInt bool
	val   int
	list  []*NestedInteger
}

func (this NestedInteger) IsInteger() bool { return this.isInt }

func (this NestedInteger) GetInteger() int { return this.val }

func (this NestedInteger) GetList() []*NestedInteger { return this.list }

type NestedIterator struct {
	st []*NestedInteger // reversed: the next element is on top
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	st := make([]*NestedInteger, len(nestedList))
	for i, x := range nestedList {
		st[len(nestedList)-1-i] = x
	}
	return &NestedIterator{st: st}
}

func (this *NestedIterator) Next() int {
	top := this.st[len(this.st)-1] // HasNext() left an integer on top
	this.st = this.st[:len(this.st)-1]
	return top.GetInteger()
}

func (this *NestedIterator) HasNext() bool {
	for len(this.st) > 0 && !this.st[len(this.st)-1].IsInteger() { // unpack lists lazily
		top := this.st[len(this.st)-1]
		this.st = this.st[:len(this.st)-1]
		l := top.GetList()
		for i := len(l) - 1; i >= 0; i-- {
			this.st = append(this.st, l[i])
		}
	}
	return len(this.st) > 0
}
// test
var wrap func(x any) *NestedInteger
wrap = func(x any) *NestedInteger {
	if v, ok := x.(int); ok {
		return &NestedInteger{isInt: true, val: v}
	}
	n := &NestedInteger{}
	for _, y := range x.([]any) {
		n.list = append(n.list, wrap(y))
	}
	return n
}
flatten := func(items []any) []int {
	nested := []*NestedInteger{}
	for _, x := range items {
		nested = append(nested, wrap(x))
	}
	it, out := Constructor(nested), []int{}
	for it.HasNext() {
		out = append(out, it.Next())
	}
	return out
}
assertEq(flatten([]any{[]any{1, 1}, 2, []any{1, 1}}), []int{1, 1, 2, 1, 1})
assertEq(flatten([]any{1, []any{4, []any{6}}, []any{[]any{}}}), []int{1, 4, 6})
```

## 432. All O`one Data Structure | Hard
related: 460, 146
tags: linked-list, bucket-sort

```problem
Design `AllOne`, which stores string counts: `inc(key)` adds 1 (inserting the key with count 1 if new), `dec(key)` subtracts 1 (removing the key at 0; the key always exists), `getMaxKey()` and `getMinKey()` return a key with the highest or lowest count (`""` if empty). Every operation must run in O(1) average time.

**Example:** inc("hello") ×2 → getMaxKey() = "hello"; inc("leet") → getMaxKey() = "hello", getMinKey() = "leet"
```

Like LFU: a doubly linked list of **count buckets** in increasing order, each holding a set of keys, plus `key → bucket`. `inc`/`dec` move a key to the neighbouring bucket (creating it if the count isn't there), deleting emptied buckets. Min/max = first/last bucket.

```viz
head <-> [1: {b}] <-> [2: {a, c}] <-> [5: {d}] <-> tail
getMinKey -> b      getMaxKey -> d      inc(c) -> c moves to new bucket [3]
```

**Complexity:** O(1) per operation.

```python
class Bucket:
    def __init__(self, cnt):
        self.cnt, self.keys, self.prev, self.next = cnt, set(), None, None

class AllOne:
    def __init__(self):
        self.head, self.tail = Bucket(0), Bucket(0)
        self.head.next, self.tail.prev = self.tail, self.head
        self.where = {}

    def _after(self, node, cnt):
        b = Bucket(cnt)
        b.prev, b.next = node, node.next
        node.next.prev = b
        node.next = b
        return b

    def _drop_if_empty(self, b):
        if not b.keys:
            b.prev.next, b.next.prev = b.next, b.prev

    def inc(self, key: str) -> None:
        cur = self.where.get(key, self.head)
        nxt = cur.next
        if nxt is self.tail or nxt.cnt != cur.cnt + 1:
            nxt = self._after(cur, cur.cnt + 1)
        nxt.keys.add(key)
        self.where[key] = nxt
        if cur is not self.head:
            cur.keys.discard(key)
            self._drop_if_empty(cur)

    def dec(self, key: str) -> None:
        cur = self.where[key]
        if cur.cnt == 1:
            del self.where[key]
        else:
            prv = cur.prev
            if prv is self.head or prv.cnt != cur.cnt - 1:
                prv = self._after(cur.prev, cur.cnt - 1)
            prv.keys.add(key)
            self.where[key] = prv
        cur.keys.discard(key)
        self._drop_if_empty(cur)

    def getMaxKey(self) -> str:
        return next(iter(self.tail.prev.keys)) if self.tail.prev is not self.head else ""

    def getMinKey(self) -> str:
        return next(iter(self.head.next.keys)) if self.head.next is not self.tail else ""
# test
a = AllOne()
a.inc("hello"); a.inc("hello")
assert a.getMaxKey() == "hello" and a.getMinKey() == "hello"
a.inc("leet")
assert a.getMaxKey() == "hello" and a.getMinKey() == "leet"
a.dec("hello"); a.dec("hello")
assert a.getMaxKey() == "leet" and a.getMinKey() == "leet"
a.dec("leet")
assert a.getMaxKey() == "" and a.getMinKey() == ""
```

```java
class AllOne {
    private static class Bucket {
        int cnt; Set<String> keys = new HashSet<>(); Bucket prev, next;
        Bucket(int c) { cnt = c; }
    }

    private final Bucket head = new Bucket(0), tail = new Bucket(0);
    private final Map<String, Bucket> where = new HashMap<>();

    public AllOne() { head.next = tail; tail.prev = head; }

    private Bucket after(Bucket node, int cnt) {
        Bucket b = new Bucket(cnt);
        b.prev = node; b.next = node.next;
        node.next.prev = b; node.next = b;
        return b;
    }

    private void dropIfEmpty(Bucket b) {
        if (b.keys.isEmpty()) { b.prev.next = b.next; b.next.prev = b.prev; }
    }

    public void inc(String key) {
        Bucket cur = where.getOrDefault(key, head), nxt = cur.next;
        if (nxt == tail || nxt.cnt != cur.cnt + 1) nxt = after(cur, cur.cnt + 1);
        nxt.keys.add(key);
        where.put(key, nxt);
        if (cur != head) { cur.keys.remove(key); dropIfEmpty(cur); }
    }

    public void dec(String key) {
        Bucket cur = where.get(key);
        if (cur.cnt == 1) where.remove(key);
        else {
            Bucket prv = cur.prev;
            if (prv == head || prv.cnt != cur.cnt - 1) prv = after(cur.prev, cur.cnt - 1);
            prv.keys.add(key);
            where.put(key, prv);
        }
        cur.keys.remove(key);
        dropIfEmpty(cur);
    }

    public String getMaxKey() { return tail.prev == head ? "" : tail.prev.keys.iterator().next(); }
    public String getMinKey() { return head.next == tail ? "" : head.next.keys.iterator().next(); }
}
// test
AllOne a = new AllOne();
a.inc("hello"); a.inc("hello"); a.inc("leet");
check(a.getMaxKey(), "hello");
check(a.getMinKey(), "leet");
a.dec("hello"); a.dec("hello");
check(a.getMaxKey(), "leet");
```

```go
type bucket struct {
	cnt        int
	keys       map[string]bool
	prev, next *bucket
}

type AllOne struct {
	head, tail *bucket            // sentinels; counts increase from head to tail
	where      map[string]*bucket // key -> the bucket holding it
}

func Constructor() AllOne {
	head, tail := &bucket{}, &bucket{}
	head.next, tail.prev = tail, head
	return AllOne{head: head, tail: tail, where: map[string]*bucket{}}
}

func (this *AllOne) after(node *bucket, cnt int) *bucket {
	b := &bucket{cnt: cnt, keys: map[string]bool{}, prev: node, next: node.next}
	node.next.prev = b
	node.next = b
	return b
}

func (this *AllOne) dropIfEmpty(b *bucket) {
	if len(b.keys) == 0 {
		b.prev.next, b.next.prev = b.next, b.prev
	}
}

func (this *AllOne) Inc(key string) {
	cur, ok := this.where[key]
	if !ok {
		cur = this.head // count 0
	}
	nxt := cur.next
	if nxt == this.tail || nxt.cnt != cur.cnt+1 {
		nxt = this.after(cur, cur.cnt+1)
	}
	nxt.keys[key] = true
	this.where[key] = nxt
	if cur != this.head {
		delete(cur.keys, key)
		this.dropIfEmpty(cur)
	}
}

func (this *AllOne) Dec(key string) {
	cur := this.where[key]
	if cur.cnt == 1 {
		delete(this.where, key)
	} else {
		prv := cur.prev
		if prv == this.head || prv.cnt != cur.cnt-1 {
			prv = this.after(cur.prev, cur.cnt-1)
		}
		prv.keys[key] = true
		this.where[key] = prv
	}
	delete(cur.keys, key)
	this.dropIfEmpty(cur)
}

func anyKey(b *bucket) string {
	for k := range b.keys {
		return k
	}
	return ""
}

func (this *AllOne) GetMaxKey() string {
	if this.tail.prev == this.head {
		return ""
	}
	return anyKey(this.tail.prev)
}

func (this *AllOne) GetMinKey() string {
	if this.head.next == this.tail {
		return ""
	}
	return anyKey(this.head.next)
}
// test
a := Constructor()
a.Inc("hello")
a.Inc("hello")
assertEq([]string{a.GetMaxKey(), a.GetMinKey()}, []string{"hello", "hello"})
a.Inc("leet")
assertEq([]string{a.GetMaxKey(), a.GetMinKey()}, []string{"hello", "leet"})
a.Dec("hello")
a.Dec("hello")
assertEq([]string{a.GetMaxKey(), a.GetMinKey()}, []string{"leet", "leet"})
a.Dec("leet")
assertEq([]string{a.GetMaxKey(), a.GetMinKey()}, []string{"", ""})
```

## 895. Maximum Frequency Stack | Hard
related: 460, 155, 716
tags: design, stack, buckets

```problem
Design `FreqStack`: `push(val)` pushes a value, and `pop()` removes and returns the most frequent value in the stack; among values tied for most frequent, it pops the one pushed most recently.

**Example:** push 5, 7, 5, 7, 4, 5 → pop() = 5, then 7, then 5, then 4

**Constraints:** up to 2·10⁴ calls.
```

Keep `freq[val]` and one **stack per frequency level**: when a value reaches frequency f, push it onto `group[f]` (it stays in the lower levels too). `pop` takes from `group[maxFreq]` — the most recent among the most frequent — decrements its frequency, and lowers `maxFreq` when that level empties.

```viz
push 5 7 5 7 4 5      group[1] = [5 7 4]   group[2] = [5 7]   group[3] = [5]
pop -> 5 (group[3])   pop -> 7 (group[2], newest)   pop -> 5   pop -> 4
```

**Complexity:** O(1) per operation.

```python
class FreqStack:
    def __init__(self):
        self.freq = Counter()
        self.group = defaultdict(list)
        self.maxf = 0

    def push(self, val: int) -> None:
        f = self.freq[val] + 1
        self.freq[val] = f
        self.group[f].append(val)
        self.maxf = max(self.maxf, f)

    def pop(self) -> int:
        val = self.group[self.maxf].pop()
        self.freq[val] -= 1
        if not self.group[self.maxf]:
            self.maxf -= 1
        return val
# test
fs = FreqStack()
for v in (5, 7, 5, 7, 4, 5):
    fs.push(v)
assert [fs.pop(), fs.pop(), fs.pop(), fs.pop()] == [5, 7, 5, 4]
fs.push(9)
assert [fs.pop(), fs.pop(), fs.pop()] == [9, 7, 5]
```

```java
class FreqStack {
    private final Map<Integer, Integer> freq = new HashMap<>();
    private final Map<Integer, Deque<Integer>> group = new HashMap<>();
    private int maxf = 0;

    public void push(int val) {
        int f = freq.merge(val, 1, Integer::sum);
        group.computeIfAbsent(f, k -> new ArrayDeque<>()).push(val);
        maxf = Math.max(maxf, f);
    }

    public int pop() {
        int val = group.get(maxf).pop();
        freq.merge(val, -1, Integer::sum);
        if (group.get(maxf).isEmpty()) maxf--;
        return val;
    }
}
// test
FreqStack fs = new FreqStack();
for (int v : new int[]{5, 7, 5, 7, 4, 5}) fs.push(v);
check(new int[]{fs.pop(), fs.pop(), fs.pop(), fs.pop()}, new int[]{5, 7, 5, 4});
```

```go
type FreqStack struct {
	freq  map[int]int   // value -> current frequency
	group map[int][]int // frequency -> stack of values that reached it
	maxf  int
}

func Constructor() FreqStack {
	return FreqStack{freq: map[int]int{}, group: map[int][]int{}}
}

func (this *FreqStack) Push(val int) {
	f := this.freq[val] + 1
	this.freq[val] = f
	this.group[f] = append(this.group[f], val) // val stays in the lower groups too
	this.maxf = max(this.maxf, f)
}

func (this *FreqStack) Pop() int {
	g := this.group[this.maxf]
	val := g[len(g)-1]
	this.group[this.maxf] = g[:len(g)-1]
	this.freq[val]--
	if len(g) == 1 {
		this.maxf--
	}
	return val
}
// test
fs := Constructor()
for _, v := range []int{5, 7, 5, 7, 4, 5} {
	fs.Push(v)
}
assertEq([]int{fs.Pop(), fs.Pop(), fs.Pop(), fs.Pop()}, []int{5, 7, 5, 4})
fs.Push(9)
assertEq([]int{fs.Pop(), fs.Pop(), fs.Pop()}, []int{9, 7, 5})
```

## 588. Design In-Memory File System | Hard
related: 208, 71, 1166
tags: design, trie, premium

```problem
Design an in-memory `FileSystem`: `ls(path)` returns `[name]` for a file path or the sorted names inside a directory; `mkdir(path)` creates the directory and any missing parents; `addContentToFile(filePath, content)` creates the file or appends to it; `readContentFromFile(filePath)` returns the file's content.

**Example:** `ls("/")` → `[]`; `mkdir("/a/b/c")`; `addContentToFile("/a/b/c/d","hello")` → `ls("/")` = `["a"]`, `readContentFromFile("/a/b/c/d")` = `"hello"`
```

**Problem (Premium):** implement `ls(path)` (a file path → `[file name]`; a directory → its children's names sorted), `mkdir(path)` (creating missing parents), `addContentToFile(filePath, content)` (create or append) and `readContentFromFile(filePath)`. Paths are absolute, like `/a/b/c`.

A **trie keyed by path components**: each node has a sorted map of children and, for files, a content buffer. Walking a path is O(depth); `ls` of a directory sorts its children (or keep a sorted map).

**Complexity:** O(path length) per walk, plus O(k log k) for `ls` of a directory with k entries.

```python
class FileSystem:
    def __init__(self):
        self.root = {"kids": {}, "file": None}       # file: None for a directory, else its content

    def _walk(self, path):
        node = self.root
        for part in filter(None, path.split("/")):
            node = node["kids"].setdefault(part, {"kids": {}, "file": None})
        return node

    def ls(self, path: str) -> List[str]:
        node = self._walk(path)
        return [path.rsplit("/", 1)[1]] if node["file"] is not None else sorted(node["kids"])

    def mkdir(self, path: str) -> None:
        self._walk(path)

    def addContentToFile(self, filePath: str, content: str) -> None:
        node = self._walk(filePath)
        node["file"] = (node["file"] or "") + content

    def readContentFromFile(self, filePath: str) -> str:
        return self._walk(filePath)["file"]
# test
fs = FileSystem()
assert fs.ls("/") == []
fs.mkdir("/a/b/c")
fs.addContentToFile("/a/b/c/d", "hello")
assert fs.ls("/") == ["a"]
assert fs.readContentFromFile("/a/b/c/d") == "hello"
fs.addContentToFile("/a/b/c/d", " world")
assert fs.readContentFromFile("/a/b/c/d") == "hello world"
assert fs.ls("/a/b/c/d") == ["d"]
fs.mkdir("/a/b/e")
fs.addContentToFile("/a/b/aa", "x")
assert fs.ls("/a/b") == ["aa", "c", "e"]
```

```java
class FileSystem {
    private static class Node { TreeMap<String, Node> kids = new TreeMap<>(); StringBuilder file; }
    private final Node root = new Node();

    private Node walk(String path) {
        Node node = root;
        for (String part : path.split("/")) if (!part.isEmpty()) node = node.kids.computeIfAbsent(part, k -> new Node());
        return node;
    }

    public List<String> ls(String path) {
        Node node = walk(path);
        if (node.file != null) return List.of(path.substring(path.lastIndexOf('/') + 1));
        return new ArrayList<>(node.kids.keySet());
    }

    public void mkdir(String path) { walk(path); }

    public void addContentToFile(String filePath, String content) {
        Node node = walk(filePath);
        if (node.file == null) node.file = new StringBuilder();
        node.file.append(content);
    }

    public String readContentFromFile(String filePath) { return walk(filePath).file.toString(); }
}
// test
FileSystem fs = new FileSystem();
check(fs.ls("/"), List.of());
fs.mkdir("/a/b/c");
fs.addContentToFile("/a/b/c/d", "hello");
check(fs.ls("/"), List.of("a"));
check(fs.ls("/a/b/c/d"), List.of("d"));
check(fs.readContentFromFile("/a/b/c/d"), "hello");
```

```go
type fsNode struct {
	kids    map[string]*fsNode
	isFile  bool
	content string
}

type FileSystem struct {
	root *fsNode
}

func Constructor() FileSystem {
	return FileSystem{root: &fsNode{kids: map[string]*fsNode{}}}
}

func (this *FileSystem) walk(path string) *fsNode { // creates missing directories on the way
	node := this.root
	for _, part := range strings.Split(path, "/") {
		if part == "" {
			continue
		}
		next, ok := node.kids[part]
		if !ok {
			next = &fsNode{kids: map[string]*fsNode{}}
			node.kids[part] = next
		}
		node = next
	}
	return node
}

func (this *FileSystem) Ls(path string) []string {
	node := this.walk(path)
	if node.isFile {
		return []string{path[strings.LastIndex(path, "/")+1:]}
	}
	names := []string{}
	for name := range node.kids {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func (this *FileSystem) Mkdir(path string) {
	this.walk(path)
}

func (this *FileSystem) AddContentToFile(filePath string, content string) {
	node := this.walk(filePath)
	node.isFile = true
	node.content += content
}

func (this *FileSystem) ReadContentFromFile(filePath string) string {
	return this.walk(filePath).content
}
// test
fs := Constructor()
assertEq(fs.Ls("/"), []string{})
fs.Mkdir("/a/b/c")
fs.AddContentToFile("/a/b/c/d", "hello")
assertEq(fs.Ls("/"), []string{"a"})
assertEq(fs.ReadContentFromFile("/a/b/c/d"), "hello")
fs.AddContentToFile("/a/b/c/d", " world")
assertEq(fs.ReadContentFromFile("/a/b/c/d"), "hello world")
assertEq(fs.Ls("/a/b/c/d"), []string{"d"})
fs.Mkdir("/a/b/e")
fs.AddContentToFile("/a/b/aa", "x")
assertEq(fs.Ls("/a/b"), []string{"aa", "c", "e"})
```

## 362. Design Hit Counter | Medium
related: 933, 359, 1396
tags: design, queue, premium

```problem
Design `HitCounter`: `hit(timestamp)` records a hit at that second, and `getHits(timestamp)` returns the number of hits in the past 5 minutes (the 300 seconds up to and including `timestamp`). Timestamps never decrease, and several hits can share one timestamp.

**Example:** hits at 1, 2, 3 → getHits(4) = 3; hit at 300 → getHits(300) = 4, getHits(301) = 3

**Constraints:** follow-up: a huge number of hits per second.
```

**Problem (Premium):** `hit(timestamp)` records a hit (several may share a timestamp); `getHits(timestamp)` returns the hits in the past 5 minutes — timestamps in `(t − 300, t]`. Calls arrive in non-decreasing time order.

A queue of `[timestamp, count]` buckets (collapsing equal timestamps) plus a running total: expire buckets from the front on read. For huge hit rates, the follow-up answer is a fixed ring of 300 one-second buckets (`times[t % 300]`, `counts[t % 300]`), O(300) memory regardless of traffic.

**Complexity:** amortized O(1) per call.

```python
class HitCounter:
    def __init__(self):
        self.q = deque()          # [timestamp, count], oldest first
        self.total = 0

    def hit(self, timestamp: int) -> None:
        if self.q and self.q[-1][0] == timestamp:
            self.q[-1][1] += 1
        else:
            self.q.append([timestamp, 1])
        self.total += 1

    def getHits(self, timestamp: int) -> int:
        while self.q and self.q[0][0] <= timestamp - 300:
            self.total -= self.q.popleft()[1]
        return self.total
# test
hc = HitCounter()
hc.hit(1); hc.hit(2); hc.hit(3)
assert hc.getHits(4) == 3
hc.hit(300)
assert hc.getHits(300) == 4 and hc.getHits(301) == 3
hc2 = HitCounter()
for _ in range(1000):
    hc2.hit(1)
assert hc2.getHits(300) == 1000 and hc2.getHits(301) == 0
```

```java
class HitCounter {
    private final int[] times = new int[300], counts = new int[300];   // ring of one-second buckets

    public void hit(int timestamp) {
        int i = timestamp % 300;
        if (times[i] != timestamp) { times[i] = timestamp; counts[i] = 0; }
        counts[i]++;
    }

    public int getHits(int timestamp) {
        int total = 0;
        for (int i = 0; i < 300; i++) if (timestamp - times[i] < 300) total += counts[i];
        return total;
    }
}
// test
HitCounter hc = new HitCounter();
hc.hit(1); hc.hit(2); hc.hit(3);
check(hc.getHits(4), 3);
hc.hit(300);
check(hc.getHits(300), 4);
check(hc.getHits(301), 3);
```

```go
type HitCounter struct {
	q     [][2]int // (timestamp, hits), oldest first
	total int
}

func Constructor() HitCounter {
	return HitCounter{}
}

func (this *HitCounter) Hit(timestamp int) {
	if n := len(this.q); n > 0 && this.q[n-1][0] == timestamp { // same second: one entry
		this.q[n-1][1]++
	} else {
		this.q = append(this.q, [2]int{timestamp, 1})
	}
	this.total++
}

func (this *HitCounter) GetHits(timestamp int) int {
	for len(this.q) > 0 && this.q[0][0] <= timestamp-300 { // older than 5 minutes
		this.total -= this.q[0][1]
		this.q = this.q[1:]
	}
	return this.total
}
// test
hc := Constructor()
hc.Hit(1)
hc.Hit(2)
hc.Hit(3)
assertEq(hc.GetHits(4), 3)
hc.Hit(300)
assertEq([]int{hc.GetHits(300), hc.GetHits(301)}, []int{4, 3})
hc2 := Constructor()
for i := 0; i < 1000; i++ {
	hc2.Hit(1)
}
assertEq([]int{hc2.GetHits(300), hc2.GetHits(301)}, []int{1000, 0})
```

## 933. Number of Recent Calls | Easy
related: 362, 346
tags: queue, design

```problem
Design `RecentCounter`, whose `ping(t)` records a request at time t (in ms, strictly increasing) and returns the number of requests in the window [t − 3000, t].

**Example:** pings at 1, 100, 3001, 3002 → `1, 2, 3, 3`
```

Pings arrive in increasing time order, so a queue works: push `t`, pop from the front while it's older than `t − 3000`; the queue length is the answer.

**Complexity:** amortized O(1) per ping.

```python
class RecentCounter:
    def __init__(self):
        self.q = deque()

    def ping(self, t: int) -> int:
        self.q.append(t)
        while self.q[0] < t - 3000:
            self.q.popleft()
        return len(self.q)
# test
rc = RecentCounter()
assert [rc.ping(t) for t in (1, 100, 3001, 3002)] == [1, 2, 3, 3]
rc2 = RecentCounter()
assert [rc2.ping(t) for t in (1, 3001, 3002, 6003)] == [1, 2, 2, 1]
```

```java
class RecentCounter {
    private final Deque<Integer> q = new ArrayDeque<>();

    public int ping(int t) {
        q.add(t);
        while (q.peek() < t - 3000) q.poll();
        return q.size();
    }
}
// test
RecentCounter rc = new RecentCounter();
check(new int[]{rc.ping(1), rc.ping(100), rc.ping(3001), rc.ping(3002)}, new int[]{1, 2, 3, 3});
```

```go
type RecentCounter struct {
	q []int // ping times within the last 3000 ms
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.q = append(this.q, t)
	for this.q[0] < t-3000 {
		this.q = this.q[1:]
	}
	return len(this.q)
}
// test
rc := Constructor()
assertEq([]int{rc.Ping(1), rc.Ping(100), rc.Ping(3001), rc.Ping(3002)}, []int{1, 2, 3, 3})
rc2 := Constructor()
assertEq([]int{rc2.Ping(1), rc2.Ping(3001), rc2.Ping(3002), rc2.Ping(6003)}, []int{1, 2, 2, 1})
```

## 705. Design HashSet | Easy
related: 706, 380
tags: design, hashing

```problem
Design `MyHashSet` without built-in hash tables: `add(key)`, `remove(key)` and `contains(key)`.

**Example:** add 1, add 2 → contains(1) = true, contains(3) = false; add 2, remove 2 → contains(2) = false

**Constraints:** keys ≤ 10⁶; up to 10⁴ calls.
```

Separate chaining: an array of buckets (a prime count spreads keys evenly), each a small list; hash = `key % buckets`. Operations scan one bucket. Follow-ups: resize (rehash into 2× buckets) when the load factor grows; open addressing with tombstones as the alternative.

**Complexity:** O(1) average per operation.

```python
class MyHashSet:
    def __init__(self):
        self.size = 1009
        self.buckets = [[] for _ in range(self.size)]

    def add(self, key: int) -> None:
        b = self.buckets[key % self.size]
        if key not in b:
            b.append(key)

    def remove(self, key: int) -> None:
        b = self.buckets[key % self.size]
        if key in b:
            b.remove(key)

    def contains(self, key: int) -> bool:
        return key in self.buckets[key % self.size]
# test
hs = MyHashSet()
hs.add(1); hs.add(2)
assert hs.contains(1) and not hs.contains(3)
hs.add(2)
assert hs.contains(2)
hs.remove(2)
assert not hs.contains(2)
hs.add(0); hs.add(1009); hs.add(2018)                 # same bucket
hs.remove(1009)
assert hs.contains(0) and not hs.contains(1009) and hs.contains(2018)
```

```java
class MyHashSet {
    private static final int SIZE = 1009;
    private final List<List<Integer>> buckets = new ArrayList<>();

    public MyHashSet() { for (int i = 0; i < SIZE; i++) buckets.add(new LinkedList<>()); }

    public void add(int key) { if (!contains(key)) buckets.get(key % SIZE).add(key); }

    public void remove(int key) { buckets.get(key % SIZE).remove(Integer.valueOf(key)); }

    public boolean contains(int key) { return buckets.get(key % SIZE).contains(key); }
}
// test
MyHashSet hs = new MyHashSet();
hs.add(1); hs.add(1009); hs.remove(1);
check(hs.contains(1) + " " + hs.contains(1009), "false true");
```

```go
type MyHashSet struct {
	buckets [][]int // separate chaining
}

const size = 1009 // a prime number of buckets

func Constructor() MyHashSet {
	return MyHashSet{buckets: make([][]int, size)}
}

func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		this.buckets[key%size] = append(this.buckets[key%size], key)
	}
}

func (this *MyHashSet) Remove(key int) {
	b := this.buckets[key%size]
	for i, k := range b {
		if k == key {
			this.buckets[key%size] = append(b[:i], b[i+1:]...)
			return
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	for _, k := range this.buckets[key%size] {
		if k == key {
			return true
		}
	}
	return false
}
// test
hs := Constructor()
hs.Add(1)
hs.Add(2)
assertEq([]bool{hs.Contains(1), hs.Contains(3)}, []bool{true, false})
hs.Add(2)
assertEq(hs.Contains(2), true)
hs.Remove(2)
assertEq(hs.Contains(2), false)
hs.Add(0) // 0, 1009 and 2018 share a bucket
hs.Add(1009)
hs.Add(2018)
hs.Remove(1009)
assertEq([]bool{hs.Contains(0), hs.Contains(1009), hs.Contains(2018)}, []bool{true, false, true})
```
