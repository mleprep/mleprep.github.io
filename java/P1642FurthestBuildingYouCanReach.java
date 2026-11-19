// 1642. Furthest Building You Can Reach (Medium)
// https://leetcode.com/problems/furthest-building-you-can-reach/
// Run (Java 11+): java P1642FurthestBuildingYouCanReach.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1642FurthestBuildingYouCanReach {
    public static void main(String[] args) {
        check(new Solution().furthestBuilding(new int[]{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2), 7);
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
    public int furthestBuilding(int[] heights, int bricks, int ladders) {
        PriorityQueue<Integer> pq = new PriorityQueue<>();
        for (int i = 0; i < heights.length - 1; i++) {
            int d = heights[i + 1] - heights[i];
            if (d <= 0) continue;
            pq.add(d);
            if (pq.size() > ladders && (bricks -= pq.poll()) < 0) return i;
        }
        return heights.length - 1;
    }
}
