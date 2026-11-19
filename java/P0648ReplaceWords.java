// 648. Replace Words (Medium)
// https://leetcode.com/problems/replace-words/
// Run (Java 11+): java P0648ReplaceWords.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0648ReplaceWords {
    public static void main(String[] args) {
        check(new Solution().replaceWords(List.of("cat", "bat", "rat"), "the cattle was rattled by the battery"), "the cat was rat by the bat");
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
