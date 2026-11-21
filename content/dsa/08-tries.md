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

```problem
Implement a `Trie` (prefix tree) with `insert(word)`, `search(word)` (was this exact word inserted?) and `startsWith(prefix)` (was any word with this prefix inserted?).

**Example:** `insert("apple")` → `search("apple")` = true, `search("app")` = false, `startsWith("app")` = true

**Constraints:** lowercase letters; up to 3·10⁴ calls.
```

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

```go
type Trie struct {
	children [26]*Trie
	end      bool // a word ends here
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if node.children[c] == nil {
			node.children[c] = &Trie{}
		}
		node = node.children[c]
	}
	node.end = true
}

func (this *Trie) walk(s string) *Trie {
	node := this
	for i := 0; i < len(s) && node != nil; i++ {
		node = node.children[s[i]-'a']
	}
	return node
}

func (this *Trie) Search(word string) bool {
	node := this.walk(word)
	return node != nil && node.end
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.walk(prefix) != nil
}
// test
t := Constructor()
t.Insert("apple")
assertEq([]bool{t.Search("apple"), t.Search("app"), t.StartsWith("app")}, []bool{true, false, true})
t.Insert("app")
assertEq(t.Search("app"), true)
```

## 211. Design Add and Search Words Data Structure | Medium
related: 208, 212
tags: blind75, design, dfs

```problem
Design `WordDictionary` with `addWord(word)` and `search(word)`, where the search pattern may contain `.`, which matches any single letter.

**Example:** add `"bad"`, `"dad"`, `"mad"` → `search("pad")` = false, `search(".ad")` = true, `search("b..")` = true

**Constraints:** at most 2 dots per search; word length ≤ 25.
```

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
assert wd.search("...") and not wd.search(".") and not wd.search("ba") and not wd.search("bad.")
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

```go
type WordDictionary struct {
	children [26]*WordDictionary
	end      bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	node := this
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if node.children[c] == nil {
			node.children[c] = &WordDictionary{}
		}
		node = node.children[c]
	}
	node.end = true
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(node *WordDictionary, i int) bool
	dfs = func(node *WordDictionary, i int) bool {
		if i == len(word) {
			return node.end
		}
		if word[i] == '.' { // try every child
			for _, child := range node.children {
				if child != nil && dfs(child, i+1) {
					return true
				}
			}
			return false
		}
		child := node.children[word[i]-'a']
		return child != nil && dfs(child, i+1)
	}
	return dfs(this, 0)
}
// test
wd := Constructor()
for _, w := range []string{"bad", "dad", "mad"} {
	wd.AddWord(w)
}
assertEq([]bool{wd.Search("pad"), wd.Search("bad"), wd.Search(".ad"), wd.Search("b..")}, []bool{false, true, true, true})
assertEq([]bool{wd.Search("..."), wd.Search("."), wd.Search("ba"), wd.Search("bad.")}, []bool{true, false, false, false})
```

## 212. Word Search II | Hard
related: 79, 208
tags: blind75, backtracking

```problem
Given an m×n `board` of letters and a list of `words`, return every word that can be traced on the board through horizontally or vertically adjacent cells, using each cell at most once per word.

**Example:** `board = ["oaan","etae","ihkr","iflv"], words = ["oath","pea","eat","rain"]` → `["eat","oath"]`

**Constraints:** board up to 12 × 12; up to 3·10⁴ words of length ≤ 10.
```

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
assert Solution().findWords([["a", "b"], ["c", "d"]], ["abcb"]) == []
assert Solution().findWords([["a"]], ["a"]) == ["a"]
assert sorted(Solution().findWords([["a", "b"], ["c", "d"]], ["ab", "cb", "ad", "bd", "ac", "ca", "da", "bc", "db", "adcb", "dabc", "abb", "acb"])) == ['ab', 'ac', 'bd', 'ca', 'db']
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

