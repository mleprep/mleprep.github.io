// 345. Reverse Vowels of a String (Easy)
// https://leetcode.com/problems/reverse-vowels-of-a-string/
// Run (Java 11+): java P0345ReverseVowelsOfAString.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0345ReverseVowelsOfAString {
    public static void main(String[] args) {
        check(new Solution().reverseVowels("IceCreAm"), "AceCreIm");
        check(new Solution().reverseVowels("leetcode"), "leotcede");
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
    public String reverseVowels(String s) {
        char[] a = s.toCharArray();
        for (int i = 0, j = a.length - 1; i < j; ) {
            if ("aeiouAEIOU".indexOf(a[i]) < 0) i++;
            else if ("aeiouAEIOU".indexOf(a[j]) < 0) j--;
            else { char t = a[i]; a[i++] = a[j]; a[j--] = t; }
        }
        return new String(a);
    }
}
