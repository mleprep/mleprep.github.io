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
assert Solution().kSmallestPairs([1, 7, 11], [2, 4, 6], 3) == [[1, 2], [1, 4], [1, 6]]
assert Solution().kSmallestPairs([1, 1, 2], [1, 2, 3], 2) == [[1, 1], [1, 1]]
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
