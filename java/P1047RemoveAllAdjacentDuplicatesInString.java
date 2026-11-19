// 1047. Remove All Adjacent Duplicates In String (Easy)
// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/
// Run (Java 11+): java P1047RemoveAllAdjacentDuplicatesInString.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1047RemoveAllAdjacentDuplicatesInString {
    public static void main(String[] args) {
        check(new Solution().removeDuplicates("azxxzy"), "ay");
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
    public String removeDuplicates(String s) {
        StringBuilder st = new StringBuilder();
        for (char c : s.toCharArray()) {
            if (st.length() > 0 && st.charAt(st.length() - 1) == c) st.deleteCharAt(st.length() - 1);
            else st.append(c);
        }
        return st.toString();
    }
}
