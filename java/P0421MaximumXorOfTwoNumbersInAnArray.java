// 421. Maximum XOR of Two Numbers in an Array (Medium)
// https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/
// Run (Java 11+): java P0421MaximumXorOfTwoNumbersInAnArray.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0421MaximumXorOfTwoNumbersInAnArray {
    public static void main(String[] args) {
        check(new Solution().findMaximumXOR(new int[]{3, 10, 5, 25, 2, 8}), 28);
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
    public int findMaximumXOR(int[] nums) {
        int[][] trie = new int[nums.length * 31 + 1][2];   // array-backed trie, 0 = no child
        int size = 1, best = 0;
        for (int x : nums) {
            int n = 0;
            for (int b = 30; b >= 0; b--) {
                int bit = (x >> b) & 1;
                if (trie[n][bit] == 0) trie[n][bit] = size++;
                n = trie[n][bit];
            }
        }
        for (int x : nums) {
            int n = 0, cur = 0;
            for (int b = 30; b >= 0; b--) {
                int bit = (x >> b) & 1;
                if (trie[n][1 - bit] != 0) { cur |= 1 << b; n = trie[n][1 - bit]; }
                else n = trie[n][bit];
            }
            best = Math.max(best, cur);
        }
        return best;
    }
}
