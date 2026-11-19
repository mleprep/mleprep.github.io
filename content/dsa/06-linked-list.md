# Linked List

**Toolkit**
- **Dummy head** — removes "is this the head?" special cases when deleting/inserting.
- **Fast & slow pointers** — middle of list, cycle detection (Floyd), n-th from end.
- **In-place reversal** — `prev, cur = None, head; while cur: cur.next, prev, cur = prev, cur, cur.next`.
- Many hard problems = combination of the three: *find middle → reverse second half → merge/compare*.

```viz
reverse:   prev   cur    nxt
           None   1 ->   2 -> 3
           None <-1      2 -> 3       (cur.next = prev; prev = cur; cur = nxt)
```

Locally, Python tests use `make_list([..])` / `list_vals(node)`; Java tests use `ListNode.of(..)` (prints like `[1, 2, 3]`).

## 206. Reverse Linked List | Easy
related: 92, 25, 234
tags: blind75, reversal

Iteratively flip each `next` pointer while walking forward.

**Complexity:** O(n) time, O(1) space (recursive version uses O(n) stack).

```python
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        prev, cur = None, head
        while cur:
            cur.next, prev, cur = prev, cur, cur.next
        return prev
# test
assert list_vals(Solution().reverseList(make_list([1, 2, 3, 4, 5]))) == [5, 4, 3, 2, 1]
assert Solution().reverseList(None) is None
```

```java
class Solution {
    public ListNode reverseList(ListNode head) {
        ListNode prev = null, cur = head;
        while (cur != null) {
            ListNode nxt = cur.next;
            cur.next = prev;
            prev = cur;
            cur = nxt;
        }
        return prev;
    }
}
// test
check(new Solution().reverseList(ListNode.of(1, 2, 3, 4, 5)), ListNode.of(5, 4, 3, 2, 1));
```

## 21. Merge Two Sorted Lists | Easy
related: 23, 88, 148
tags: blind75, dummy-head

Dummy head + tail pointer; repeatedly attach the smaller head. Append the leftover list.

**Complexity:** O(m + n) time, O(1) space.

```python
class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        dummy = tail = ListNode()
        while list1 and list2:
            if list1.val <= list2.val:
                tail.next, list1 = list1, list1.next
            else:
                tail.next, list2 = list2, list2.next
            tail = tail.next
        tail.next = list1 or list2
        return dummy.next
# test
assert list_vals(Solution().mergeTwoLists(make_list([1, 2, 4]), make_list([1, 3, 4]))) == [1, 1, 2, 3, 4, 4]
assert list_vals(Solution().mergeTwoLists(make_list([]), make_list([]))) == []
assert list_vals(Solution().mergeTwoLists(make_list([]), make_list([0]))) == [0]
assert list_vals(Solution().mergeTwoLists(make_list([5]), make_list([1, 2, 4]))) == [1, 2, 4, 5]
```

```java
class Solution {
    public ListNode mergeTwoLists(ListNode a, ListNode b) {
        ListNode dummy = new ListNode(), tail = dummy;
        while (a != null && b != null) {
            if (a.val <= b.val) { tail.next = a; a = a.next; }
            else { tail.next = b; b = b.next; }
            tail = tail.next;
        }
        tail.next = a != null ? a : b;
        return dummy.next;
    }
}
// test
check(new Solution().mergeTwoLists(ListNode.of(1, 2, 4), ListNode.of(1, 3, 4)), ListNode.of(1, 1, 2, 3, 4, 4));
```

## 141. Linked List Cycle | Easy
related: 142, 287, 202
tags: blind75, fast-slow

Floyd: slow moves 1, fast moves 2. If there's a cycle they must meet; otherwise fast hits null.

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        slow = fast = head
        while fast and fast.next:
            slow, fast = slow.next, fast.next.next
            if slow is fast:
                return True
        return False
# test
h = make_list([3, 2, 0, -4]); h.next.next.next.next = h.next   # tail -> node 1
assert Solution().hasCycle(h) is True
assert Solution().hasCycle(make_list([1])) is False
```

```java
class Solution {
    public boolean hasCycle(ListNode head) {
        ListNode slow = head, fast = head;
        while (fast != null && fast.next != null) {
            slow = slow.next;
            fast = fast.next.next;
            if (slow == fast) return true;
        }
        return false;
    }
}
// test
ListNode h = ListNode.of(3, 2, 0, -4);
h.next.next.next.next = h.next;
check(new Solution().hasCycle(h), true);
check(new Solution().hasCycle(ListNode.of(1, 2)), false);
```

## 142. Linked List Cycle II | Medium
related: 141, 287
tags: fast-slow

After slow and fast meet, restart one pointer at `head`; moving both one step at a time, they meet at the cycle entrance. (Distance head→entry ≡ distance meet→entry mod cycle length.)

```viz
head -> a -> b -> [c] -> d -> e
                   ^           |
                   +-----------+     x = head..c, meet inside; x ≡ meet..c
