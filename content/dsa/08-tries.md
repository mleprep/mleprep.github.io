# Tries

**When to use:** many prefix queries, autocomplete, word dictionaries with wildcards, searching many words in a grid at once, and **bitwise** tries for XOR maximization.

Each node = map of `char → child` + `end` flag. Insert/search cost O(word length), independent of the number of words.

```viz
insert "app", "apple", "bat"
(root)
 ├─ a ─ p ─ p* ─ l ─ e*
 └─ b ─ a ─ t*                 * = end of word
```

## 208. Implement Trie (Prefix Tree) | Medium
related: 211, 212, 14
tags: blind75, design

Walk/create child nodes char by char. `search` needs `end` at the last node; `startsWith` only needs the path to exist.

**Complexity:** O(L) per operation.

```python
class Trie:
    def __init__(self):
        self.root = {}

    def insert(self, word: str) -> None:
        node = self.root
        for ch in word:
            node = node.setdefault(ch, {})
        node["$"] = True                 # end marker

    def _walk(self, s):
        node = self.root
        for ch in s:
            if ch not in node:
                return None
            node = node[ch]
        return node

    def search(self, word: str) -> bool:
        node = self._walk(word)
        return node is not None and "$" in node

    def startsWith(self, prefix: str) -> bool:
        return self._walk(prefix) is not None
# test
t = Trie()
t.insert("apple")
assert t.search("apple") and not t.search("app") and t.startsWith("app")
t.insert("app")
assert t.search("app")
```

```java
class Trie {
    private static class Node { Node[] next = new Node[26]; boolean end; }
    private final Node root = new Node();

    public void insert(String word) {
        Node n = root;
        for (char c : word.toCharArray()) {
            if (n.next[c - 'a'] == null) n.next[c - 'a'] = new Node();
            n = n.next[c - 'a'];
        }
        n.end = true;
    }

    private Node walk(String s) {
        Node n = root;
        for (char c : s.toCharArray()) if ((n = n.next[c - 'a']) == null) return null;
        return n;
    }

    public boolean search(String word) { Node n = walk(word); return n != null && n.end; }
    public boolean startsWith(String prefix) { return walk(prefix) != null; }
}
// test
Trie t = new Trie();
t.insert("apple");
check(t.search("apple") && !t.search("app") && t.startsWith("app"), true);
t.insert("app");
check(t.search("app"), true);
```

## 211. Design Add and Search Words Data Structure | Medium
related: 208, 212
tags: blind75, design, dfs

Trie + DFS: on `.`, try every child.

**Complexity:** add O(L); search O(26^dots · L) worst case.

```python
class WordDictionary:
    def __init__(self):
        self.root = {}

    def addWord(self, word: str) -> None:
        node = self.root
        for ch in word:
            node = node.setdefault(ch, {})
        node["$"] = True

    def search(self, word: str) -> bool:
        def dfs(node, i):
            if i == len(word):
                return "$" in node
            ch = word[i]
            if ch == ".":
                return any(dfs(child, i + 1) for k, child in node.items() if k != "$")
            return ch in node and dfs(node[ch], i + 1)

        return dfs(self.root, 0)
# test
wd = WordDictionary()
for w in ("bad", "dad", "mad"):
    wd.addWord(w)
assert not wd.search("pad") and wd.search("bad") and wd.search(".ad") and wd.search("b..")
```

```java
class WordDictionary {
    private static class Node { Node[] next = new Node[26]; boolean end; }
    private final Node root = new Node();

    public void addWord(String word) {
        Node n = root;
        for (char c : word.toCharArray()) {
            if (n.next[c - 'a'] == null) n.next[c - 'a'] = new Node();
            n = n.next[c - 'a'];
        }
        n.end = true;
    }

    public boolean search(String word) { return dfs(root, word, 0); }

    private boolean dfs(Node n, String w, int i) {
        if (i == w.length()) return n.end;
        char c = w.charAt(i);
        if (c == '.') {
            for (Node child : n.next) if (child != null && dfs(child, w, i + 1)) return true;
            return false;
        }
        Node child = n.next[c - 'a'];
        return child != null && dfs(child, w, i + 1);
    }
}
// test
WordDictionary wd = new WordDictionary();
wd.addWord("bad"); wd.addWord("dad"); wd.addWord("mad");
check(!wd.search("pad") && wd.search("bad") && wd.search(".ad") && wd.search("b.."), true);
```

