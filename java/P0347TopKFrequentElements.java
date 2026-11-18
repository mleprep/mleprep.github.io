// 347. Top K Frequent Elements (Medium)
// https://leetcode.com/problems/top-k-frequent-elements/
// Run (Java 11+): java P0347TopKFrequentElements.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0347TopKFrequentElements {
    public static void main(String[] args) {
        int[] t = new Solution().topKFrequent(new int[]{1, 1, 1, 2, 2, 3}, 2);
        Arrays.sort(t);
        check(t, new int[]{1, 2});
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
    public int[] topKFrequent(int[] nums, int k) {
        Map<Integer, Integer> cnt = new HashMap<>();
        for (int x : nums) cnt.merge(x, 1, Integer::sum);
        List<List<Integer>> bucket = new ArrayList<>();
        for (int i = 0; i <= nums.length; i++) bucket.add(new ArrayList<>());
        for (Map.Entry<Integer, Integer> e : cnt.entrySet()) bucket.get(e.getValue()).add(e.getKey());
        int[] res = new int[k];
        int i = 0;
        for (int f = nums.length; f > 0 && i < k; f--)
            for (int x : bucket.get(f)) if (i < k) res[i++] = x;
        return res;
    }
}