```

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def detectCycle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        slow = fast = head
        while fast and fast.next:
            slow, fast = slow.next, fast.next.next
            if slow is fast:
                p = head
                while p is not slow:
                    p, slow = p.next, slow.next
                return p
        return None
# test
h = make_list([3, 2, 0, -4]); h.next.next.next.next = h.next
assert Solution().detectCycle(h) is h.next
assert Solution().detectCycle(make_list([1])) is None
```

```java
class Solution {
    public ListNode detectCycle(ListNode head) {
        ListNode slow = head, fast = head;
        while (fast != null && fast.next != null) {
            slow = slow.next;
            fast = fast.next.next;
            if (slow == fast) {
                ListNode p = head;
                while (p != slow) { p = p.next; slow = slow.next; }
                return p;
            }
        }
        return null;
    }
}
// test
ListNode h = ListNode.of(3, 2, 0, -4);
h.next.next.next.next = h.next;
check(new Solution().detectCycle(h) == h.next, true);
```

## 143. Reorder List | Medium
related: 206, 234, 876
tags: blind75, fast-slow, reversal

`L0 → Ln → L1 → Ln-1 …` = **find middle**, **reverse second half**, **interleave** the halves.

```viz
1 2 3 4 5    middle at 3 -> split  1 2 3 | 5 4 (reversed)
merge:       1 5 2 4 3
```

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        slow = fast = head
        while fast.next and fast.next.next:
            slow, fast = slow.next, fast.next.next
        prev, cur = None, slow.next
        slow.next = None
        while cur:
            cur.next, prev, cur = prev, cur, cur.next
        a, b = head, prev
        while b:
            an, bn = a.next, b.next
            a.next, b.next = b, an
            a, b = an, bn
# test
h = make_list([1, 2, 3, 4, 5]); Solution().reorderList(h)
assert list_vals(h) == [1, 5, 2, 4, 3]
h = make_list([1, 2, 3, 4]); Solution().reorderList(h)
assert list_vals(h) == [1, 4, 2, 3]
```

```java
class Solution {
    public void reorderList(ListNode head) {
        ListNode slow = head, fast = head;
        while (fast.next != null && fast.next.next != null) { slow = slow.next; fast = fast.next.next; }
        ListNode prev = null, cur = slow.next;
        slow.next = null;
        while (cur != null) { ListNode nxt = cur.next; cur.next = prev; prev = cur; cur = nxt; }
        ListNode a = head, b = prev;
        while (b != null) {
            ListNode an = a.next, bn = b.next;
            a.next = b; b.next = an;
            a = an; b = bn;
        }
    }
}
// test
ListNode h = ListNode.of(1, 2, 3, 4, 5);
new Solution().reorderList(h);
check(h, ListNode.of(1, 5, 2, 4, 3));
```

## 19. Remove Nth Node From End of List | Medium
related: 876, 61
tags: blind75, two-pointers, dummy-head

Advance `fast` n steps ahead from a dummy, then move both until `fast` is last; `slow.next` is the node to delete.

```viz
dummy 1 2 3 4 5, n=2
fast moves 2 ahead; then both move until fast at 5:
slow at 3 -> skip 4 -> 1 2 3 5
```

**Complexity:** O(n) time, one pass.

```python
class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        dummy = ListNode(0, head)
        slow = fast = dummy
        for _ in range(n):
            fast = fast.next
        while fast.next:
            slow, fast = slow.next, fast.next
        slow.next = slow.next.next
        return dummy.next