## 212. Word Search II | Hard
related: 79, 208
tags: blind75, backtracking

Put all words in a trie, then DFS from every cell **following the trie** — one traversal finds all words. Store the full word at its end node; clear it after finding (dedupe) and prune empty branches.

```viz
board  o a a n        trie of {oath, pea, eat, rain}
       e t a e        DFS from 'o': o -> a -> t -> h  ✓ "oath"
       i h k r        DFS from 'e': e -> a -> t       ✓ "eat"
       i f l v
```

**Complexity:** O(cells · 4 · 3^(L−1)) worst case, heavily pruned in practice.

```python
class Solution:
    def findWords(self, board: List[List[str]], words: List[str]) -> List[str]:
        root = {}
        for w in words:
            node = root
            for ch in w:
                node = node.setdefault(ch, {})
            node["$"] = w
        R, C, res = len(board), len(board[0]), []

        def dfs(r, c, parent):
            ch = board[r][c]
            node = parent[ch]
            if "$" in node:
                res.append(node.pop("$"))
            board[r][c] = "#"
            for nr, nc in ((r + 1, c), (r - 1, c), (r, c + 1), (r, c - 1)):
                if 0 <= nr < R and 0 <= nc < C and board[nr][nc] in node:
                    dfs(nr, nc, node)
            board[r][c] = ch
            if not node:
                parent.pop(ch)               # prune exhausted branch

        for r in range(R):
            for c in range(C):
                if board[r][c] in root:
                    dfs(r, c, root)
        return res
# test
b = [list("oaan"), list("etae"), list("ihkr"), list("iflv")]
assert sorted(Solution().findWords(b, ["oath", "pea", "eat", "rain"])) == ["eat", "oath"]
```

```java
class Solution {
    private static class Node { Node[] next = new Node[26]; String word; }

    public List<String> findWords(char[][] board, String[] words) {
        Node root = new Node();
        for (String w : words) {
            Node n = root;
            for (char c : w.toCharArray()) {
                if (n.next[c - 'a'] == null) n.next[c - 'a'] = new Node();
                n = n.next[c - 'a'];
            }
            n.word = w;
        }
        List<String> res = new ArrayList<>();
        for (int r = 0; r < board.length; r++)
            for (int c = 0; c < board[0].length; c++) dfs(board, r, c, root, res);
        return res;
    }

    private void dfs(char[][] b, int r, int c, Node parent, List<String> res) {
        if (r < 0 || c < 0 || r >= b.length || c >= b[0].length || b[r][c] == '#') return;
        char ch = b[r][c];
        Node n = parent.next[ch - 'a'];
        if (n == null) return;
        if (n.word != null) { res.add(n.word); n.word = null; }
        b[r][c] = '#';
        dfs(b, r + 1, c, n, res); dfs(b, r - 1, c, n, res);
        dfs(b, r, c + 1, n, res); dfs(b, r, c - 1, n, res);
        b[r][c] = ch;
    }
}
// test
char[][] b = {"oaan".toCharArray(), "etae".toCharArray(), "ihkr".toCharArray(), "iflv".toCharArray()};
List<String> r = new Solution().findWords(b, new String[]{"oath", "pea", "eat", "rain"});
Collections.sort(r);
check(r, "[eat, oath]");
```

## 421. Maximum XOR of Two Numbers in an Array | Medium
related: 208, 136
tags: bit-trie, greedy

