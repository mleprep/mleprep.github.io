// 56. Merge Intervals (Medium)
// https://leetcode.com/problems/merge-intervals/
// Run (Java 11+): java P0056MergeIntervals.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0056MergeIntervals {
    public static void main(String[] args) {
        check(new Solution().merge(new int[][]{{1, 3}, {2, 6}, {8, 10}, {15, 18}}), new int[][]{{1, 6}, {8, 10}, {15, 18}});
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
    public int[][] merge(int[][] intervals) {
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[0], b[0]));
        List<int[]> res = new ArrayList<>();
        for (int[] iv : intervals) {
            if (!res.isEmpty() && iv[0] <= res.get(res.size() - 1)[1])
                res.get(res.size() - 1)[1] = Math.max(res.get(res.size() - 1)[1], iv[1]);
            else res.add(new int[]{iv[0], iv[1]});
        }
        return res.toArray(new int[0][]);
    }
}