# test
assert list_vals(Solution().removeNthFromEnd(make_list([1, 2, 3, 4, 5]), 2)) == [1, 2, 3, 5]
assert list_vals(Solution().removeNthFromEnd(make_list([1]), 1)) == []
```

```java
class Solution {
    public ListNode removeNthFromEnd(ListNode head, int n) {
        ListNode dummy = new ListNode(0, head), slow = dummy, fast = dummy;
        for (int i = 0; i < n; i++) fast = fast.next;
        while (fast.next != null) { slow = slow.next; fast = fast.next; }
        slow.next = slow.next.next;
        return dummy.next;
    }
}
// test
check(new Solution().removeNthFromEnd(ListNode.of(1, 2, 3, 4, 5), 2), ListNode.of(1, 2, 3, 5));
check(new Solution().removeNthFromEnd(ListNode.of(1), 1), null);
```

## 138. Copy List with Random Pointer | Medium
related: 133
tags: hash-map, interleave

Map each original node → its clone (first pass), then wire `next`/`random` through the map (second pass). O(1)-space variant: interleave clones `A A' B B' …`, set `A'.random = A.random.next`, then unweave.

**Complexity:** O(n) time, O(n) space.

```python
class Node:  # provided by LeetCode
    def __init__(self, x: int, next: "Node" = None, random: "Node" = None):
        self.val, self.next, self.random = int(x), next, random

class Solution:
    def copyRandomList(self, head: "Optional[Node]") -> "Optional[Node]":
        clone = {None: None}
        cur = head
        while cur:
            clone[cur] = Node(cur.val)
            cur = cur.next
        cur = head
        while cur:
            clone[cur].next = clone[cur.next]
            clone[cur].random = clone[cur.random]
            cur = cur.next
        return clone[head]
# test
spec = [[7, None], [13, 0], [11, 4], [10, 2], [1, 0]]
nodes = [Node(v) for v, _ in spec]
for i, (_, r) in enumerate(spec):
    nodes[i].next = nodes[i + 1] if i + 1 < len(nodes) else None
    nodes[i].random = nodes[r] if r is not None else None
c, out = Solution().copyRandomList(nodes[0]), []
copies = []
while c:
    assert c not in nodes
    copies.append(c); c = c.next
assert [[n.val, copies.index(n.random) if n.random else None] for n in copies] == spec
```

```java
class Node { // provided by LeetCode
    int val; Node next, random;
    Node(int val) { this.val = val; }
}

class Solution {
    public Node copyRandomList(Node head) {
        Map<Node, Node> clone = new HashMap<>();
        for (Node cur = head; cur != null; cur = cur.next) clone.put(cur, new Node(cur.val));
        for (Node cur = head; cur != null; cur = cur.next) {
            clone.get(cur).next = clone.get(cur.next);     // get(null) -> null
            clone.get(cur).random = clone.get(cur.random);
        }
        return clone.get(head);
    }
}
// test
Node a = new Node(7), b = new Node(13), c = new Node(11);
a.next = b; b.next = c; b.random = a; c.random = c;
Node x = new Solution().copyRandomList(a);
check(x != a && x.val == 7 && x.next.val == 13 && x.next.random == x && x.next.next.random == x.next.next, true);
```

## 2. Add Two Numbers | Medium
related: 445, 67
tags: dummy-head, carry

Digits are stored least-significant first, so add node by node with a carry, like grade-school addition.

```viz
  2 -> 4 -> 3     (342)
+ 5 -> 6 -> 4     (465)
= 7 -> 0 -> 8     (807)   carry: 0, 1, 0
```

**Complexity:** O(max(m, n)).

```python
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        dummy = tail = ListNode()
        carry = 0
        while l1 or l2 or carry:
            s = carry + (l1.val if l1 else 0) + (l2.val if l2 else 0)
            carry, d = divmod(s, 10)
            tail.next = ListNode(d)
            tail = tail.next
            l1 = l1.next if l1 else None
            l2 = l2.next if l2 else None
        return dummy.next
# test
assert list_vals(Solution().addTwoNumbers(make_list([2, 4, 3]), make_list([5, 6, 4]))) == [7, 0, 8]
assert list_vals(Solution().addTwoNumbers(make_list([9, 9, 9]), make_list([1]))) == [0, 0, 0, 1]
```

```java
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode dummy = new ListNode(), tail = dummy;
        int carry = 0;
        while (l1 != null || l2 != null || carry != 0) {
            int s = carry;
            if (l1 != null) { s += l1.val; l1 = l1.next; }
            if (l2 != null) { s += l2.val; l2 = l2.next; }
            carry = s / 10;
            tail = tail.next = new ListNode(s % 10);
        }
        return dummy.next;
    }
}
// test
check(new Solution().addTwoNumbers(ListNode.of(2, 4, 3), ListNode.of(5, 6, 4)), ListNode.of(7, 0, 8));
check(new Solution().addTwoNumbers(ListNode.of(9, 9, 9), ListNode.of(1)), ListNode.of(0, 0, 0, 1));
```

