// 282. Expression Add Operators (Hard)
// https://leetcode.com/problems/expression-add-operators/
// Run (Java 11+): java P0282ExpressionAddOperators.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0282ExpressionAddOperators {
    public static void main(String[] args) {
        List<String> r = new Solution().addOperators("232", 8);
        Collections.sort(r);
        check(r, "[2*3+2, 2+3*2]");
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
    public List<String> addOperators(String num, int target) {
        List<String> res = new ArrayList<>();
        bt(num, target, 0, "", 0, 0, res);
        return res;
    }

    private void bt(String num, int target, int i, String expr, long value, long prev, List<String> res) {
        if (i == num.length()) { if (value == target) res.add(expr); return; }
        for (int j = i + 1; j <= num.length(); j++) {
            String s = num.substring(i, j);
            if (s.length() > 1 && s.charAt(0) == '0') break;
            long cur = Long.parseLong(s);
            if (i == 0) bt(num, target, j, s, cur, cur, res);
            else {
                bt(num, target, j, expr + "+" + s, value + cur, cur, res);
                bt(num, target, j, expr + "-" + s, value - cur, -cur, res);
                bt(num, target, j, expr + "*" + s, value - prev + prev * cur, prev * cur, res);
            }
        }
    }
}
