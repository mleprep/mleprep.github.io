# Heap / Priority Queue

**When to use:** "k-th largest/smallest", "top k", repeatedly take the best item, merge k sorted streams, scheduling, running median.

- **Top-k largest** → keep a **min-heap of size k** (evict the smallest). O(n log k).
- Python `heapq` is a min-heap: push `-x` for a max-heap, push tuples `(priority, tiebreak, item)`.
- Java `PriorityQueue<>((a, b) -> ...)`; max-heap via `Collections.reverseOrder()`.
- **Two heaps** (max-heap low half + min-heap high half) for medians.

```python
heap = []
for x in nums:
    heapq.heappush(heap, x)
    if len(heap) > k:
        heapq.heappop(heap)      # heap[0] is now the k-th largest
```

## 703. Kth Largest Element in a Stream | Easy
related: 215, 295
tags: min-heap

Min-heap of the k largest seen so far; its root is the answer.

**Complexity:** O(log k) per add.

```python
class KthLargest:
    def __init__(self, k: int, nums: List[int]):
        self.k, self.h = k, []
        for x in nums:
            self.add(x)

    def add(self, val: int) -> int:
        heappush(self.h, val)
        if len(self.h) > self.k:
            heappop(self.h)
        return self.h[0]
# test
kl = KthLargest(3, [4, 5, 8, 2])
assert [kl.add(v) for v in (3, 5, 10, 9, 4)] == [4, 5, 5, 8, 8]
kl = KthLargest(1, [])
assert [kl.add(v) for v in (-3, -2, -4, 0, 4)] == [-3, -2, -2, 0, 4]
```

```java
class KthLargest {
    private final PriorityQueue<Integer> pq = new PriorityQueue<>();
    private final int k;

    public KthLargest(int k, int[] nums) {
        this.k = k;
        for (int x : nums) add(x);
    }

    public int add(int val) {
        pq.add(val);
        if (pq.size() > k) pq.poll();
        return pq.peek();
    }
}
// test
KthLargest kl = new KthLargest(3, new int[]{4, 5, 8, 2});
int[] got = {kl.add(3), kl.add(5), kl.add(10), kl.add(9), kl.add(4)};
check(got, new int[]{4, 5, 5, 8, 8});
```

```go
type IntHeap []int // min-heap

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

type KthLargest struct {
	k int
	h IntHeap // the k largest values; the root is the k-th largest
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	for _, x := range nums {
		kl.Add(x)
	}
	return kl
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.h, val)
	if this.h.Len() > this.k {
		heap.Pop(&this.h)
	}
	return this.h[0]
}
// test
kl := Constructor(3, []int{4, 5, 8, 2})
got := []int{}
for _, v := range []int{3, 5, 10, 9, 4} {
	got = append(got, kl.Add(v))
}
assertEq(got, []int{4, 5, 5, 8, 8})
kl2 := Constructor(1, []int{})
got2 := []int{}
for _, v := range []int{-3, -2, -4, 0, 4} {
	got2 = append(got2, kl2.Add(v))
}
assertEq(got2, []int{-3, -2, -2, 0, 4})
```

## 1046. Last Stone Weight | Easy
related: 215
tags: max-heap

Simulate with a max-heap: smash the two heaviest, push back the difference.

**Complexity:** O(n log n).

```python
class Solution:
    def lastStoneWeight(self, stones: List[int]) -> int:
        h = [-s for s in stones]
        heapify(h)
        while len(h) > 1:
            a, b = -heappop(h), -heappop(h)
            if a != b:
                heappush(h, -(a - b))
        return -h[0] if h else 0
# test
assert Solution().lastStoneWeight([2, 7, 4, 1, 8, 1]) == 1
assert Solution().lastStoneWeight([1]) == 1
```

```java
class Solution {
    public int lastStoneWeight(int[] stones) {
        PriorityQueue<Integer> pq = new PriorityQueue<>(Collections.reverseOrder());
        for (int s : stones) pq.add(s);
        while (pq.size() > 1) {
            int a = pq.poll(), b = pq.poll();
            if (a != b) pq.add(a - b);
        }
        return pq.isEmpty() ? 0 : pq.peek();
    }
}
// test
check(new Solution().lastStoneWeight(new int[]{2, 7, 4, 1, 8, 1}), 1);
```

```go
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func lastStoneWeight(stones []int) int {
	h := MaxHeap(append([]int{}, stones...))
	heap.Init(&h)
	for h.Len() > 1 { // smash the two heaviest
		a, b := heap.Pop(&h).(int), heap.Pop(&h).(int)
		if a != b {
			heap.Push(&h, a-b)
		}
	}
	if h.Len() == 0 {
		return 0
	}
	return h[0]
}
// test
assertEq(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}), 1)
assertEq(lastStoneWeight([]int{1}), 1)
```

## 973. K Closest Points to Origin | Medium
related: 215, 347
tags: max-heap, quickselect

Keep a **max-heap of size k** on squared distance; evict the farthest. (Quickselect gives average O(n).)

**Complexity:** O(n log k).

```python
class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        h = []
        for x, y in points:
            heappush(h, (-(x * x + y * y), x, y))
            if len(h) > k:
                heappop(h)
        return [[x, y] for _, x, y in h]
# test
assert Solution().kClosest([[1, 3], [-2, 2]], 1) == [[-2, 2]]
assert sorted(Solution().kClosest([[3, 3], [5, -1], [-2, 4]], 2)) == [[-2, 4], [3, 3]]
```

```java
class Solution {
    public int[][] kClosest(int[][] points, int k) {
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> (b[0] * b[0] + b[1] * b[1]) - (a[0] * a[0] + a[1] * a[1]));
        for (int[] p : points) {
            pq.add(p);
            if (pq.size() > k) pq.poll();
        }
        return pq.toArray(new int[0][]);
    }
}
// test
check(new Solution().kClosest(new int[][]{{1, 3}, {-2, 2}}, 1), new int[][]{{-2, 2}});
```

```go
type point struct{ d, x, y int }
type pointHeap []point // max-heap by distance

func (h pointHeap) Len() int           { return len(h) }
func (h pointHeap) Less(i, j int) bool { return h[i].d > h[j].d }
func (h pointHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *pointHeap) Push(x any)        { *h = append(*h, x.(point)) }
func (h *pointHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	h := &pointHeap{}
	for _, p := range points {
		heap.Push(h, point{p[0]*p[0] + p[1]*p[1], p[0], p[1]})
		if h.Len() > k { // drop the farthest
			heap.Pop(h)
		}
	}
	res := [][]int{}
	for _, p := range *h {
		res = append(res, []int{p.x, p.y})
	}
	return res
}
// test
assertEq(kClosest([][]int{{1, 3}, {-2, 2}}, 1), [][]int{{-2, 2}})
assertUnordered(kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2), [][]int{{-2, 4}, {3, 3}})
```

## 215. Kth Largest Element in an Array | Medium
related: 703, 347, 973
tags: min-heap, quickselect

Min-heap of size k → root is the k-th largest. Interview follow-up: **quickselect** (partition around a pivot, recurse into one side) for average O(n).

**Complexity:** heap O(n log k); quickselect O(n) average.

```python
class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        h = []
        for x in nums:
            heappush(h, x)
            if len(h) > k:
                heappop(h)
        return h[0]
# test
assert Solution().findKthLargest([3, 2, 1, 5, 6, 4], 2) == 5
assert Solution().findKthLargest([3, 2, 3, 1, 2, 4, 5, 5, 6], 4) == 4
```

```java
class Solution {
    // quickselect with 3-way partition (random pivot): average O(n)
    public int findKthLargest(int[] nums, int k) {
        int target = nums.length - k, lo = 0, hi = nums.length - 1;
        Random rnd = new Random(7);
        while (true) {
            int pivot = nums[lo + rnd.nextInt(hi - lo + 1)];
            int lt = lo, i = lo, gt = hi;          // [lo,lt) < p, [lt,i) == p, (gt,hi] > p
            while (i <= gt) {
                if (nums[i] < pivot) swap(nums, lt++, i++);
                else if (nums[i] > pivot) swap(nums, i, gt--);
                else i++;
            }
            if (target < lt) hi = lt - 1;
            else if (target > gt) lo = gt + 1;
            else return pivot;
        }
    }

    private void swap(int[] a, int i, int j) { int t = a[i]; a[i] = a[j]; a[j] = t; }
}
// test
check(new Solution().findKthLargest(new int[]{3, 2, 1, 5, 6, 4}, 2), 5);
check(new Solution().findKthLargest(new int[]{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4), 4);
```

