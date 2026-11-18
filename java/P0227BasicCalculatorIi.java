// 227. Basic Calculator II (Medium)
// https://leetcode.com/problems/basic-calculator-ii/
// Run (Java 11+): java P0227BasicCalculatorIi.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0227BasicCalculatorIi {
    public static void main(String[] args) {
        check(new Solution().calculate("3+2*2"), 7);
        check(new Solution().calculate("14-3/2"), 13);
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
        Deque<Integer> st = new ArrayDeque<>();
        int num = 0;
        char op = '+';
        for (int i = 0; i < s.length(); i++) {
            char c = s.charAt(i);
            if (Character.isDigit(c)) num = num * 10 + (c - '0');
            if ("+-*/".indexOf(c) >= 0 || i == s.length() - 1) {
                if (op == '+') st.push(num);
                else if (op == '-') st.push(-num);
                else if (op == '*') st.push(st.pop() * num);
                else st.push(st.pop() / num);
                op = c; num = 0;
            }
        }
        int res = 0;
        for (int x : st) res += x;
        return res;
    }
}