## 287. Find the Duplicate Number | Medium
related: 142, 41, 268
tags: fast-slow, cycle-detection

Values in `[1, n]` → view `i → nums[i]` as a linked list. The duplicate value is where two arrows enter the same node = **cycle entrance** (Floyd, like 142). No modification, O(1) space.

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def findDuplicate(self, nums: List[int]) -> int:
        slow = fast = 0
        while True:
            slow, fast = nums[slow], nums[nums[fast]]
            if slow == fast:
                break
        p = 0
        while p != slow:
            p, slow = nums[p], nums[slow]
        return p
# test
assert Solution().findDuplicate([1, 3, 4, 2, 2]) == 2
assert Solution().findDuplicate([3, 1, 3, 4, 2]) == 3
```

```java
class Solution {
    public int findDuplicate(int[] nums) {
        int slow = 0, fast = 0;
        do { slow = nums[slow]; fast = nums[nums[fast]]; } while (slow != fast);
        int p = 0;
        while (p != slow) { p = nums[p]; slow = nums[slow]; }
        return p;
    }
}
// test
check(new Solution().findDuplicate(new int[]{1, 3, 4, 2, 2}), 2);
check(new Solution().findDuplicate(new int[]{3, 1, 3, 4, 2}), 3);
```

## 23. Merge k Sorted Lists | Hard
related: 21, 378, 373
tags: blind75, heap, divide-and-conquer

Min-heap of the current head of each list; pop the smallest, push its successor. (Or pairwise merge, like merge sort.)

**Complexity:** O(N log k) for N total nodes.

```python
class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        heap = [(node.val, i, node) for i, node in enumerate(lists) if node]
        heapify(heap)
        dummy = tail = ListNode()
        while heap:
            _, i, node = heappop(heap)      # i breaks ties (nodes aren't comparable)
            tail.next = tail = node
            if node.next:
                heappush(heap, (node.next.val, i, node.next))
        return dummy.next
# test
ls = [make_list([1, 4, 5]), make_list([1, 3, 4]), make_list([2, 6])]
assert list_vals(Solution().mergeKLists(ls)) == [1, 1, 2, 3, 4, 4, 5, 6]
assert Solution().mergeKLists([]) is None
```

```java
class Solution {
    public ListNode mergeKLists(ListNode[] lists) {
        PriorityQueue<ListNode> pq = new PriorityQueue<>((a, b) -> Integer.compare(a.val, b.val));
        for (ListNode l : lists) if (l != null) pq.add(l);
        ListNode dummy = new ListNode(), tail = dummy;
        while (!pq.isEmpty()) {
            ListNode n = pq.poll();
            tail = tail.next = n;
            if (n.next != null) pq.add(n.next);
        }
        return dummy.next;
    }
}
// test
check(new Solution().mergeKLists(new ListNode[]{ListNode.of(1, 4, 5), ListNode.of(1, 3, 4), ListNode.of(2, 6)}),
      ListNode.of(1, 1, 2, 3, 4, 4, 5, 6));
```

## 25. Reverse Nodes in k-Group | Hard
related: 206, 92
tags: reversal, dummy-head

For each group: check k nodes exist, reverse them in place, then connect the previous group's tail to the new head and move on. A dummy head keeps the first group uniform.

```viz
k=2:  d -> 1 -> 2 -> 3 -> 4 -> 5
      d -> 2 -> 1 -> 4 -> 3 -> 5   (last group < k stays)
```

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def reverseKGroup(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        dummy = ListNode(0, head)
        group_prev = dummy
        while True:
            kth = group_prev
            for _ in range(k):
                kth = kth.next
                if not kth:
                    return dummy.next
            group_next = kth.next
            prev, cur = group_next, group_prev.next      # reverse [group_prev.next .. kth]
            while cur is not group_next:
                cur.next, prev, cur = prev, cur, cur.next
            first = group_prev.next
            group_prev.next = kth
            group_prev = first
# test
assert list_vals(Solution().reverseKGroup(make_list([1, 2, 3, 4, 5]), 2)) == [2, 1, 4, 3, 5]
assert list_vals(Solution().reverseKGroup(make_list([1, 2, 3, 4, 5]), 3)) == [3, 2, 1, 4, 5]
```

