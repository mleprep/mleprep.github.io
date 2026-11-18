// 127. Word Ladder (Hard)
// https://leetcode.com/problems/word-ladder/
// Run (Java 11+): java P0127WordLadder.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0127WordLadder {
    public static void main(String[] args) {
        check(new Solution().ladderLength("hit", "cog", Arrays.asList("hot", "dot", "dog", "lot", "log", "cog")), 5);
        check(new Solution().ladderLength("hit", "cog", Arrays.asList("hot", "dot", "dog", "lot", "log")), 0);
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
    public int ladderLength(String beginWord, String endWord, List<String> wordList) {
        Set<String> words = new HashSet<>(wordList);
        if (!words.contains(endWord)) return 0;
        Deque<String> q = new ArrayDeque<>(List.of(beginWord));
        for (int steps = 1; !q.isEmpty(); steps++) {
            for (int k = q.size(); k > 0; k--) {
                String w = q.poll();
                if (w.equals(endWord)) return steps;
                char[] cs = w.toCharArray();
                for (int i = 0; i < cs.length; i++) {
                    char orig = cs[i];
                    for (char ch = 'a'; ch <= 'z'; ch++) {
                        cs[i] = ch;
                        String nxt = new String(cs);
                        if (words.remove(nxt)) q.add(nxt);
                    }
                    cs[i] = orig;
                }
            }
        }
        return 0;
    }
}
