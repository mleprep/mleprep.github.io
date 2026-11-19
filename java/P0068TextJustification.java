// 68. Text Justification (Hard)
// https://leetcode.com/problems/text-justification/
// Run (Java 11+): java P0068TextJustification.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0068TextJustification {
    public static void main(String[] args) {
        check(new Solution().fullJustify(new String[]{"This", "is", "an", "example", "of", "text", "justification."}, 16),
              List.of("This    is    an", "example  of text", "justification.  "));
        check(new Solution().fullJustify(new String[]{"What", "must", "be", "acknowledgment", "shall", "be"}, 16),
              List.of("What   must   be", "acknowledgment  ", "shall be        "));
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
    public List<String> fullJustify(String[] words, int maxWidth) {
        List<String> res = new ArrayList<>();
        int i = 0;
        while (i < words.length) {
            int j = i, letters = 0;
            while (j < words.length && letters + words[j].length() + (j - i) <= maxWidth) letters += words[j++].length();
            StringBuilder sb = new StringBuilder(words[i]);
            int gaps = j - i - 1;
            if (j == words.length || gaps == 0) {              // last line or single word: left-justify
                for (int k = i + 1; k < j; k++) sb.append(' ').append(words[k]);
                sb.append(" ".repeat(maxWidth - sb.length()));
            } else {
                int spaces = maxWidth - letters, each = spaces / gaps, extra = spaces % gaps;
                for (int k = i + 1; k < j; k++) sb.append(" ".repeat(each + (k - i <= extra ? 1 : 0))).append(words[k]);
            }
            res.add(sb.toString());
            i = j;
        }
        return res;
    }
}