```java
class Solution {
    public ListNode reverseKGroup(ListNode head, int k) {
        ListNode dummy = new ListNode(0, head), groupPrev = dummy;
        while (true) {
            ListNode kth = groupPrev;
            for (int i = 0; i < k; i++) {
                kth = kth.next;
                if (kth == null) return dummy.next;
            }
            ListNode groupNext = kth.next, prev = groupNext, cur = groupPrev.next;
            while (cur != groupNext) { ListNode nxt = cur.next; cur.next = prev; prev = cur; cur = nxt; }
            ListNode first = groupPrev.next;
            groupPrev.next = kth;
            groupPrev = first;
        }
    }
}
// test
check(new Solution().reverseKGroup(ListNode.of(1, 2, 3, 4, 5), 2), ListNode.of(2, 1, 4, 3, 5));
check(new Solution().reverseKGroup(ListNode.of(1, 2, 3, 4, 5), 3), ListNode.of(3, 2, 1, 4, 5));
```

## 234. Palindrome Linked List | Easy
related: 206, 143, 125
tags: fast-slow, reversal

Find middle, reverse the second half, compare with the first half.

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def isPalindrome(self, head: Optional[ListNode]) -> bool:
        slow = fast = head
        while fast and fast.next:
            slow, fast = slow.next, fast.next.next
        prev = None
        while slow:
            slow.next, prev, slow = prev, slow, slow.next
        while prev:
            if prev.val != head.val:
                return False
            prev, head = prev.next, head.next
        return True
# test
assert Solution().isPalindrome(make_list([1, 2, 2, 1])) is True
assert Solution().isPalindrome(make_list([1, 2, 3, 2, 1])) is True
assert Solution().isPalindrome(make_list([1, 2])) is False
```

```java
class Solution {
    public boolean isPalindrome(ListNode head) {
        ListNode slow = head, fast = head;
        while (fast != null && fast.next != null) { slow = slow.next; fast = fast.next.next; }
        ListNode prev = null;
        while (slow != null) { ListNode nxt = slow.next; slow.next = prev; prev = slow; slow = nxt; }
        for (; prev != null; prev = prev.next, head = head.next)
            if (prev.val != head.val) return false;
        return true;
    }
}
// test
check(new Solution().isPalindrome(ListNode.of(1, 2, 2, 1)), true);
check(new Solution().isPalindrome(ListNode.of(1, 2)), false);
```

## 160. Intersection of Two Linked Lists | Easy
related: 141
tags: two-pointers

Walk `a` over A then B, `b` over B then A. Both travel `lenA + lenB` steps, so they meet at the intersection (or both reach `null` together).

**Complexity:** O(m + n) time, O(1) space.

```python
class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> Optional[ListNode]:
        a, b = headA, headB
        while a is not b:
            a = a.next if a else headB
            b = b.next if b else headA
        return a
# test
common = make_list([8, 4, 5])
A = make_list([4, 1]); A.next.next = common
B = make_list([5, 6, 1]); B.next.next.next = common
assert Solution().getIntersectionNode(A, B) is common
assert Solution().getIntersectionNode(make_list([1]), make_list([2])) is None
```

```java
class Solution {
    public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
        ListNode a = headA, b = headB;
        while (a != b) {
            a = a == null ? headB : a.next;
            b = b == null ? headA : b.next;
        }
        return a;
    }
}
// test
ListNode common = ListNode.of(8, 4, 5), A = ListNode.of(4, 1), B = ListNode.of(5, 6, 1);
A.next.next = common; B.next.next.next = common;
check(new Solution().getIntersectionNode(A, B) == common, true);
```

## 328. Odd Even Linked List | Medium
related: 86
tags: in-place

Build two chains (odd positions, even positions) in one pass, then attach the even chain after the odd chain.

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def oddEvenList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head:
            return head
        odd, even = head, head.next
        even_head = even
        while even and even.next:
            odd.next = even.next
            odd = odd.next
            even.next = odd.next
            even = even.next
        odd.next = even_head
        return head
# test
assert list_vals(Solution().oddEvenList(make_list([1, 2, 3, 4, 5]))) == [1, 3, 5, 2, 4]
assert list_vals(Solution().oddEvenList(make_list([2, 1, 3, 5, 6, 4, 7]))) == [2, 3, 6, 7, 1, 5, 4]
```