```go
type IntHeap []int // min-heap

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &IntHeap{} // the k largest seen so far
	for _, x := range nums {
		heap.Push(h, x)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}
// test
assertEq(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2), 5)
assertEq(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4), 4)
```

## 621. Task Scheduler | Medium
related: 767, 358
tags: greedy, counting

The most frequent task(s) dictate the frame: `(maxf − 1)` full blocks of size `n + 1`, plus one slot per task tied at `maxf`. If there are more tasks than slots, no idle is needed → answer is `len(tasks)`.

```viz
tasks A A A B B B, n=2      maxf=3, tied=2
A B _ | A B _ | A B         (3-1)*(2+1) + 2 = 8
```

**Complexity:** O(n).

```python
class Solution:
    def leastInterval(self, tasks: List[str], n: int) -> int:
        cnt = Counter(tasks).values()
        maxf = max(cnt)
        tied = sum(1 for c in cnt if c == maxf)
        return max(len(tasks), (maxf - 1) * (n + 1) + tied)
# test
assert Solution().leastInterval(["A", "A", "A", "B", "B", "B"], 2) == 8
assert Solution().leastInterval(["A", "C", "A", "B", "D", "B"], 1) == 6
assert Solution().leastInterval(["A", "A", "A", "B", "B", "B"], 3) == 10
```

```java
class Solution {
    public int leastInterval(char[] tasks, int n) {
        int[] cnt = new int[26];
        int maxf = 0, tied = 0;
        for (char t : tasks) maxf = Math.max(maxf, ++cnt[t - 'A']);
        for (int c : cnt) if (c == maxf) tied++;
        return Math.max(tasks.length, (maxf - 1) * (n + 1) + tied);
    }
}
// test
check(new Solution().leastInterval("AAABBB".toCharArray(), 2), 8);
check(new Solution().leastInterval("AAABBB".toCharArray(), 3), 10);
```

```go
func leastInterval(tasks []byte, n int) int {
	var cnt [26]int
	for _, t := range tasks {
		cnt[t-'A']++
	}
	maxf, tied := 0, 0
	for _, c := range cnt {
		if c > maxf {
			maxf, tied = c, 1
		} else if c == maxf {
			tied++
		}
	}
	// (maxf-1) full frames of n+1 slots, plus the tasks tied for most frequent
	return max(len(tasks), (maxf-1)*(n+1)+tied)
}
// test
assertEq(leastInterval([]byte("AAABBB"), 2), 8)
assertEq(leastInterval([]byte("ACABDB"), 1), 6)
assertEq(leastInterval([]byte("AAABBB"), 3), 10)
```

## 355. Design Twitter | Medium
related: 23
tags: design, k-way-merge

Store each user's tweets as `(time, id)` lists. `getNewsFeed` = k-way merge of the followees' lists (newest first) with a heap, stopping at 10.

**Complexity:** feed O(F log F) for F followees (each contributes ≤ 10).

```python
class Twitter:
    def __init__(self):
        self.t = 0
        self.tweets = defaultdict(list)    # user -> [(time, id)]
        self.follows = defaultdict(set)

    def postTweet(self, userId: int, tweetId: int) -> None:
        self.t += 1
        self.tweets[userId].append((self.t, tweetId))

    def getNewsFeed(self, userId: int) -> List[int]:
        h = []
        for u in self.follows[userId] | {userId}:
            if self.tweets[u]:
                i = len(self.tweets[u]) - 1
                t, tid = self.tweets[u][i]
                heappush(h, (-t, tid, u, i))
        res = []
        while h and len(res) < 10:
            _, tid, u, i = heappop(h)
            res.append(tid)
            if i:
                t, nid = self.tweets[u][i - 1]
                heappush(h, (-t, nid, u, i - 1))
        return res

    def follow(self, followerId: int, followeeId: int) -> None:
        self.follows[followerId].add(followeeId)

    def unfollow(self, followerId: int, followeeId: int) -> None:
        self.follows[followerId].discard(followeeId)
# test
tw = Twitter()
tw.postTweet(1, 5)
assert tw.getNewsFeed(1) == [5]
tw.follow(1, 2); tw.postTweet(2, 6)
assert tw.getNewsFeed(1) == [6, 5]
tw.unfollow(1, 2)
assert tw.getNewsFeed(1) == [5]
```

```java
class Twitter {
    private int time = 0;
    private final Map<Integer, List<int[]>> tweets = new HashMap<>();   // user -> [{time, id}]
    private final Map<Integer, Set<Integer>> follows = new HashMap<>();

    public void postTweet(int userId, int tweetId) {
        tweets.computeIfAbsent(userId, k -> new ArrayList<>()).add(new int[]{time++, tweetId});
    }

    public List<Integer> getNewsFeed(int userId) {
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> b[0] - a[0]); // {time, id, user, idx}
        Set<Integer> users = new HashSet<>(follows.getOrDefault(userId, Set.of()));
        users.add(userId);
        for (int u : users) {
            List<int[]> l = tweets.get(u);
            if (l != null && !l.isEmpty()) {
                int i = l.size() - 1;
                pq.add(new int[]{l.get(i)[0], l.get(i)[1], u, i});
            }
        }
        List<Integer> res = new ArrayList<>();
        while (!pq.isEmpty() && res.size() < 10) {
            int[] top = pq.poll();
            res.add(top[1]);
            if (top[3] > 0) {
                int[] prev = tweets.get(top[2]).get(top[3] - 1);
                pq.add(new int[]{prev[0], prev[1], top[2], top[3] - 1});
            }
        }
        return res;
    }

    public void follow(int followerId, int followeeId) {
        follows.computeIfAbsent(followerId, k -> new HashSet<>()).add(followeeId);
    }

    public void unfollow(int followerId, int followeeId) {
        Set<Integer> s = follows.get(followerId);
        if (s != null) s.remove(followeeId);
    }
}
// test
Twitter tw = new Twitter();
tw.postTweet(1, 5);
tw.follow(1, 2); tw.postTweet(2, 6);
check(tw.getNewsFeed(1), "[6, 5]");
tw.unfollow(1, 2);
check(tw.getNewsFeed(1), "[5]");
```

```go
type tweet struct{ time, id int }

type feedItem struct{ time, id, user, i int }
type feedHeap []feedItem // newest first

func (h feedHeap) Len() int           { return len(h) }
func (h feedHeap) Less(i, j int) bool { return h[i].time > h[j].time }
func (h feedHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *feedHeap) Push(x any)        { *h = append(*h, x.(feedItem)) }
func (h *feedHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

type Twitter struct {
	t       int
	tweets  map[int][]tweet
	follows map[int]map[int]bool
}

func Constructor() Twitter {
	return Twitter{tweets: map[int][]tweet{}, follows: map[int]map[int]bool{}}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.t++
	this.tweets[userId] = append(this.tweets[userId], tweet{this.t, tweetId})
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	h := &feedHeap{}
	users := []int{userId}
	for u := range this.follows[userId] {
		if u != userId {
			users = append(users, u)
		}
	}
	for _, u := range users { // merge each user's newest tweets, k-way
		if ts := this.tweets[u]; len(ts) > 0 {
			i := len(ts) - 1
			heap.Push(h, feedItem{ts[i].time, ts[i].id, u, i})
		}
	}
	res := []int{}
	for h.Len() > 0 && len(res) < 10 {
		it := heap.Pop(h).(feedItem)
		res = append(res, it.id)
		if it.i > 0 {
			prev := this.tweets[it.user][it.i-1]
			heap.Push(h, feedItem{prev.time, prev.id, it.user, it.i - 1})
		}
	}
	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.follows[followerId] == nil {
		this.follows[followerId] = map[int]bool{}
	}
	this.follows[followerId][followeeId] = true
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.follows[followerId], followeeId)
}
// test
tw := Constructor()
tw.PostTweet(1, 5)
assertEq(tw.GetNewsFeed(1), []int{5})
tw.Follow(1, 2)
tw.PostTweet(2, 6)
assertEq(tw.GetNewsFeed(1), []int{6, 5})
tw.Unfollow(1, 2)
assertEq(tw.GetNewsFeed(1), []int{5})
```

## 295. Find Median from Data Stream | Hard
related: 4, 480, 703
tags: blind75, two-heaps, design

`lo` = max-heap of the smaller half, `hi` = min-heap of the larger half, `len(lo) ∈ {len(hi), len(hi)+1}`. Push into `lo`, move its max to `hi`, rebalance. Median = top of `lo`, or the average of both tops.

