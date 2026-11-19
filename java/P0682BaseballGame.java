// 682. Baseball Game (Easy)
// https://leetcode.com/problems/baseball-game/
// Run (Java 11+): java P0682BaseballGame.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0682BaseballGame {
    public static void main(String[] args) {
        check(new Solution().calPoints(new String[]{"5", "2", "C", "D", "+"}), 30);
        check(new Solution().calPoints(new String[]{"5", "-2", "4", "C", "D", "9", "+", "+"}), 27);
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
    public int calPoints(String[] operations) {
        Deque<Integer> st = new ArrayDeque<>();
        for (String op : operations) {
            if (op.equals("C")) st.pop();
            else if (op.equals("D")) st.push(2 * st.peek());
            else if (op.equals("+")) { int top = st.pop(), sum = top + st.peek(); st.push(top); st.push(sum); }
            else st.push(Integer.parseInt(op));
        }
        int total = 0;
        for (int v : st) total += v;
        return total;
    }
}