```java
class Solution {
    public ListNode oddEvenList(ListNode head) {
        if (head == null) return null;
        ListNode odd = head, even = head.next, evenHead = even;
        while (even != null && even.next != null) {
            odd = odd.next = even.next;
            even = even.next = odd.next;
        }
        odd.next = evenHead;
        return head;
    }
}
// test
check(new Solution().oddEvenList(ListNode.of(1, 2, 3, 4, 5)), ListNode.of(1, 3, 5, 2, 4));
```

## 148. Sort List | Medium
related: 21, 23
tags: merge-sort

Merge sort is natural for lists: split at the middle (fast/slow), sort halves recursively, merge (21).

**Complexity:** O(n log n) time, O(log n) recursion.

```python
class Solution:
    def sortList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        if not head or not head.next:
            return head
        slow, fast = head, head.next
        while fast and fast.next:
            slow, fast = slow.next, fast.next.next
        mid, slow.next = slow.next, None
        a, b = self.sortList(head), self.sortList(mid)
        dummy = tail = ListNode()
        while a and b:
            if a.val <= b.val:
                tail.next, a = a, a.next
            else:
                tail.next, b = b, b.next
            tail = tail.next
        tail.next = a or b
        return dummy.next
# test
assert list_vals(Solution().sortList(make_list([4, 2, 1, 3]))) == [1, 2, 3, 4]
assert list_vals(Solution().sortList(make_list([-1, 5, 3, 4, 0]))) == [-1, 0, 3, 4, 5]
```

```java
class Solution {
    public ListNode sortList(ListNode head) {
        if (head == null || head.next == null) return head;
        ListNode slow = head, fast = head.next;
        while (fast != null && fast.next != null) { slow = slow.next; fast = fast.next.next; }
        ListNode mid = slow.next;
        slow.next = null;
        ListNode a = sortList(head), b = sortList(mid), dummy = new ListNode(), tail = dummy;
        while (a != null && b != null) {
            if (a.val <= b.val) { tail.next = a; a = a.next; } else { tail.next = b; b = b.next; }
            tail = tail.next;
        }
        tail.next = a != null ? a : b;
        return dummy.next;
    }
}
// test
check(new Solution().sortList(ListNode.of(-1, 5, 3, 4, 0)), ListNode.of(-1, 0, 3, 4, 5));
```

## 92. Reverse Linked List II | Medium
related: 206, 25
tags: reversal, dummy-head

Walk to the node before `left`. Then repeatedly move the node after `cur` to the front of the sublist (`right − left` times) — "head insertion".

```viz
1 [2 3 4] 5   left=2 right=4
1 [3 2 4] 5   move 3 to front
1 [4 3 2] 5   move 4 to front
```

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def reverseBetween(self, head: Optional[ListNode], left: int, right: int) -> Optional[ListNode]:
        dummy = ListNode(0, head)
        prev = dummy
        for _ in range(left - 1):
            prev = prev.next
        cur = prev.next
        for _ in range(right - left):
            nxt = cur.next
            cur.next = nxt.next
            nxt.next = prev.next
            prev.next = nxt
        return dummy.next
# test
assert list_vals(Solution().reverseBetween(make_list([1, 2, 3, 4, 5]), 2, 4)) == [1, 4, 3, 2, 5]
assert list_vals(Solution().reverseBetween(make_list([5]), 1, 1)) == [5]
```

```java
class Solution {
    public ListNode reverseBetween(ListNode head, int left, int right) {
        ListNode dummy = new ListNode(0, head), prev = dummy;
        for (int i = 1; i < left; i++) prev = prev.next;
        ListNode cur = prev.next;
        for (int i = 0; i < right - left; i++) {
            ListNode nxt = cur.next;
            cur.next = nxt.next;
            nxt.next = prev.next;
            prev.next = nxt;
        }
        return dummy.next;
    }
}
// test
check(new Solution().reverseBetween(ListNode.of(1, 2, 3, 4, 5), 2, 4), ListNode.of(1, 4, 3, 2, 5));
```


## 876. Middle of the Linked List | Easy
related: 143, 234, 148
tags: fast-slow

Fast moves two steps, slow one; when fast runs out, slow is at the middle (the second middle for even lengths).

**Complexity:** O(n) time, O(1) space.

```python
class Solution:
    def middleNode(self, head: Optional[ListNode]) -> Optional[ListNode]:
        slow = fast = head
        while fast and fast.next:
            slow, fast = slow.next, fast.next.next
        return slow