```viz
stream 5 2 8 1
lo (max-heap)  [2, 1]    |    hi (min-heap) [5, 8]
median = (2 + 5) / 2 = 3.5
```

**Complexity:** O(log n) add, O(1) median.

```python
class MedianFinder:
    def __init__(self):
        self.lo, self.hi = [], []        # lo stores negatives (max-heap)

    def addNum(self, num: int) -> None:
        heappush(self.lo, -num)
        heappush(self.hi, -heappop(self.lo))
        if len(self.hi) > len(self.lo):
            heappush(self.lo, -heappop(self.hi))

    def findMedian(self) -> float:
        if len(self.lo) > len(self.hi):
            return float(-self.lo[0])
        return (-self.lo[0] + self.hi[0]) / 2
# test
mf = MedianFinder()
mf.addNum(1); mf.addNum(2)
assert mf.findMedian() == 1.5
mf.addNum(3)
assert mf.findMedian() == 2.0
```

```java
class MedianFinder {
    private final PriorityQueue<Integer> lo = new PriorityQueue<>(Collections.reverseOrder());
    private final PriorityQueue<Integer> hi = new PriorityQueue<>();

    public void addNum(int num) {
        lo.add(num);
        hi.add(lo.poll());
        if (hi.size() > lo.size()) lo.add(hi.poll());
    }

    public double findMedian() {
        return lo.size() > hi.size() ? lo.peek() : (lo.peek() + (double) hi.peek()) / 2;
    }
}
// test
MedianFinder mf = new MedianFinder();
mf.addNum(1); mf.addNum(2);
check(mf.findMedian(), 1.5);
mf.addNum(3);
check(mf.findMedian(), 2.0);
```

```go
type IntHeap []int // min-heap

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

type MedianFinder struct {
	lo IntHeap // smaller half, stored negated (a max-heap)
	hi IntHeap // larger half
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.lo, -num)
	heap.Push(&this.hi, -heap.Pop(&this.lo).(int)) // the largest of the small half moves up
	if this.hi.Len() > this.lo.Len() {
		heap.Push(&this.lo, -heap.Pop(&this.hi).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.lo.Len() > this.hi.Len() {
		return float64(-this.lo[0])
	}
	return float64(-this.lo[0]+this.hi[0]) / 2
}
// test
mf := Constructor()
mf.AddNum(1)
mf.AddNum(2)
assertEq(mf.FindMedian(), 1.5)
mf.AddNum(3)
assertEq(mf.FindMedian(), 2.0)
```

## 378. Kth Smallest Element in a Sorted Matrix | Medium
related: 23, 373, 240
tags: k-way-merge, binary-search

Rows are sorted lists → k-way merge: heap seeded with each row's first element; pop k−1 times, pushing the next in the same row. (Alt: binary search on value + count ≤ mid staircase, O(n log range).)

**Complexity:** O(k log n).

```python
class Solution:
    def kthSmallest(self, matrix: List[List[int]], k: int) -> int:
        n = len(matrix)
        h = [(matrix[r][0], r, 0) for r in range(min(n, k))]
        heapify(h)
        for _ in range(k - 1):
            _, r, c = heappop(h)
            if c + 1 < n:
                heappush(h, (matrix[r][c + 1], r, c + 1))
        return h[0][0]
# test
assert Solution().kthSmallest([[1, 5, 9], [10, 11, 13], [12, 13, 15]], 8) == 13
assert Solution().kthSmallest([[-5]], 1) == -5
```

```java
class Solution {
    public int kthSmallest(int[][] matrix, int k) {
        int n = matrix.length;
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[0] - b[0]); // {val, r, c}
        for (int r = 0; r < Math.min(n, k); r++) pq.add(new int[]{matrix[r][0], r, 0});
        for (int i = 0; i < k - 1; i++) {
            int[] t = pq.poll();
            if (t[2] + 1 < n) pq.add(new int[]{matrix[t[1]][t[2] + 1], t[1], t[2] + 1});
        }
        return pq.peek()[0];
    }
}
// test
check(new Solution().kthSmallest(new int[][]{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8), 13);
```

```go
type cell struct{ v, r, c int }
type cellHeap []cell

func (h cellHeap) Len() int           { return len(h) }
func (h cellHeap) Less(i, j int) bool { return h[i].v < h[j].v }
func (h cellHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *cellHeap) Push(x any)        { *h = append(*h, x.(cell)) }
func (h *cellHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	h := &cellHeap{}
	for r := 0; r < min(n, k); r++ { // the first column of each row
		heap.Push(h, cell{matrix[r][0], r, 0})
	}
	for i := 0; i < k-1; i++ { // pop k-1 smallest, advancing along their rows
		c := heap.Pop(h).(cell)
		if c.c+1 < n {
			heap.Push(h, cell{matrix[c.r][c.c+1], c.r, c.c + 1})
		}
	}
	return (*h)[0].v
}
// test
assertEq(kthSmallest([][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8), 13)
assertEq(kthSmallest([][]int{{-5}}, 1), -5)
```

## 767. Reorganize String | Medium
related: 621, 358
tags: greedy, max-heap

Greedy: always place the most frequent char that isn't the previous one. With a max-heap, pop the top, append it, then push back the *previous* char (held out for one turn). Impossible if any count > ⌈n/2⌉.

**Complexity:** O(n log 26).

```python
class Solution:
    def reorganizeString(self, s: str) -> str:
        h = [(-c, ch) for ch, c in Counter(s).items()]
        heapify(h)
        res, prev = [], None
        while h:
            c, ch = heappop(h)
            res.append(ch)
            if prev:
                heappush(h, prev)
            prev = (c + 1, ch) if c + 1 < 0 else None
        return "".join(res) if not prev else ""
# test
r = Solution().reorganizeString("aab")
assert r == "aba"
assert Solution().reorganizeString("aaab") == ""
```

```java
class Solution {
    public String reorganizeString(String s) {
        int[] cnt = new int[26];
        for (char c : s.toCharArray()) cnt[c - 'a']++;
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> b[1] - a[1]); // {char, count}
        for (int i = 0; i < 26; i++) if (cnt[i] > 0) pq.add(new int[]{i, cnt[i]});
        StringBuilder sb = new StringBuilder();
        int[] prev = null;
        while (!pq.isEmpty()) {
            int[] cur = pq.poll();
            sb.append((char) ('a' + cur[0]));
            if (prev != null) pq.add(prev);
            prev = --cur[1] > 0 ? cur : null;
        }
        return prev == null ? sb.toString() : "";
    }
}
// test
check(new Solution().reorganizeString("aab"), "aba");
check(new Solution().reorganizeString("aaab"), "");
```

```go
type charCount struct {
	n  int
	ch byte
}
type countHeap []charCount // most frequent first

func (h countHeap) Len() int           { return len(h) }
func (h countHeap) Less(i, j int) bool { return h[i].n > h[j].n }
func (h countHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *countHeap) Push(x any)        { *h = append(*h, x.(charCount)) }
func (h *countHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func reorganizeString(s string) string {
	var cnt [26]int
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}
	h := &countHeap{}
	for i, n := range cnt {
		if n > 0 {
			heap.Push(h, charCount{n, byte('a' + i)})
		}
	}
	res := []byte{}
	var prev *charCount // the letter just used waits one turn
	for h.Len() > 0 {
		cur := heap.Pop(h).(charCount)
		res = append(res, cur.ch)
		if prev != nil {
			heap.Push(h, *prev)
			prev = nil
		}
		if cur.n--; cur.n > 0 {
			prev = &cur
		}
	}
	if prev != nil { // a letter is left over with nothing to separate it
		return ""
	}
	return string(res)
}
// test
assertEq(reorganizeString("aab"), "aba")
assertEq(reorganizeString("aaab"), "")
```

## 502. IPO | Hard
related: 630
tags: greedy, two-heaps

Sort projects by capital. Each round, push every newly affordable project's profit into a max-heap and take the best one. Repeat k times.

**Complexity:** O(n log n + k log n).

```python
class Solution:
    def findMaximizedCapital(self, k: int, w: int, profits: List[int], capital: List[int]) -> int:
        projects = sorted(zip(capital, profits))
        h, i = [], 0
        for _ in range(k):
            while i < len(projects) and projects[i][0] <= w:
                heappush(h, -projects[i][1])
                i += 1
            if not h:
                break
            w -= heappop(h)
        return w
# test
assert Solution().findMaximizedCapital(2, 0, [1, 2, 3], [0, 1, 1]) == 4
assert Solution().findMaximizedCapital(3, 0, [1, 2, 3], [0, 1, 2]) == 6
```

