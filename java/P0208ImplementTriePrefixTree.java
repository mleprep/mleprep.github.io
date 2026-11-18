// 208. Implement Trie (Prefix Tree) (Medium)
// https://leetcode.com/problems/implement-trie-prefix-tree/
// Run (Java 11+): java P0208ImplementTriePrefixTree.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0208ImplementTriePrefixTree {
    public static void main(String[] args) {
        Trie t = new Trie();
        t.insert("apple");
        check(t.search("apple") && !t.search("app") && t.startsWith("app"), true);
        t.insert("app");
        check(t.search("app"), true);
        if (fails > 0) { System.out.println(fails + " test(s) FAILED"); System.exit(1); }
        System.out.println("OK");
    }

    static int fails = 0;

    static void check(Object got, Object expected) {
        String g = str(got), e = str(expected);
        if (!g.equals(e)) { fails++; System.out.println("FAIL: got " + g + ", expected " + e); }
    }

    static String str(Object o) {
        if (o == null) return "null";
        if (o instanceof int[]) return Arrays.toString((int[]) o);
        if (o instanceof long[]) return Arrays.toString((long[]) o);
        if (o instanceof double[]) return Arrays.toString((double[]) o);
        if (o instanceof char[]) return Arrays.toString((char[]) o);
        if (o instanceof boolean[]) return Arrays.toString((boolean[]) o);
        if (o instanceof Object[]) return Arrays.deepToString((Object[]) o);
        return String.valueOf(o);
    }
}

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