# test
assert Solution().middleNode(make_list([1, 2, 3, 4, 5])).val == 3
assert Solution().middleNode(make_list([1, 2, 3, 4, 5, 6])).val == 4
```

```java
class Solution {
    public ListNode middleNode(ListNode head) {
        ListNode slow = head, fast = head;
        while (fast != null && fast.next != null) { slow = slow.next; fast = fast.next.next; }
        return slow;
    }
}
// test
check(new Solution().middleNode(ListNode.of(1, 2, 3, 4, 5, 6)), ListNode.of(4, 5, 6));
```

## 83. Remove Duplicates from Sorted List | Easy
related: 82, 26
tags: in-place

Sorted → duplicates are adjacent. Skip `cur.next` while it equals `cur`.

**Complexity:** O(n).

```python
class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        cur = head
        while cur and cur.next:
            if cur.next.val == cur.val:
                cur.next = cur.next.next
            else:
                cur = cur.next
        return head
# test
assert list_vals(Solution().deleteDuplicates(make_list([1, 1, 2, 3, 3]))) == [1, 2, 3]
assert list_vals(Solution().deleteDuplicates(make_list([]))) == []
assert list_vals(Solution().deleteDuplicates(make_list([1, 1, 1]))) == [1]
```

```java
class Solution {
    public ListNode deleteDuplicates(ListNode head) {
        for (ListNode cur = head; cur != null && cur.next != null; )
            if (cur.next.val == cur.val) cur.next = cur.next.next; else cur = cur.next;
        return head;
    }
}
// test
check(new Solution().deleteDuplicates(ListNode.of(1, 1, 2, 3, 3)), ListNode.of(1, 2, 3));
```

## 82. Remove Duplicates from Sorted List II | Medium
related: 83
tags: dummy-head

Remove **every** value that repeats. With a dummy, `prev` points to the last kept node; if the run starting at `prev.next` has length > 1, skip the whole run.

**Complexity:** O(n).

```python
class Solution:
    def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy = prev = ListNode(0, head)
        while prev.next:
            cur = prev.next
            if cur.next and cur.next.val == cur.val:
                while cur.next and cur.next.val == cur.val:
                    cur = cur.next
                prev.next = cur.next          # drop the whole run
            else:
                prev = cur
        return dummy.next
# test
assert list_vals(Solution().deleteDuplicates(make_list([1, 2, 3, 3, 4, 4, 5]))) == [1, 2, 5]
assert list_vals(Solution().deleteDuplicates(make_list([1, 1, 1, 2, 3]))) == [2, 3]
```

```java
class Solution {
    public ListNode deleteDuplicates(ListNode head) {
        ListNode dummy = new ListNode(0, head), prev = dummy;
        while (prev.next != null) {
            ListNode cur = prev.next;
            if (cur.next != null && cur.next.val == cur.val) {
                while (cur.next != null && cur.next.val == cur.val) cur = cur.next;
                prev.next = cur.next;
            } else prev = cur;
        }
        return dummy.next;
    }
}
// test
check(new Solution().deleteDuplicates(ListNode.of(1, 2, 3, 3, 4, 4, 5)), ListNode.of(1, 2, 5));
```

## 24. Swap Nodes in Pairs | Medium
related: 25
tags: dummy-head

With `prev` before each pair `a → b`: rewire to `prev → b → a → (rest)` and advance `prev = a`.

**Complexity:** O(n).

```python
class Solution:
    def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy = prev = ListNode(0, head)
        while prev.next and prev.next.next:
            a, b = prev.next, prev.next.next
            prev.next, a.next, b.next = b, b.next, a
            prev = a
        return dummy.next
