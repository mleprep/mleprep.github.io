// 368. Largest Divisible Subset (Medium)
// https://leetcode.com/problems/largest-divisible-subset/
// Run (Java 11+): java P0368LargestDivisibleSubset.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0368LargestDivisibleSubset {
    public static void main(String[] args) {
        check(new Solution().largestDivisibleSubset(new int[]{1, 2, 4, 8}), "[1, 2, 4, 8]");
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
    public List<Integer> largestDivisibleSubset(int[] nums) {
        Arrays.sort(nums);
        int n = nums.length, best = 0;
        int[] size = new int[n], parent = new int[n];
        for (int i = 0; i < n; i++) {
            size[i] = 1; parent[i] = -1;
            for (int j = 0; j < i; j++)
                if (nums[i] % nums[j] == 0 && size[j] + 1 > size[i]) { size[i] = size[j] + 1; parent[i] = j; }
            if (size[i] > size[best]) best = i;
        }
        LinkedList<Integer> res = new LinkedList<>();
        for (int i = best; i != -1; i = parent[i]) res.addFirst(nums[i]);
        return res;
    }
}
