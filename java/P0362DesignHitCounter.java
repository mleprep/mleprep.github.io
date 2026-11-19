// 362. Design Hit Counter (Medium)
// https://leetcode.com/problems/design-hit-counter/
// Run (Java 11+): java P0362DesignHitCounter.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0362DesignHitCounter {
    public static void main(String[] args) {
        HitCounter hc = new HitCounter();
        hc.hit(1); hc.hit(2); hc.hit(3);
        check(hc.getHits(4), 3);
        hc.hit(300);
        check(hc.getHits(300), 4);
        check(hc.getHits(301), 3);
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

class HitCounter {
    private final int[] times = new int[300], counts = new int[300];   // ring of one-second buckets

    public void hit(int timestamp) {
        int i = timestamp % 300;
        if (times[i] != timestamp) { times[i] = timestamp; counts[i] = 0; }
        counts[i]++;
    }

    public int getHits(int timestamp) {
        int total = 0;
        for (int i = 0; i < 300; i++) if (timestamp - times[i] < 300) total += counts[i];
        return total;
    }
}