# test
assert list_vals(Solution().swapPairs(make_list([1, 2, 3, 4]))) == [2, 1, 4, 3]
assert list_vals(Solution().swapPairs(make_list([1, 2, 3]))) == [2, 1, 3]
```

```java
class Solution {
    public ListNode swapPairs(ListNode head) {
        ListNode dummy = new ListNode(0, head), prev = dummy;
        while (prev.next != null && prev.next.next != null) {
            ListNode a = prev.next, b = a.next;
            a.next = b.next;
            b.next = a;
            prev.next = b;
            prev = a;
        }
        return dummy.next;
    }
}
// test
check(new Solution().swapPairs(ListNode.of(1, 2, 3, 4)), ListNode.of(2, 1, 4, 3));
```

## 61. Rotate List | Medium
related: 189
tags: ring

Find the length and tail; connect tail → head to form a ring; the new tail is `n − k % n − 1` steps from the head; break the ring after it.

**Complexity:** O(n).

```python
class Solution:
    def rotateRight(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        if not head:
            return None
        n, tail = 1, head
        while tail.next:
            tail, n = tail.next, n + 1
        tail.next = head                      # ring
        for _ in range(n - k % n - 1):
            head = head.next
        new_head, head.next = head.next, None
        return new_head
# test
assert list_vals(Solution().rotateRight(make_list([1, 2, 3, 4, 5]), 2)) == [4, 5, 1, 2, 3]
assert list_vals(Solution().rotateRight(make_list([0, 1, 2]), 4)) == [2, 0, 1]
```

```java
class Solution {
    public ListNode rotateRight(ListNode head, int k) {
        if (head == null) return null;
        int n = 1;
        ListNode tail = head;
        while (tail.next != null) { tail = tail.next; n++; }
        tail.next = head;
        ListNode newTail = head;
        for (int i = 0; i < n - k % n - 1; i++) newTail = newTail.next;
        ListNode newHead = newTail.next;
        newTail.next = null;
        return newHead;
    }
}
// test
check(new Solution().rotateRight(ListNode.of(1, 2, 3, 4, 5), 2), ListNode.of(4, 5, 1, 2, 3));
```

## 86. Partition List | Medium
related: 75, 328
tags: two-lists

Build two chains (`< x` and `≥ x`) with dummy heads, preserving order, then join them. Terminate the second chain!

**Complexity:** O(n).

```python
class Solution:
    def partition(self, head: Optional[ListNode], x: int) -> Optional[ListNode]:
        small = s = ListNode()
        big = b = ListNode()
        while head:
            if head.val < x:
                s.next = s = head
            else:
                b.next = b = head
            head = head.next
        b.next = None
        s.next = big.next
        return small.next
# test
assert list_vals(Solution().partition(make_list([1, 4, 3, 2, 5, 2]), 3)) == [1, 2, 2, 4, 3, 5]
assert list_vals(Solution().partition(make_list([2, 1]), 2)) == [1, 2]
assert list_vals(Solution().partition(make_list([]), 0)) == []
```

```java
class Solution {
    public ListNode partition(ListNode head, int x) {
        ListNode small = new ListNode(), big = new ListNode(), s = small, b = big;
        for (; head != null; head = head.next) {
            if (head.val < x) s = s.next = head; else b = b.next = head;
        }
        b.next = null;
        s.next = big.next;
        return small.next;
    }
}
// test
check(new Solution().partition(ListNode.of(1, 4, 3, 2, 5, 2), 3), ListNode.of(1, 2, 2, 4, 3, 5));
```

## 445. Add Two Numbers II | Medium
related: 2
tags: stack

Most-significant digit first → push both lists' digits onto stacks, pop and add from the least significant, **prepending** each result node.

**Complexity:** O(m + n).

```python
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        a, b = [], []
        while l1:
            a.append(l1.val); l1 = l1.next
        while l2:
            b.append(l2.val); l2 = l2.next
        head, carry = None, 0
        while a or b or carry:
            s = carry + (a.pop() if a else 0) + (b.pop() if b else 0)
            carry, d = divmod(s, 10)
            head = ListNode(d, head)
        return head
# test
assert list_vals(Solution().addTwoNumbers(make_list([7, 2, 4, 3]), make_list([5, 6, 4]))) == [7, 8, 0, 7]
assert list_vals(Solution().addTwoNumbers(make_list([5]), make_list([5]))) == [1, 0]
```

```java
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        Deque<Integer> a = new ArrayDeque<>(), b = new ArrayDeque<>();
        for (; l1 != null; l1 = l1.next) a.push(l1.val);
        for (; l2 != null; l2 = l2.next) b.push(l2.val);
        ListNode head = null;
        int carry = 0;
        while (!a.isEmpty() || !b.isEmpty() || carry != 0) {
            int s = carry + (a.isEmpty() ? 0 : a.pop()) + (b.isEmpty() ? 0 : b.pop());
            carry = s / 10;
            head = new ListNode(s % 10, head);
        }
        return head;
    }
}
// test
check(new Solution().addTwoNumbers(ListNode.of(7, 2, 4, 3), ListNode.of(5, 6, 4)), ListNode.of(7, 8, 0, 7));
```
