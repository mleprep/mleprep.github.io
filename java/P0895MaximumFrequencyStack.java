// 895. Maximum Frequency Stack (Hard)
// https://leetcode.com/problems/maximum-frequency-stack/
// Run (Java 11+): java P0895MaximumFrequencyStack.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0895MaximumFrequencyStack {
    public static void main(String[] args) {
        FreqStack fs = new FreqStack();
        for (int v : new int[]{5, 7, 5, 7, 4, 5}) fs.push(v);
        check(new int[]{fs.pop(), fs.pop(), fs.pop(), fs.pop()}, new int[]{5, 7, 5, 4});
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

class FreqStack {
    private final Map<Integer, Integer> freq = new HashMap<>();
    private final Map<Integer, Deque<Integer>> group = new HashMap<>();
    private int maxf = 0;

    public void push(int val) {
        int f = freq.merge(val, 1, Integer::sum);
        group.computeIfAbsent(f, k -> new ArrayDeque<>()).push(val);
        maxf = Math.max(maxf, f);
    }

    public int pop() {
        int val = group.get(maxf).pop();
        freq.merge(val, -1, Integer::sum);
        if (group.get(maxf).isEmpty()) maxf--;
        return val;
    }
}