Insert numbers as 31-bit binary strings into a trie (high bit first). For each number, walk greedily toward the **opposite** bit at every level to maximize XOR.

```viz
x = 5 = 00101   want bits 11010 -> follow opposite branches when they exist
best partner for 5 in [3,10,5,25,2,8] is 25 (11001) -> 5^25 = 28
```

**Complexity:** O(32 · n).

```python
class Solution:
    def findMaximumXOR(self, nums: List[int]) -> int:
        root, best = {}, 0
        for x in nums:
            node = root
            for b in range(30, -1, -1):
                node = node.setdefault((x >> b) & 1, {})
        for x in nums:
            node, cur = root, 0
            for b in range(30, -1, -1):
                bit = (x >> b) & 1
                if 1 - bit in node:
                    cur |= 1 << b
                    node = node[1 - bit]
                else:
                    node = node[bit]
            best = max(best, cur)
        return best
# test
assert Solution().findMaximumXOR([3, 10, 5, 25, 2, 8]) == 28
assert Solution().findMaximumXOR([14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70]) == 127
```

```java
class Solution {
    public int findMaximumXOR(int[] nums) {
        int[][] trie = new int[nums.length * 31 + 1][2];   // array-backed trie, 0 = no child
        int size = 1, best = 0;
        for (int x : nums) {
            int n = 0;
            for (int b = 30; b >= 0; b--) {
                int bit = (x >> b) & 1;
                if (trie[n][bit] == 0) trie[n][bit] = size++;
                n = trie[n][bit];
            }
        }
        for (int x : nums) {
            int n = 0, cur = 0;
            for (int b = 30; b >= 0; b--) {
                int bit = (x >> b) & 1;
                if (trie[n][1 - bit] != 0) { cur |= 1 << b; n = trie[n][1 - bit]; }
                else n = trie[n][bit];
            }
            best = Math.max(best, cur);
        }
        return best;
    }
}
// test
check(new Solution().findMaximumXOR(new int[]{3, 10, 5, 25, 2, 8}), 28);
```


## 648. Replace Words | Medium
related: 208
tags: trie

Insert the roots into a trie. For each word, walk the trie and stop at the **first** end-of-root marker; replace the word with that prefix.

**Complexity:** O(total characters).

```python
class Solution:
    def replaceWords(self, dictionary: List[str], sentence: str) -> str:
        root = {}
        for w in dictionary:
            node = root
            for ch in w:
                node = node.setdefault(ch, {})
            node["$"] = True

        def shortest(word):
            node = root
            for i, ch in enumerate(word):
                if ch not in node:
                    return word
                node = node[ch]
                if "$" in node:
                    return word[:i + 1]
            return word

        return " ".join(shortest(w) for w in sentence.split())
# test
assert Solution().replaceWords(["cat", "bat", "rat"], "the cattle was rattled by the battery") == "the cat was rat by the bat"
assert Solution().replaceWords(["a", "b", "c"], "aadsfasf absbs bbab cadsfafs") == "a a b c"
```

```java
class Solution {
    private static class Node { Node[] next = new Node[26]; boolean end; }

    public String replaceWords(List<String> dictionary, String sentence) {
        Node root = new Node();
        for (String w : dictionary) {
            Node n = root;
            for (char c : w.toCharArray()) {
                if (n.next[c - 'a'] == null) n.next[c - 'a'] = new Node();
                n = n.next[c - 'a'];
            }
            n.end = true;
        }
        StringBuilder sb = new StringBuilder();
        for (String w : sentence.split(" ")) {
            if (sb.length() > 0) sb.append(' ');
            Node n = root;
            int i = 0;
            for (; i < w.length() && n != null && !n.end; i++) n = n.next[w.charAt(i) - 'a'];
            sb.append(n != null && n.end ? w.substring(0, i) : w);
        }
        return sb.toString();
    }
}
// test
check(new Solution().replaceWords(List.of("cat", "bat", "rat"), "the cattle was rattled by the battery"), "the cat was rat by the bat");
```

