// 698. Partition to K Equal Sum Subsets (Medium)
// https://leetcode.com/problems/partition-to-k-equal-sum-subsets/
// Run (Java 11+): java P0698PartitionToKEqualSumSubsets.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0698PartitionToKEqualSumSubsets {
    public static void main(String[] args) {
        check(new Solution().canPartitionKSubsets(new int[]{4, 3, 2, 3, 5, 2, 1}, 4), true);
        check(new Solution().canPartitionKSubsets(new int[]{1, 2, 3, 4}, 3), false);
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
    public boolean canPartitionKSubsets(int[] nums, int k) {
        int total = 0;
        for (int x : nums) total += x;
        if (total % k != 0) return false;
        Arrays.sort(nums);
        int target = total / k;
        if (nums[nums.length - 1] > target) return false;
        return bt(nums, nums.length - 1, new int[k], target);   // largest first
    }

    private boolean bt(int[] nums, int i, int[] buckets, int target) {
        if (i < 0) return true;
        for (int b = 0; b < buckets.length; b++) {
            if (buckets[b] + nums[i] <= target) {
                buckets[b] += nums[i];
                if (bt(nums, i - 1, buckets, target)) return true;
                buckets[b] -= nums[i];
            }
            if (buckets[b] == 0) break;
        }
        return false;
    }
}
