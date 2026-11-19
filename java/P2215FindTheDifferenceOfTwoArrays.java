// 2215. Find the Difference of Two Arrays (Easy)
// https://leetcode.com/problems/find-the-difference-of-two-arrays/
// Run (Java 11+): java P2215FindTheDifferenceOfTwoArrays.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P2215FindTheDifferenceOfTwoArrays {
    public static void main(String[] args) {
        List<List<Integer>> d = new Solution().findDifference(new int[]{1, 2, 3}, new int[]{2, 4, 6});
        for (List<Integer> l : d) Collections.sort(l);
        check(d, List.of(List.of(1, 3), List.of(4, 6)));
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
    public List<List<Integer>> findDifference(int[] nums1, int[] nums2) {
        Set<Integer> a = new HashSet<>(), b = new HashSet<>();
        for (int x : nums1) a.add(x);
        for (int x : nums2) b.add(x);
        List<Integer> onlyA = new ArrayList<>(), onlyB = new ArrayList<>();
        for (int x : a) if (!b.contains(x)) onlyA.add(x);
        for (int x : b) if (!a.contains(x)) onlyB.add(x);
        return List.of(onlyA, onlyB);
    }
}