```go
type trieNode struct {
	next [26]*trieNode
	word string // set at the node where a word ends
	kids int    // live children, for pruning
}

func findWords(board [][]byte, words []string) []string {
	root := &trieNode{}
	for _, w := range words {
		node := root
		for i := 0; i < len(w); i++ {
			c := w[i] - 'a'
			if node.next[c] == nil {
				node.next[c] = &trieNode{}
				node.kids++
			}
			node = node.next[c]
		}
		node.word = w
	}
	R, C, res := len(board), len(board[0]), []string{}
	var dfs func(r, c int, parent *trieNode)
	dfs = func(r, c int, parent *trieNode) {
		ch := board[r][c]
		node := parent.next[ch-'a']
		if node.word != "" {
			res = append(res, node.word)
			node.word = "" // report each word once
		}
		board[r][c] = '#'
		for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < R && nc >= 0 && nc < C && board[nr][nc] != '#' && node.next[board[nr][nc]-'a'] != nil {
				dfs(nr, nc, node)
			}
		}
		board[r][c] = ch
		if node.kids == 0 && node.word == "" { // prune an exhausted branch
			parent.next[ch-'a'] = nil
			parent.kids--
		}
	}
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			if root.next[board[r][c]-'a'] != nil {
				dfs(r, c, root)
			}
		}
	}
	return res
}
// test
assertUnordered(findWords(byteGrid("oaan", "etae", "ihkr", "iflv"), []string{"oath", "pea", "eat", "rain"}), []string{"eat", "oath"})
assertEq(findWords(byteGrid("ab", "cd"), []string{"abcb"}), []string{})
assertEq(findWords(byteGrid("a"), []string{"a"}), []string{"a"})
assertUnordered(findWords(byteGrid("ab", "cd"), []string{"ab", "cb", "ad", "bd", "ac", "ca", "da", "bc", "db", "adcb", "dabc", "abb", "acb"}), []string{"ab", "ac", "bd", "ca", "db"})
```

## 421. Maximum XOR of Two Numbers in an Array | Medium
related: 208, 136
tags: bit-trie, greedy

```problem
Return the maximum of `nums[i] XOR nums[j]` over all pairs i ≤ j.

**Example:** `[3,10,5,25,2,8]` → `28` (5 XOR 25)

**Constraints:** n ≤ 2·10⁵; values < 2³¹.
```

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

```go
func findMaximumXOR(nums []int) int {
	type node struct{ next [2]*node }
	root := &node{}
	for _, x := range nums { // binary trie of the numbers, high bit first
		n := root
		for b := 30; b >= 0; b-- {
			bit := x >> b & 1
			if n.next[bit] == nil {
				n.next[bit] = &node{}
			}
			n = n.next[bit]
		}
	}
	best := 0
	for _, x := range nums {
		n, cur := root, 0
		for b := 30; b >= 0; b-- {
			bit := x >> b & 1
			if n.next[1-bit] != nil { // take the opposite bit whenever possible
				cur |= 1 << b
				n = n.next[1-bit]
			} else {
				n = n.next[bit]
			}
		}
		best = max(best, cur)
	}
	return best
}
// test
assertEq(findMaximumXOR([]int{3, 10, 5, 25, 2, 8}), 28)
assertEq(findMaximumXOR([]int{14, 70, 53, 83, 49, 91, 36, 80, 92, 51, 66, 70}), 127)
```


## 648. Replace Words | Medium
related: 208
tags: trie

```problem
Replace every word of `sentence` that starts with one of the `dictionary` roots by the shortest such root; other words stay unchanged.

**Example:** `dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"` → `"the cat was rat by the bat"`
```

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

```go
func replaceWords(dictionary []string, sentence string) string {
	type node struct {
		next [26]*node
		end  bool
	}
	root := &node{}
	for _, w := range dictionary {
		n := root
		for i := 0; i < len(w); i++ {
			c := w[i] - 'a'
			if n.next[c] == nil {
				n.next[c] = &node{}
			}
			n = n.next[c]
		}
		n.end = true
	}
	shortest := func(word string) string { // the shortest root that prefixes word
		n := root
		for i := 0; i < len(word); i++ {
			if n = n.next[word[i]-'a']; n == nil {
				return word
			}
			if n.end {
				return word[:i+1]
			}
		}
		return word
	}
	words := strings.Fields(sentence)
	for i, w := range words {
		words[i] = shortest(w)
	}
	return strings.Join(words, " ")
}
// test
assertEq(replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"), "the cat was rat by the bat")
assertEq(replaceWords([]string{"a", "b", "c"}, "aadsfasf absbs bbab cadsfafs"), "a a b c")
```

## 1268. Search Suggestions System | Medium
related: 208, 642
tags: sort, binary-search, trie

```problem
As each character of `searchWord` is typed, suggest up to three `products` that start with the prefix typed so far, choosing the lexicographically smallest ones. Return the suggestion list after every character.

**Example:** `products = ["mobile","mouse","moneypot","monitor","mousepad"], searchWord = "mouse"` → `[["mobile","moneypot","monitor"],["mobile","moneypot","monitor"],["mouse","mousepad"],["mouse","mousepad"],["mouse","mousepad"]]`
```

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
assert Solution().suggestedProducts(["havana"], "tatiana") == [[], [], [], [], [], [], []]
assert Solution().suggestedProducts(["bags", "baggage", "banner", "box", "cloths"], "bags") == [["baggage", "bags", "banner"], ["baggage", "bags", "banner"], ["baggage", "bags"], ["bags"]]
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

```go
func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	res := [][]string{}
	for k := 1; k <= len(searchWord); k++ {
		prefix := searchWord[:k]
		i := sort.SearchStrings(products, prefix) // first product >= prefix
		group := []string{}
		for j := i; j < len(products) && j < i+3 && strings.HasPrefix(products[j], prefix); j++ {
			group = append(group, products[j])
		}
		res = append(res, group)
	}
	return res
}
// test
P := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
assertEq(suggestedProducts(P, "mouse"), [][]string{
	{"mobile", "moneypot", "monitor"}, {"mobile", "moneypot", "monitor"},
	{"mouse", "mousepad"}, {"mouse", "mousepad"}, {"mouse", "mousepad"}})
