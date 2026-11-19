// 2709. Greatest Common Divisor Traversal (Hard)
// https://leetcode.com/problems/greatest-common-divisor-traversal/
// Run (Java 11+): java P2709GreatestCommonDivisorTraversal.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P2709GreatestCommonDivisorTraversal {
    public static void main(String[] args) {
        check(new Solution().canTraverseAllPairs(new int[]{2, 3, 6}), true);
        check(new Solution().canTraverseAllPairs(new int[]{3, 9, 5}), false);
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
    public boolean canTraverseAllPairs(int[] nums) {
        int n = nums.length, M = 0;
        if (n == 1) return true;
        for (int x : nums) { if (x == 1) return false; M = Math.max(M, x); }
        int[] spf = new int[M + 1];
        for (int i = 2; i <= M; i++) if (spf[i] == 0) for (int j = i; j <= M; j += i) if (spf[j] == 0) spf[j] = i;
        int[] parent = new int[n];
        for (int i = 0; i < n; i++) parent[i] = i;
        Map<Integer, Integer> first = new HashMap<>();
        for (int i = 0; i < n; i++) {
            for (int x = nums[i]; x > 1; ) {
                int p = spf[x];
                Integer j = first.putIfAbsent(p, i);
                if (j != null) parent[find(parent, i)] = find(parent, j);
                while (x % p == 0) x /= p;
            }
        }
        int root = find(parent, 0);
        for (int i = 1; i < n; i++) if (find(parent, i) != root) return false;
        return true;
    }

    private int find(int[] p, int x) {
        while (p[x] != x) x = p[x] = p[p[x]];
        return x;
    }
}
