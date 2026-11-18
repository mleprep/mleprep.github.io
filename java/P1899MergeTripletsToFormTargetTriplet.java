// 1899. Merge Triplets to Form Target Triplet (Medium)
// https://leetcode.com/problems/merge-triplets-to-form-target-triplet/
// Run (Java 11+): java P1899MergeTripletsToFormTargetTriplet.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1899MergeTripletsToFormTargetTriplet {
    public static void main(String[] args) {
        check(new Solution().mergeTriplets(new int[][]{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}}, new int[]{2, 7, 5}), true);
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
    public boolean mergeTriplets(int[][] triplets, int[] target) {
        boolean[] good = new boolean[3];
        for (int[] t : triplets) {
            if (t[0] > target[0] || t[1] > target[1] || t[2] > target[2]) continue;
            for (int i = 0; i < 3; i++) if (t[i] == target[i]) good[i] = true;
        }
        return good[0] && good[1] && good[2];
    }
}