assertEq(suggestedProducts([]string{"havana"}, "tatiana"), [][]string{{}, {}, {}, {}, {}, {}, {}})
assertEq(suggestedProducts([]string{"bags", "baggage", "banner", "box", "cloths"}, "bags"), [][]string{{"baggage", "bags", "banner"}, {"baggage", "bags", "banner"}, {"baggage", "bags"}, {"bags"}})
```

## 745. Prefix and Suffix Search | Hard
related: 208, 211
tags: trie, hashing

```problem
Design `WordFilter(words)` whose `f(pref, suff)` returns the largest index of a word that starts with `pref` and ends with `suff`, or −1 if there is none.

**Example:** `WordFilter(["apple"])` → `f("a","e")` = 0

**Constraints:** up to 10⁴ words of length ≤ 7; up to 10⁴ calls to `f`.
```

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
assert WordFilter(["apple"]).f("a", "e") == 0
assert WordFilter(["abc", "abc"]).f("a", "c") == 1
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

```go
type WordFilter struct {
	idx map[string]int // "prefix#suffix" -> largest index
}

func Constructor(words []string) WordFilter {
	idx := map[string]int{}
	for i, w := range words {
		for p := 0; p <= len(w); p++ {
			for s := 0; s <= len(w); s++ {
				idx[w[:p]+"#"+w[s:]] = i // later words overwrite earlier ones
			}
		}
	}
	return WordFilter{idx}
}

func (this *WordFilter) F(pref string, suff string) int {
	if i, ok := this.idx[pref+"#"+suff]; ok {
		return i
	}
	return -1
}
// test
wf := Constructor([]string{"apple", "ape", "maple"})
assertEq([]int{wf.F("a", "e"), wf.F("ap", "le"), wf.F("", "le"), wf.F("b", "")}, []int{1, 0, 2, -1})
wf2 := Constructor([]string{"apple"})
assertEq(wf2.F("a", "e"), 0)
wf3 := Constructor([]string{"abc", "abc"})
assertEq(wf3.F("a", "c"), 1)
```

## 336. Palindrome Pairs | Hard
related: 214, 5, 131
tags: hash-map, palindrome

```problem
`words` holds distinct strings. Return every pair of indices `[i, j]` (i ≠ j) for which `words[i] + words[j]` is a palindrome.

**Example:** `["bat","tab","cat"]` → `[[0,1],[1,0]]`

**Constraints:** up to 5000 words of length ≤ 300; aim for better than checking every pair.
```

Map each **reversed** word to its index. For every word `w` and every split `w = pre + suf`:
- if `pre` is some other word reversed and `suf` is a palindrome → `w + that word` is a palindrome;
- if `suf` is some other word reversed and `pre` is a palindrome → `that word + w` is a palindrome (skip the empty `pre` here so each pair is found once).

A trie of reversed words gives the same complexity; the hash map is simpler to write.

```viz
w = "lls" split "" + "lls"? ...  "ll" + "s": pre "ll" is a palindrome, rev(suf) = "s" is a word
    -> "s" + "lls" = "slls"  ✓
