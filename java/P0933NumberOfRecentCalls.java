// 933. Number of Recent Calls (Easy)
// https://leetcode.com/problems/number-of-recent-calls/
// Run (Java 11+): java P0933NumberOfRecentCalls.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0933NumberOfRecentCalls {
    public static void main(String[] args) {
        RecentCounter rc = new RecentCounter();
        check(new int[]{rc.ping(1), rc.ping(100), rc.ping(3001), rc.ping(3002)}, new int[]{1, 2, 3, 3});
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

class RecentCounter {
    private final Deque<Integer> q = new ArrayDeque<>();

    public int ping(int t) {
        q.add(t);
        while (q.peek() < t - 3000) q.poll();
        return q.size();
    }
}