```java
class Solution {
    public int findMaximizedCapital(int k, int w, int[] profits, int[] capital) {
        int n = profits.length;
        Integer[] idx = new Integer[n];
        for (int i = 0; i < n; i++) idx[i] = i;
        Arrays.sort(idx, (a, b) -> capital[a] - capital[b]);
        PriorityQueue<Integer> pq = new PriorityQueue<>(Collections.reverseOrder());
        int i = 0;
        while (k-- > 0) {
            while (i < n && capital[idx[i]] <= w) pq.add(profits[idx[i++]]);
            if (pq.isEmpty()) break;
            w += pq.poll();
        }
        return w;
    }
}
// test
check(new Solution().findMaximizedCapital(3, 0, new int[]{1, 2, 3}, new int[]{0, 1, 2}), 6);
```

```go
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	idx := make([]int, len(profits))
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(a, b int) bool { return capital[idx[a]] < capital[idx[b]] })
	h, i := &MaxHeap{}, 0
	for ; k > 0; k-- {
		for i < len(idx) && capital[idx[i]] <= w { // newly affordable projects
			heap.Push(h, profits[idx[i]])
			i++
		}
		if h.Len() == 0 {
			break
		}
		w += heap.Pop(h).(int) // take the most profitable one
	}
	return w
}
// test
assertEq(findMaximizedCapital(2, 0, []int{1, 2, 3}, []int{0, 1, 1}), 4)
assertEq(findMaximizedCapital(3, 0, []int{1, 2, 3}, []int{0, 1, 2}), 6)
```

## 373. Find K Pairs with Smallest Sums | Medium
related: 378, 23
tags: k-way-merge

Think of an implicit sorted matrix `nums1[i] + nums2[j]`. Seed the heap with `(i, 0)` for the first k rows; each pop pushes `(i, j+1)`.

**Complexity:** O(k log k).

```python
class Solution:
    def kSmallestPairs(self, nums1: List[int], nums2: List[int], k: int) -> List[List[int]]:
        h = [(nums1[i] + nums2[0], i, 0) for i in range(min(k, len(nums1)))]
        heapify(h)
        res = []
        while h and len(res) < k:
            _, i, j = heappop(h)
            res.append([nums1[i], nums2[j]])
            if j + 1 < len(nums2):
                heappush(h, (nums1[i] + nums2[j + 1], i, j + 1))
        return res
# test
assert sorted(Solution().kSmallestPairs([1, 7, 11], [2, 4, 6], 3)) == sorted([[1, 2], [1, 4], [1, 6]])
assert sorted(Solution().kSmallestPairs([1, 1, 2], [1, 2, 3], 2)) == sorted([[1, 1], [1, 1]])
```

```java
class Solution {
    public List<List<Integer>> kSmallestPairs(int[] nums1, int[] nums2, int k) {
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> a[0] - b[0]); // {sum, i, j}
        for (int i = 0; i < Math.min(k, nums1.length); i++) pq.add(new int[]{nums1[i] + nums2[0], i, 0});
        List<List<Integer>> res = new ArrayList<>();
        while (!pq.isEmpty() && res.size() < k) {
            int[] t = pq.poll();
            int i = t[1], j = t[2];
            res.add(List.of(nums1[i], nums2[j]));
            if (j + 1 < nums2.length) pq.add(new int[]{nums1[i] + nums2[j + 1], i, j + 1});
        }
        return res;
    }
}
// test
check(new Solution().kSmallestPairs(new int[]{1, 7, 11}, new int[]{2, 4, 6}, 3), "[[1, 2], [1, 4], [1, 6]]");
```

```go
type pair struct{ sum, i, j int }
type pairHeap []pair

func (h pairHeap) Len() int           { return len(h) }
func (h pairHeap) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h pairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *pairHeap) Push(x any)        { *h = append(*h, x.(pair)) }
func (h *pairHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := &pairHeap{}
	for i := 0; i < min(k, len(nums1)); i++ { // (i, 0) for every row
		heap.Push(h, pair{nums1[i] + nums2[0], i, 0})
	}
	res := [][]int{}
	for h.Len() > 0 && len(res) < k {
		p := heap.Pop(h).(pair)
		res = append(res, []int{nums1[p.i], nums2[p.j]})
		if p.j+1 < len(nums2) {
			heap.Push(h, pair{nums1[p.i] + nums2[p.j+1], p.i, p.j + 1})
		}
	}
	return res
}
// test
assertUnordered(kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3), [][]int{{1, 2}, {1, 4}, {1, 6}})
assertUnordered(kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 2), [][]int{{1, 1}, {1, 1}})
```

## 692. Top K Frequent Words | Medium
related: 347
tags: heap, sort

Count, then order by `(−count, word)`. A size-k heap needs the *reverse* order (min count, then larger word evicted first).

**Complexity:** O(n log k).

```python
class Solution:
    def topKFrequent(self, words: List[str], k: int) -> List[str]:
        cnt = Counter(words)
        return heapq.nsmallest(k, cnt, key=lambda w: (-cnt[w], w))
# test
assert Solution().topKFrequent(["i", "love", "leetcode", "i", "love", "coding"], 2) == ["i", "love"]
assert Solution().topKFrequent(["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], 4) == ["the", "is", "sunny", "day"]
```

```java
class Solution {
    public List<String> topKFrequent(String[] words, int k) {
        Map<String, Integer> cnt = new HashMap<>();
        for (String w : words) cnt.merge(w, 1, Integer::sum);
        // min-heap whose top is the "worst": lowest count, then lexicographically largest
        PriorityQueue<String> pq = new PriorityQueue<>((a, b) ->
                cnt.get(a).equals(cnt.get(b)) ? b.compareTo(a) : cnt.get(a) - cnt.get(b));
        for (String w : cnt.keySet()) {
            pq.add(w);
            if (pq.size() > k) pq.poll();
        }
        LinkedList<String> res = new LinkedList<>();
        while (!pq.isEmpty()) res.addFirst(pq.poll());
        return res;
    }
}
// test
check(new Solution().topKFrequent(new String[]{"i", "love", "leetcode", "i", "love", "coding"}, 2), "[i, love]");
```

```go
func topKFrequent(words []string, k int) []string {
	cnt := map[string]int{}
	for _, w := range words {
		cnt[w]++
	}
	keys := make([]string, 0, len(cnt))
	for w := range cnt {
		keys = append(keys, w)
	}
	sort.Slice(keys, func(i, j int) bool { // by count desc, then alphabetically
		a, b := keys[i], keys[j]
		if cnt[a] != cnt[b] {
			return cnt[a] > cnt[b]
		}
		return a < b
	})
	return keys[:k]
}
// test
assertEq(topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2), []string{"i", "love"})
assertEq(topKFrequent([]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4), []string{"the", "is", "sunny", "day"})
```


## 451. Sort Characters By Frequency | Medium
related: 347, 692
tags: bucket-sort, counting

Count characters, then emit them from most to least frequent (bucket sort by count, or sort the ≤ 62 distinct chars).

**Complexity:** O(n).

```python
class Solution:
    def frequencySort(self, s: str) -> str:
        return "".join(ch * c for ch, c in Counter(s).most_common())
# test
assert Solution().frequencySort("tree") in ("eert", "eetr")
assert Solution().frequencySort("Aabb") in ("bbAa", "bbaA")
```

```java
class Solution {
    public String frequencySort(String s) {
        Map<Character, Integer> cnt = new HashMap<>();
        for (char c : s.toCharArray()) cnt.merge(c, 1, Integer::sum);
        List<Character>[] bucket = new List[s.length() + 1];
        for (Map.Entry<Character, Integer> e : cnt.entrySet()) {
            if (bucket[e.getValue()] == null) bucket[e.getValue()] = new ArrayList<>();
            bucket[e.getValue()].add(e.getKey());
        }
        StringBuilder sb = new StringBuilder();
        for (int f = s.length(); f > 0; f--)
            if (bucket[f] != null) for (char c : bucket[f]) sb.append(String.valueOf(c).repeat(f));
        return sb.toString();
    }
}
// test
String r = new Solution().frequencySort("tree");
check(r.equals("eert") || r.equals("eetr"), true);
```

