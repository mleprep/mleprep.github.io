// 224. Basic Calculator (Hard)
// https://leetcode.com/problems/basic-calculator/
// Run (Java 11+): java P0224BasicCalculator.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0224BasicCalculator {
    public static void main(String[] args) {
        check(new Solution().calculate("(1+(4+5+2)-3)+(6+8)"), 23);
        check(new Solution().calculate("-(2+3)"), -5);
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
    public int calculate(String s) {
        int res = 0, num = 0, sign = 1;
        Deque<Integer> st = new ArrayDeque<>();
        for (char c : s.toCharArray()) {
            if (Character.isDigit(c)) num = num * 10 + (c - '0');
            else if (c == '+' || c == '-') { res += sign * num; num = 0; sign = c == '+' ? 1 : -1; }
            else if (c == '(') { st.push(res); st.push(sign); res = 0; sign = 1; }
            else if (c == ')') {
                res += sign * num; num = 0;
                res = st.pop() * res + st.pop();   // pop sign, then previous result
            }
        }
        return res + sign * num;
    }
}