## 1268. Search Suggestions System | Medium
related: 208, 642
tags: sort, binary-search, trie

Sort the products once. For each typed prefix, binary-search its lower bound; the next 3 products that start with the prefix are the suggestions. (A trie storing the top-3 at each node is the alternative.)

**Complexity:** O(n log n + L log n).

```python
class Solution:
    def suggestedProducts(self, products: List[str], searchWord: str) -> List[List[str]]:
        products.sort()
        res, prefix = [], ""
        for ch in searchWord:
            prefix += ch
            i = bisect.bisect_left(products, prefix)
            res.append([p for p in products[i:i + 3] if p.startswith(prefix)])
        return res
# test
P = ["mobile", "mouse", "moneypot", "monitor", "mousepad"]
assert Solution().suggestedProducts(P, "mouse") == [
    ["mobile", "moneypot", "monitor"], ["mobile", "moneypot", "monitor"],
    ["mouse", "mousepad"], ["mouse", "mousepad"], ["mouse", "mousepad"]]
```

```java
class Solution {
    public List<List<String>> suggestedProducts(String[] products, String searchWord) {
        Arrays.sort(products);
        List<List<String>> res = new ArrayList<>();
        String prefix = "";
        for (char c : searchWord.toCharArray()) {
            prefix += c;
            int lo = 0, hi = products.length;
            while (lo < hi) {
                int mid = (lo + hi) / 2;
                if (products[mid].compareTo(prefix) >= 0) hi = mid; else lo = mid + 1;
            }
            List<String> cur = new ArrayList<>();
            for (int i = lo; i < Math.min(lo + 3, products.length) && products[i].startsWith(prefix); i++) cur.add(products[i]);
            res.add(cur);
        }
        return res;
    }
}
// test
check(new Solution().suggestedProducts(new String[]{"havana"}, "tatiana"), "[[], [], [], [], [], [], []]");
check(new Solution().suggestedProducts(new String[]{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mou").get(2), "[mouse, mousepad]");
```

## 745. Prefix and Suffix Search | Hard
related: 208, 211
tags: trie, hashing

Words are short (≤ 7 chars), so precompute every `(prefix, suffix)` combination per word into a map `prefix + "#" + suffix → index` (later words overwrite → largest index wins). Queries become O(1). Trie alternative: insert `suffix + "#" + word` for every suffix and store the max index at each node.

**Complexity:** O(N · L²) build (L³ including string copies), O(L) query.

```python
class WordFilter:
    def __init__(self, words: List[str]):
        self.idx = {}
        for i, w in enumerate(words):
            for p in range(len(w) + 1):
                for s in range(len(w) + 1):
                    self.idx[w[:p] + "#" + w[s:]] = i

    def f(self, pref: str, suff: str) -> int:
        return self.idx.get(pref + "#" + suff, -1)
# test
wf = WordFilter(["apple", "ape", "maple"])
assert wf.f("a", "e") == 1 and wf.f("ap", "le") == 0 and wf.f("", "le") == 2 and wf.f("b", "") == -1
```

```java
class WordFilter {
    private final Map<String, Integer> idx = new HashMap<>();

    public WordFilter(String[] words) {
        for (int i = 0; i < words.length; i++) {
            String w = words[i];
            for (int p = 0; p <= w.length(); p++)
                for (int s = 0; s <= w.length(); s++) idx.put(w.substring(0, p) + "#" + w.substring(s), i);
        }
    }

    public int f(String pref, String suff) { return idx.getOrDefault(pref + "#" + suff, -1); }
}
// test
WordFilter wf = new WordFilter(new String[]{"apple", "ape", "maple"});
check(wf.f("a", "e"), 1);
check(wf.f("", "le"), 2);
check(wf.f("b", ""), -1);
```