```go
func frequencySort(s string) string {
	cnt := map[byte]int{}
	for i := 0; i < len(s); i++ {
		cnt[s[i]]++
	}
	chars := make([]byte, 0, len(cnt))
	for c := range cnt {
		chars = append(chars, c)
	}
	sort.Slice(chars, func(i, j int) bool { return cnt[chars[i]] > cnt[chars[j]] })
	var sb strings.Builder
	for _, c := range chars {
		sb.WriteString(strings.Repeat(string(c), cnt[c]))
	}
	return sb.String()
}
// test
r := frequencySort("tree")
assertEq(r == "eert" || r == "eetr", true)
r2 := frequencySort("Aabb")
assertEq(r2 == "bbAa" || r2 == "bbaA", true)
```

## 658. Find K Closest Elements | Medium
related: 973, 35
tags: binary-search, fixed-window

The answer is a window `arr[i .. i+k)`. Binary search its start `i ∈ [0, n−k]`: if `x − arr[mid] > arr[mid+k] − x`, the window should move right.

**Complexity:** O(log(n − k) + k).

```python
class Solution:
    def findClosestElements(self, arr: List[int], k: int, x: int) -> List[int]:
        lo, hi = 0, len(arr) - k
        while lo < hi:
            mid = (lo + hi) // 2
            if x - arr[mid] > arr[mid + k] - x:
                lo = mid + 1
            else:
                hi = mid
        return arr[lo:lo + k]
# test
assert Solution().findClosestElements([1, 2, 3, 4, 5], 4, 3) == [1, 2, 3, 4]
assert Solution().findClosestElements([1, 1, 2, 3, 4, 5], 4, -1) == [1, 1, 2, 3]
assert Solution().findClosestElements([1, 3], 1, 2) == [1]
```

```java
class Solution {
    public List<Integer> findClosestElements(int[] arr, int k, int x) {
        int lo = 0, hi = arr.length - k;
        while (lo < hi) {
            int mid = (lo + hi) / 2;
            if (x - arr[mid] > arr[mid + k] - x) lo = mid + 1; else hi = mid;
        }
        List<Integer> res = new ArrayList<>();
        for (int i = lo; i < lo + k; i++) res.add(arr[i]);
        return res;
    }
}
// test
check(new Solution().findClosestElements(new int[]{1, 2, 3, 4, 5}, 4, 3), "[1, 2, 3, 4]");
```

```go
func findClosestElements(arr []int, k int, x int) []int {
	lo, hi := 0, len(arr)-k // left end of the window
	for lo < hi {
		mid := (lo + hi) / 2
		if x-arr[mid] > arr[mid+k]-x { // arr[mid+k] is closer: move right
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return arr[lo : lo+k]
}
// test
assertEq(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3), []int{1, 2, 3, 4})
assertEq(findClosestElements([]int{1, 1, 2, 3, 4, 5}, 4, -1), []int{1, 1, 2, 3})
assertEq(findClosestElements([]int{1, 3}, 1, 2), []int{1})
```

## 1642. Furthest Building You Can Reach | Medium
related: 871, 630
tags: greedy, min-heap

Use ladders for the **largest** climbs. Push every climb into a min-heap; when the heap exceeds `ladders`, the smallest climb gets paid with bricks instead. Stop when bricks go negative.

**Complexity:** O(n log ladders).

```python
class Solution:
    def furthestBuilding(self, heights: List[int], bricks: int, ladders: int) -> int:
        h = []
        for i in range(len(heights) - 1):
            d = heights[i + 1] - heights[i]
            if d <= 0:
                continue
            heappush(h, d)
            if len(h) > ladders:
                bricks -= heappop(h)
                if bricks < 0:
                    return i
        return len(heights) - 1
# test
assert Solution().furthestBuilding([4, 2, 7, 6, 9, 14, 12], 5, 1) == 4
assert Solution().furthestBuilding([4, 12, 2, 7, 3, 18, 20, 3, 19], 10, 2) == 7
assert Solution().furthestBuilding([14, 3, 19, 3], 17, 0) == 3
```

```java
class Solution {
    public int furthestBuilding(int[] heights, int bricks, int ladders) {
        PriorityQueue<Integer> pq = new PriorityQueue<>();
        for (int i = 0; i < heights.length - 1; i++) {
            int d = heights[i + 1] - heights[i];
            if (d <= 0) continue;
            pq.add(d);
            if (pq.size() > ladders && (bricks -= pq.poll()) < 0) return i;
        }
        return heights.length - 1;
    }
}
// test
check(new Solution().furthestBuilding(new int[]{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2), 7);
```

```go
type IntHeap []int // min-heap

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	h := &IntHeap{} // climbs covered by ladders: the largest ones
	for i := 0; i+1 < len(heights); i++ {
		d := heights[i+1] - heights[i]
		if d <= 0 {
			continue
		}
		heap.Push(h, d)
		if h.Len() > ladders { // the smallest climb pays with bricks
			bricks -= heap.Pop(h).(int)
			if bricks < 0 {
				return i
			}
		}
	}
	return len(heights) - 1
}
// test
assertEq(furthestBuilding([]int{4, 2, 7, 6, 9, 14, 12}, 5, 1), 4)
assertEq(furthestBuilding([]int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2), 7)
assertEq(furthestBuilding([]int{14, 3, 19, 3}, 17, 0), 3)
```

## 871. Minimum Number of Refueling Stops | Hard
related: 1642, 134
tags: greedy, max-heap

Drive as far as the fuel allows, pushing every station passed into a max-heap. When stuck, retroactively refuel at the **biggest** station passed (one stop). Fail if the heap is empty.

**Complexity:** O(n log n).

```python
class Solution:
    def minRefuelStops(self, target: int, startFuel: int, stations: List[List[int]]) -> int:
        h, reach, i, stops = [], startFuel, 0, 0
        while reach < target:
            while i < len(stations) and stations[i][0] <= reach:
                heappush(h, -stations[i][1])
                i += 1
            if not h:
                return -1
            reach -= heappop(h)
            stops += 1
        return stops
# test
assert Solution().minRefuelStops(1, 1, []) == 0
assert Solution().minRefuelStops(100, 1, [[10, 100]]) == -1
assert Solution().minRefuelStops(100, 10, [[10, 60], [20, 30], [30, 30], [60, 40]]) == 2
```

```java
class Solution {
    public int minRefuelStops(int target, int startFuel, int[][] stations) {
        PriorityQueue<Integer> pq = new PriorityQueue<>(Collections.reverseOrder());
        long reach = startFuel;
        int i = 0, stops = 0;
        while (reach < target) {
            while (i < stations.length && stations[i][0] <= reach) pq.add(stations[i++][1]);
            if (pq.isEmpty()) return -1;
            reach += pq.poll();
            stops++;
        }
        return stops;
    }
}
// test
check(new Solution().minRefuelStops(100, 10, new int[][]{{10, 60}, {20, 30}, {30, 30}, {60, 40}}), 2);
```

```go
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	h, reach, i, stops := &MaxHeap{}, startFuel, 0, 0
	for reach < target {
		for i < len(stations) && stations[i][0] <= reach { // stations we could have stopped at
			heap.Push(h, stations[i][1])
			i++
		}
		if h.Len() == 0 {
			return -1
		}
		reach += heap.Pop(h).(int) // retroactively stop at the biggest one
		stops++
	}
	return stops
}
// test
assertEq(minRefuelStops(1, 1, [][]int{}), 0)
assertEq(minRefuelStops(100, 1, [][]int{{10, 100}}), -1)
assertEq(minRefuelStops(100, 10, [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}), 2)
```

## 630. Course Schedule III | Hard
related: 1642, 502
tags: greedy, max-heap

Sort by deadline and take courses greedily, tracking total time. If the deadline is exceeded, drop the **longest** course taken so far (max-heap): the count stays the same but the time used shrinks.

**Complexity:** O(n log n).

```python
class Solution:
    def scheduleCourse(self, courses: List[List[int]]) -> int:
        h, t = [], 0
        for dur, last in sorted(courses, key=lambda c: c[1]):
            heappush(h, -dur)
            t += dur
            if t > last:
                t += heappop(h)          # remove the longest (stored negative)
        return len(h)
# test
assert Solution().scheduleCourse([[100, 200], [200, 1300], [1000, 1250], [2000, 3200]]) == 3
assert Solution().scheduleCourse([[1, 2]]) == 1
assert Solution().scheduleCourse([[3, 2], [4, 3]]) == 0
```

```java
class Solution {
    public int scheduleCourse(int[][] courses) {
        Arrays.sort(courses, (a, b) -> a[1] - b[1]);
        PriorityQueue<Integer> pq = new PriorityQueue<>(Collections.reverseOrder());
        int t = 0;
        for (int[] c : courses) {
            pq.add(c[0]);
            t += c[0];
            if (t > c[1]) t -= pq.poll();
        }
        return pq.size();
    }
}
// test
check(new Solution().scheduleCourse(new int[][]{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}), 3);
```

