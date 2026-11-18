// 2013. Detect Squares (Medium)
// https://leetcode.com/problems/detect-squares/
// Run (Java 11+): java P2013DetectSquares.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P2013DetectSquares {
    public static void main(String[] args) {
        DetectSquares ds = new DetectSquares();
        ds.add(new int[]{3, 10}); ds.add(new int[]{11, 2}); ds.add(new int[]{3, 2});
        check(ds.count(new int[]{11, 10}), 1);
        ds.add(new int[]{11, 2});
        check(ds.count(new int[]{11, 10}), 2);
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

class DetectSquares {
    private final Map<Long, Integer> cnt = new HashMap<>();
    private final List<int[]> pts = new ArrayList<>();

    private long key(int x, int y) { return (long) x * 2001 + y; }

    public void add(int[] point) {
        if (cnt.merge(key(point[0], point[1]), 1, Integer::sum) == 1) pts.add(point);
    }

    public int count(int[] point) {
        int x = point[0], y = point[1], total = 0;
        for (int[] p : pts) {
            if (Math.abs(p[0] - x) != Math.abs(p[1] - y) || p[0] == x) continue;
            total += cnt.get(key(p[0], p[1])) * cnt.getOrDefault(key(x, p[1]), 0) * cnt.getOrDefault(key(p[0], y), 0);
        }
        return total;
    }
}
