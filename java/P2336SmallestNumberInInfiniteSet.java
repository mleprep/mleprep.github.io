// 2336. Smallest Number in Infinite Set (Medium)
// https://leetcode.com/problems/smallest-number-in-infinite-set/
// Run (Java 11+): java P2336SmallestNumberInInfiniteSet.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P2336SmallestNumberInInfiniteSet {
    public static void main(String[] args) {
        SmallestInfiniteSet s = new SmallestInfiniteSet();
        s.addBack(2);
        check(new int[]{s.popSmallest(), s.popSmallest(), s.popSmallest()}, new int[]{1, 2, 3});
        s.addBack(1);
        check(new int[]{s.popSmallest(), s.popSmallest(), s.popSmallest()}, new int[]{1, 4, 5});
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

class SmallestInfiniteSet {
    private int nxt = 1;
    private final TreeSet<Integer> back = new TreeSet<>();   // re-added numbers below nxt

    public int popSmallest() { return back.isEmpty() ? nxt++ : back.pollFirst(); }

    public void addBack(int num) { if (num < nxt) back.add(num); }
}