```go
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool { return courses[i][1] < courses[j][1] }) // by deadline
	h, t := &MaxHeap{}, 0
	for _, c := range courses {
		heap.Push(h, c[0])
		t += c[0]
		if t > c[1] { // over the deadline: drop the longest course taken
			t -= heap.Pop(h).(int)
		}
	}
	return h.Len()
}
// test
assertEq(scheduleCourse([][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}), 3)
assertEq(scheduleCourse([][]int{{1, 2}}), 1)
assertEq(scheduleCourse([][]int{{3, 2}, {4, 3}}), 0)
```

## 1834. Single-Threaded CPU | Medium
related: 621
tags: simulation, min-heap

Sort tasks by enqueue time. At each step, push all tasks that have arrived by `time` into a heap keyed by `(processing time, index)`; run the top. If the heap is empty, jump `time` to the next arrival.

**Complexity:** O(n log n).

```python
class Solution:
    def getOrder(self, tasks: List[List[int]]) -> List[int]:
        order = sorted(range(len(tasks)), key=lambda i: tasks[i][0])
        h, res, time, i = [], [], 0, 0
        while len(res) < len(tasks):
            if not h and time < tasks[order[i]][0]:
                time = tasks[order[i]][0]
            while i < len(order) and tasks[order[i]][0] <= time:
                j = order[i]
                heappush(h, (tasks[j][1], j))
                i += 1
            p, j = heappop(h)
            time += p
            res.append(j)
        return res
# test
assert Solution().getOrder([[1, 2], [2, 4], [3, 2], [4, 1]]) == [0, 2, 3, 1]
assert Solution().getOrder([[7, 10], [7, 12], [7, 5], [7, 4], [7, 2]]) == [4, 3, 2, 0, 1]
```

```java
class Solution {
    public int[] getOrder(int[][] tasks) {
        int n = tasks.length;
        Integer[] order = new Integer[n];
        for (int i = 0; i < n; i++) order[i] = i;
        Arrays.sort(order, (a, b) -> tasks[a][0] - tasks[b][0]);
        PriorityQueue<Integer> pq = new PriorityQueue<>((a, b) -> tasks[a][1] != tasks[b][1] ? tasks[a][1] - tasks[b][1] : a - b);
        int[] res = new int[n];
        long time = 0;
        for (int k = 0, i = 0; k < n; k++) {
            if (pq.isEmpty() && time < tasks[order[i]][0]) time = tasks[order[i]][0];
            while (i < n && tasks[order[i]][0] <= time) pq.add(order[i++]);
            int j = pq.poll();
            time += tasks[j][1];
            res[k] = j;
        }
        return res;
    }
}
// test
check(new Solution().getOrder(new int[][]{{1, 2}, {2, 4}, {3, 2}, {4, 1}}), new int[]{0, 2, 3, 1});
```

```go
type job struct{ proc, idx int }
type jobHeap []job // shortest processing time, then smallest index

func (h jobHeap) Len() int { return len(h) }
func (h jobHeap) Less(i, j int) bool {
	if h[i].proc != h[j].proc {
		return h[i].proc < h[j].proc
	}
	return h[i].idx < h[j].idx
}
func (h jobHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *jobHeap) Push(x any)   { *h = append(*h, x.(job)) }
func (h *jobHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func getOrder(tasks [][]int) []int {
	order := make([]int, len(tasks))
	for i := range order {
		order[i] = i
	}
	sort.Slice(order, func(a, b int) bool { return tasks[order[a]][0] < tasks[order[b]][0] })
	h, res, time, i := &jobHeap{}, []int{}, 0, 0
	for len(res) < len(tasks) {
		if h.Len() == 0 && time < tasks[order[i]][0] { // idle until the next task arrives
			time = tasks[order[i]][0]
		}
		for i < len(order) && tasks[order[i]][0] <= time {
			heap.Push(h, job{tasks[order[i]][1], order[i]})
			i++
		}
		j := heap.Pop(h).(job)
		time += j.proc
		res = append(res, j.idx)
	}
	return res
}
// test
assertEq(getOrder([][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}}), []int{0, 2, 3, 1})
assertEq(getOrder([][]int{{7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2}}), []int{4, 3, 2, 0, 1})
```

## 480. Sliding Window Median | Hard
related: 295, 239
tags: sorted-window, two-heaps

Maintain the current window **sorted**: insert the new value and delete the outgoing one with binary search; the median is the middle element(s). Optimal alternative: the two-heap median finder (295) with lazy deletion, or two balanced `TreeSet`s (Java, below), for O(n log k).

**Complexity:** O(n log k) with balanced trees; O(n·k) with a plain sorted list (fine in practice for Python).

```python
class Solution:
    def medianSlidingWindow(self, nums: List[int], k: int) -> List[float]:
        win = sorted(nums[:k])
        res = []
        for i in range(k, len(nums) + 1):
            res.append((win[k // 2] + win[(k - 1) // 2]) / 2)
            if i == len(nums):
                break
            win.pop(bisect.bisect_left(win, nums[i - k]))
            bisect.insort(win, nums[i])
        return res
# test
assert Solution().medianSlidingWindow([1, 3, -1, -3, 5, 3, 6, 7], 3) == [1.0, -1.0, -1.0, 3.0, 5.0, 6.0]
assert Solution().medianSlidingWindow([1, 2, 3, 4, 2, 3, 1, 4, 2], 3) == [2.0, 3.0, 3.0, 3.0, 2.0, 3.0, 2.0]
```

```java
class Solution {
    public double[] medianSlidingWindow(int[] nums, int k) {
        // indices ordered by (value, index) so duplicates are distinct; lo holds the smaller half (size ceil(k/2))
        Comparator<Integer> cmp = (a, b) -> nums[a] != nums[b] ? Integer.compare(nums[a], nums[b]) : a - b;
        TreeSet<Integer> lo = new TreeSet<>(cmp), hi = new TreeSet<>(cmp);
        double[] res = new double[nums.length - k + 1];
        for (int i = 0; i < nums.length; i++) {
            if (i >= k && !lo.remove(i - k)) hi.remove(i - k);
            lo.add(i);
            hi.add(lo.pollLast());                       // keep every lo <= every hi
            if (hi.size() > lo.size()) lo.add(hi.pollFirst());
            if (i >= k - 1)
                res[i - k + 1] = k % 2 == 1 ? nums[lo.last()] : ((double) nums[lo.last()] + nums[hi.first()]) / 2;
        }
        return res;
    }
}
// test
check(new Solution().medianSlidingWindow(new int[]{1, 3, -1, -3, 5, 3, 6, 7}, 3), new double[]{1.0, -1.0, -1.0, 3.0, 5.0, 6.0});
check(new Solution().medianSlidingWindow(new int[]{1, 4, 2, 3}, 4), new double[]{2.5});
```

```go
func medianSlidingWindow(nums []int, k int) []float64 {
	win := append([]int{}, nums[:k]...) // the window, kept sorted
	sort.Ints(win)
	res := []float64{}
	for i := k; ; i++ {
		res = append(res, (float64(win[k/2])+float64(win[(k-1)/2]))/2)
		if i == len(nums) {
			return res
		}
		j := sort.SearchInts(win, nums[i-k]) // remove the outgoing value
		win = append(win[:j], win[j+1:]...)
		j = sort.SearchInts(win, nums[i]) // insert the incoming one
		win = append(win, 0)
		copy(win[j+1:], win[j:])
		win[j] = nums[i]
	}
}
// test
assertEq(medianSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3), []float64{1, -1, -1, 3, 5, 6})
assertEq(medianSlidingWindow([]int{1, 2, 3, 4, 2, 3, 1, 4, 2}, 3), []float64{2, 3, 3, 3, 2, 3, 2})
```

## 632. Smallest Range Covering Elements from K Lists | Hard
related: 23, 76
tags: k-way-merge, min-heap

K-way merge with a min-heap holding one element from each list, plus the current **maximum** of those elements. The heap top and the max form a range containing one element of every list. Advance the list that holds the minimum (the only move that can shrink the range); stop when that list runs out.

**Complexity:** O(N log k) for N total elements.