```

**Complexity:** O(n · k²) for n words of length ≤ k.

```python
class Solution:
    def palindromePairs(self, words: List[str]) -> List[List[int]]:
        rev = {w[::-1]: i for i, w in enumerate(words)}
        res = []
        for i, w in enumerate(words):
            for j in range(len(w) + 1):
                pre, suf = w[:j], w[j:]
                if pre in rev and rev[pre] != i and suf == suf[::-1]:
                    res.append([i, rev[pre]])
                if j and suf in rev and rev[suf] != i and pre == pre[::-1]:
                    res.append([rev[suf], i])
        return res
# test
assert sorted(Solution().palindromePairs(["abcd", "dcba", "lls", "s", "sssll"])) == [[0, 1], [1, 0], [2, 4], [3, 2]]
assert sorted(Solution().palindromePairs(["bat", "tab", "cat"])) == [[0, 1], [1, 0]]
assert sorted(Solution().palindromePairs(["a", ""])) == [[0, 1], [1, 0]]
assert sorted(Solution().palindromePairs(["a", "b", "c"])) == []
assert sorted(Solution().palindromePairs(["ab", "ba", "aba"])) == [[0, 1], [0, 2], [1, 0], [2, 1]]  # "ababa" twice
```

```java
class Solution {
    public List<List<Integer>> palindromePairs(String[] words) {
        Map<String, Integer> rev = new HashMap<>();
        for (int i = 0; i < words.length; i++) rev.put(new StringBuilder(words[i]).reverse().toString(), i);
        List<List<Integer>> res = new ArrayList<>();
        for (int i = 0; i < words.length; i++) {
            String w = words[i];
            for (int j = 0; j <= w.length(); j++) {
                String pre = w.substring(0, j), suf = w.substring(j);
                Integer k = rev.get(pre);
                if (k != null && k != i && isPal(suf)) res.add(List.of(i, k));
                k = rev.get(suf);
                if (j > 0 && k != null && k != i && isPal(pre)) res.add(List.of(k, i));
            }
        }
        return res;
    }

    private boolean isPal(String s) {
        for (int i = 0, j = s.length() - 1; i < j; i++, j--) if (s.charAt(i) != s.charAt(j)) return false;
        return true;
    }
}
// test
List<List<Integer>> pp = new ArrayList<>(new Solution().palindromePairs(new String[]{"abcd", "dcba", "lls", "s", "sssll"}));
pp.sort(Comparator.comparing((List<Integer> x) -> x.get(0)).thenComparing(x -> x.get(1)));
check(pp, List.of(List.of(0, 1), List.of(1, 0), List.of(2, 4), List.of(3, 2)));
```

```go
func palindromePairs(words []string) [][]int {
	isPal := func(s string) bool {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}
	reverse := func(s string) string {
		b := []byte(s)
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
		return string(b)
	}
	rev := map[string]int{}
	for i, w := range words {
		rev[reverse(w)] = i
	}
	res := [][]int{}
	for i, w := range words {
		for j := 0; j <= len(w); j++ {
			pre, suf := w[:j], w[j:]
			// w + words[k] is a palindrome if words[k] reversed == pre and suf is a palindrome
			if k, ok := rev[pre]; ok && k != i && isPal(suf) {
				res = append(res, []int{i, k})
			}
			// words[k] + w is a palindrome if words[k] reversed == suf and pre is a palindrome
			if k, ok := rev[suf]; ok && j > 0 && k != i && isPal(pre) {
				res = append(res, []int{k, i})
			}
		}
	}
	return res
}
// test
assertUnordered(palindromePairs([]string{"abcd", "dcba", "lls", "s", "sssll"}), [][]int{{0, 1}, {1, 0}, {2, 4}, {3, 2}})
assertUnordered(palindromePairs([]string{"bat", "tab", "cat"}), [][]int{{0, 1}, {1, 0}})
assertUnordered(palindromePairs([]string{"a", ""}), [][]int{{0, 1}, {1, 0}})
assertEq(palindromePairs([]string{"a", "b", "c"}), [][]int{})
assertUnordered(palindromePairs([]string{"ab", "ba", "aba"}), [][]int{{0, 1}, {0, 2}, {1, 0}, {2, 1}}) // "ababa" twice
```