```python
class Solution:
    def smallestRange(self, nums: List[List[int]]) -> List[int]:
        heap = [(row[0], i, 0) for i, row in enumerate(nums)]
        heapify(heap)
        hi = max(row[0] for row in nums)
        best = [heap[0][0], hi]
        while True:
            lo, i, j = heappop(heap)
            if hi - lo < best[1] - best[0]:
                best = [lo, hi]
            if j + 1 == len(nums[i]):
                return best
            nxt = nums[i][j + 1]
            hi = max(hi, nxt)
            heappush(heap, (nxt, i, j + 1))
# test
assert Solution().smallestRange([[4, 10, 15, 24, 26], [0, 9, 12, 20], [5, 18, 22, 30]]) == [20, 24]
assert Solution().smallestRange([[1, 2, 3], [1, 2, 3], [1, 2, 3]]) == [1, 1]
assert Solution().smallestRange([[1], [2], [3]]) == [1, 3]
assert Solution().smallestRange([[10, 10], [11, 11]]) == [10, 11]
```

```java
class Solution {
    public int[] smallestRange(List<List<Integer>> nums) {
        PriorityQueue<int[]> pq = new PriorityQueue<>((a, b) -> Integer.compare(a[0], b[0])); // {value, list, index}
        int hi = Integer.MIN_VALUE;
        for (int i = 0; i < nums.size(); i++) {
            pq.add(new int[]{nums.get(i).get(0), i, 0});
            hi = Math.max(hi, nums.get(i).get(0));
        }
        int[] best = {pq.peek()[0], hi};
        while (true) {
            int[] top = pq.poll();
            if (hi - top[0] < best[1] - best[0]) best = new int[]{top[0], hi};
            List<Integer> row = nums.get(top[1]);
            if (top[2] + 1 == row.size()) return best;
            int nxt = row.get(top[2] + 1);
            hi = Math.max(hi, nxt);
            pq.add(new int[]{nxt, top[1], top[2] + 1});
        }
    }
}
// test
check(new Solution().smallestRange(List.of(List.of(4, 10, 15, 24, 26), List.of(0, 9, 12, 20), List.of(5, 18, 22, 30))), new int[]{20, 24});
check(new Solution().smallestRange(List.of(List.of(1, 2, 3), List.of(1, 2, 3), List.of(1, 2, 3))), new int[]{1, 1});
```

```go
type entry struct{ v, i, j int } // value, list, position
type entryHeap []entry

func (h entryHeap) Len() int           { return len(h) }
func (h entryHeap) Less(i, j int) bool { return h[i].v < h[j].v }
func (h entryHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *entryHeap) Push(x any)        { *h = append(*h, x.(entry)) }
func (h *entryHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func smallestRange(nums [][]int) []int {
	h, hi := &entryHeap{}, math.MinInt
	for i, row := range nums { // one pointer per list; the range is [heap min, hi]
		heap.Push(h, entry{row[0], i, 0})
		hi = max(hi, row[0])
	}
	best := []int{(*h)[0].v, hi}
	for {
		e := heap.Pop(h).(entry)
		if hi-e.v < best[1]-best[0] {
			best = []int{e.v, hi}
		}
		if e.j+1 == len(nums[e.i]) { // a list is exhausted
			return best
		}
		next := nums[e.i][e.j+1]
		hi = max(hi, next)
		heap.Push(h, entry{next, e.i, e.j + 1})
	}
}
// test
assertEq(smallestRange([][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}), []int{20, 24})
assertEq(smallestRange([][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}), []int{1, 1})
assertEq(smallestRange([][]int{{1}, {2}, {3}}), []int{1, 3})
assertEq(smallestRange([][]int{{10, 10}, {11, 11}}), []int{10, 11})
```

## 2336. Smallest Number in Infinite Set | Medium
related: 1845, 703
tags: min-heap, design

Every integer `≥ nxt` is still present — no need to store them. Numbers added back **below** `nxt` go into a min-heap (with a set to avoid duplicates). Pop from the heap if it's non-empty, otherwise hand out `nxt` and increment it.

**Complexity:** O(log n) per operation.

```python
class SmallestInfiniteSet:
    def __init__(self):
        self.nxt = 1
        self.back, self.in_back = [], set()

    def popSmallest(self) -> int:
        if self.back:
            x = heappop(self.back)
            self.in_back.remove(x)
            return x
        self.nxt += 1
        return self.nxt - 1

    def addBack(self, num: int) -> None:
        if num < self.nxt and num not in self.in_back:
            self.in_back.add(num)
            heappush(self.back, num)
# test
s = SmallestInfiniteSet()
s.addBack(2)
assert [s.popSmallest(), s.popSmallest(), s.popSmallest()] == [1, 2, 3]
s.addBack(1)
assert [s.popSmallest(), s.popSmallest(), s.popSmallest()] == [1, 4, 5]
s2 = SmallestInfiniteSet()
s2.popSmallest(); s2.popSmallest()
s2.addBack(1); s2.addBack(1)
assert [s2.popSmallest(), s2.popSmallest()] == [1, 3]
```

```java
class SmallestInfiniteSet {
    private int nxt = 1;
    private final TreeSet<Integer> back = new TreeSet<>();   // re-added numbers below nxt

    public int popSmallest() { return back.isEmpty() ? nxt++ : back.pollFirst(); }

    public void addBack(int num) { if (num < nxt) back.add(num); }
}
// test
SmallestInfiniteSet s = new SmallestInfiniteSet();
s.addBack(2);
check(new int[]{s.popSmallest(), s.popSmallest(), s.popSmallest()}, new int[]{1, 2, 3});
s.addBack(1);
check(new int[]{s.popSmallest(), s.popSmallest(), s.popSmallest()}, new int[]{1, 4, 5});
```

```go
type IntHeap []int // min-heap

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

type SmallestInfiniteSet struct {
	next   int          // every number >= next is in the set
	back   IntHeap      // numbers < next that were added back
	inBack map[int]bool // ...and their membership
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{next: 1, inBack: map[int]bool{}}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.back.Len() > 0 {
		x := heap.Pop(&this.back).(int)
		delete(this.inBack, x)
		return x
	}
	this.next++
	return this.next - 1
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num < this.next && !this.inBack[num] {
		this.inBack[num] = true
		heap.Push(&this.back, num)
	}
}
// test
s := Constructor()
s.AddBack(2)
assertEq([]int{s.PopSmallest(), s.PopSmallest(), s.PopSmallest()}, []int{1, 2, 3})
s.AddBack(1)
assertEq([]int{s.PopSmallest(), s.PopSmallest(), s.PopSmallest()}, []int{1, 4, 5})
s2 := Constructor()
s2.PopSmallest()
s2.PopSmallest()
s2.AddBack(1)
s2.AddBack(1)
assertEq([]int{s2.PopSmallest(), s2.PopSmallest()}, []int{1, 3})
```

## 2542. Maximum Subsequence Score | Medium
related: 857, 1383
tags: min-heap, sort, greedy

Sort pairs by `nums2` **descending**. When we reach pair `i`, `nums2[i]` is the minimum of everything seen so far — so fix it as the multiplier and pair it with the **k largest `nums1`** among the seen pairs, kept in a size-k min-heap with a running sum.

**Complexity:** O(n log n).

```python
class Solution:
    def maxScore(self, nums1: List[int], nums2: List[int], k: int) -> int:
        heap, total, best = [], 0, 0
        for b, a in sorted(zip(nums2, nums1), reverse=True):
            heappush(heap, a)
            total += a
            if len(heap) > k:
                total -= heappop(heap)
            if len(heap) == k:
                best = max(best, total * b)
        return best
# test
assert Solution().maxScore([1, 3, 3, 2], [2, 1, 3, 4], 3) == 12
assert Solution().maxScore([4, 2, 3, 1, 1], [7, 5, 10, 9, 6], 1) == 30
assert Solution().maxScore([2, 1, 14, 12], [11, 7, 13, 6], 3) == 168
```

```java
class Solution {
    public long maxScore(int[] nums1, int[] nums2, int k) {
        int n = nums1.length;
        Integer[] idx = new Integer[n];
        for (int i = 0; i < n; i++) idx[i] = i;
        Arrays.sort(idx, (x, y) -> nums2[y] - nums2[x]);
        PriorityQueue<Integer> heap = new PriorityQueue<>();
        long total = 0, best = 0;
        for (int i : idx) {
            heap.add(nums1[i]);
            total += nums1[i];
            if (heap.size() > k) total -= heap.poll();
            if (heap.size() == k) best = Math.max(best, total * nums2[i]);
        }
        return best;
    }
}
// test
check(new Solution().maxScore(new int[]{1, 3, 3, 2}, new int[]{2, 1, 3, 4}, 3), 12L);
check(new Solution().maxScore(new int[]{4, 2, 3, 1, 1}, new int[]{7, 5, 10, 9, 6}, 1), 30L);
```

```go
type IntHeap []int // min-heap

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	idx := make([]int, len(nums1))
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(a, b int) bool { return nums2[idx[a]] > nums2[idx[b]] }) // by nums2 desc
	h, total, best := &IntHeap{}, 0, 0
	for _, i := range idx { // nums2[i] is the minimum of the chosen set so far
		heap.Push(h, nums1[i])
		total += nums1[i]
		if h.Len() > k { // keep the k largest nums1 values
			total -= heap.Pop(h).(int)
		}
		if h.Len() == k {
			best = max(best, total*nums2[i])
		}
	}
	return int64(best)
}
// test
assertEq(maxScore([]int{1, 3, 3, 2}, []int{2, 1, 3, 4}, 3), 12)
assertEq(maxScore([]int{4, 2, 3, 1, 1}, []int{7, 5, 10, 9, 6}, 1), 30)
assertEq(maxScore([]int{2, 1, 14, 12}, []int{11, 7, 13, 6}, 3), 168)
```

## 2462. Total Cost to Hire K Workers | Medium
related: 2542, 1046
tags: two-heaps

Two min-heaps: the first `candidates` workers and the last `candidates` workers (never overlapping — the pointers `i ≤ j` guarantee each worker is in at most one heap). Each round hire the cheaper top (front wins ties: smaller index), then refill that side from the middle.

**Complexity:** O((k + candidates) log candidates).

```python
class Solution:
    def totalCost(self, costs: List[int], k: int, candidates: int) -> int:
        i, j = 0, len(costs) - 1
        front, back, total = [], [], 0
        for _ in range(k):
            while len(front) < candidates and i <= j:
                heappush(front, costs[i])
                i += 1
            while len(back) < candidates and i <= j:
                heappush(back, costs[j])
                j -= 1
            if not back or (front and front[0] <= back[0]):
                total += heappop(front)
            else:
                total += heappop(back)
        return total
# test
assert Solution().totalCost([17, 12, 10, 2, 7, 2, 11, 20, 8], 3, 4) == 11
assert Solution().totalCost([1, 2, 4, 1], 3, 3) == 4
assert Solution().totalCost([5], 1, 1) == 5
assert Solution().totalCost([2, 2, 2, 2], 2, 1) == 4
```

```java
class Solution {
    public long totalCost(int[] costs, int k, int candidates) {
        PriorityQueue<Integer> front = new PriorityQueue<>(), back = new PriorityQueue<>();
        int i = 0, j = costs.length - 1;
        long total = 0;
        while (k-- > 0) {
            while (front.size() < candidates && i <= j) front.add(costs[i++]);
            while (back.size() < candidates && i <= j) back.add(costs[j--]);
            if (back.isEmpty() || (!front.isEmpty() && front.peek() <= back.peek())) total += front.poll();
            else total += back.poll();
        }
        return total;
    }
}
// test
check(new Solution().totalCost(new int[]{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 4), 11L);
check(new Solution().totalCost(new int[]{1, 2, 4, 1}, 3, 3), 4L);
```

```go
type IntHeap []int // min-heap

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func totalCost(costs []int, k int, candidates int) int64 {
	i, j := 0, len(costs)-1
	front, back := &IntHeap{}, &IntHeap{}
	total := 0
	for ; k > 0; k-- {
		for front.Len() < candidates && i <= j {
			heap.Push(front, costs[i])
			i++
		}
		for back.Len() < candidates && i <= j {
			heap.Push(back, costs[j])
			j--
		}
		if back.Len() == 0 || (front.Len() > 0 && (*front)[0] <= (*back)[0]) { // ties go to the front
			total += heap.Pop(front).(int)
		} else {
			total += heap.Pop(back).(int)
		}
	}
	return int64(total)
}
// test
assertEq(totalCost([]int{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 4), 11)
assertEq(totalCost([]int{1, 2, 4, 1}, 3, 3), 4)
assertEq(totalCost([]int{5}, 1, 1), 5)
assertEq(totalCost([]int{2, 2, 2, 2}, 2, 1), 4)
```

## 1405. Longest Happy String | Medium
related: 767, 621
tags: max-heap, greedy

Greedy with a max-heap of remaining counts: append the most plentiful letter unless the last two characters are already that letter — then append the second most plentiful instead (and stop if there is none). Using the largest pile whenever allowed is what keeps the string as long as possible.

**Complexity:** O(a + b + c).

```python
class Solution:
    def longestDiverseString(self, a: int, b: int, c: int) -> str:
        heap = [(-n, ch) for n, ch in ((a, "a"), (b, "b"), (c, "c")) if n]
        heapify(heap)
        out = []
        while heap:
            n1, c1 = heappop(heap)
            if len(out) >= 2 and out[-1] == out[-2] == c1:
                if not heap:
                    break
                n2, c2 = heappop(heap)
                out.append(c2)
                if n2 + 1:
                    heappush(heap, (n2 + 1, c2))
                heappush(heap, (n1, c1))
            else:
                out.append(c1)
                if n1 + 1:
                    heappush(heap, (n1 + 1, c1))
        return "".join(out)
# test
def happy(s, a, b, c):  # any valid answer of maximum length is accepted
    x, y, z = sorted((a, b, c), reverse=True)
    return ("aaa" not in s and "bbb" not in s and "ccc" not in s and s.count("a") <= a and s.count("b") <= b
            and s.count("c") <= c and len(s) == y + z + min(x, 2 * (y + z + 1)))
for a, b, c in ((1, 1, 7), (7, 1, 0), (2, 2, 1), (0, 0, 0), (0, 8, 11), (5, 5, 5), (0, 0, 9)):
    assert happy(Solution().longestDiverseString(a, b, c), a, b, c)
```

```java
class Solution {
    public String longestDiverseString(int a, int b, int c) {
        PriorityQueue<int[]> pq = new PriorityQueue<>((x, y) -> y[0] - x[0]);   // {count, letter}
        if (a > 0) pq.add(new int[]{a, 'a'});
        if (b > 0) pq.add(new int[]{b, 'b'});
        if (c > 0) pq.add(new int[]{c, 'c'});
        StringBuilder sb = new StringBuilder();
        while (!pq.isEmpty()) {
            int[] top = pq.poll();
            int len = sb.length();
            if (len >= 2 && sb.charAt(len - 1) == top[1] && sb.charAt(len - 2) == top[1]) {
                if (pq.isEmpty()) break;
                int[] second = pq.poll();
                sb.append((char) second[1]);
                if (--second[0] > 0) pq.add(second);
                pq.add(top);
            } else {
                sb.append((char) top[1]);
                if (--top[0] > 0) pq.add(top);
            }
        }
        return sb.toString();
    }
}
// test
String hs = new Solution().longestDiverseString(1, 1, 7);
check(hs.length() == 8 && !hs.contains("ccc"), true);
check(new Solution().longestDiverseString(7, 1, 0).length(), 5);
```

```go
func longestDiverseString(a int, b int, c int) string {
	cnt := []int{a, b, c}
	out := []byte{}
	for {
		order := []int{0, 1, 2} // letters by remaining count, most first
		sort.Slice(order, func(i, j int) bool { return cnt[order[i]] > cnt[order[j]] })
		placed := false
		for _, x := range order {
			ch := byte('a' + x)
			n := len(out)
			if cnt[x] == 0 || (n >= 2 && out[n-1] == ch && out[n-2] == ch) {
				continue // would make three in a row
			}
			out = append(out, ch)
			cnt[x]--
			placed = true
			break
		}
		if !placed {
			return string(out)
		}
	}
}
// test
happy := func(s string, a, b, c int) bool { // any valid answer of maximum length is accepted
	v := []int{a, b, c}
	sort.Sort(sort.Reverse(sort.IntSlice(v)))
	x, y, z := v[0], v[1], v[2]
	return !strings.Contains(s, "aaa") && !strings.Contains(s, "bbb") && !strings.Contains(s, "ccc") &&
		strings.Count(s, "a") <= a && strings.Count(s, "b") <= b && strings.Count(s, "c") <= c &&
		len(s) == y+z+min(x, 2*(y+z+1))
}
for _, tc := range [][3]int{{1, 1, 7}, {7, 1, 0}, {2, 2, 1}, {0, 0, 0}, {0, 8, 11}, {5, 5, 5}, {0, 0, 9}} {
	assertEq(happy(longestDiverseString(tc[0], tc[1], tc[2]), tc[0], tc[1], tc[2]), true)
}
```
