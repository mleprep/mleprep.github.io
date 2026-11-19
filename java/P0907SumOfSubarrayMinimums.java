// 907. Sum of Subarray Minimums (Medium)
// https://leetcode.com/problems/sum-of-subarray-minimums/
// Run (Java 11+): java P0907SumOfSubarrayMinimums.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0907SumOfSubarrayMinimums {
    public static void main(String[] args) {
        check(new Solution().sumSubarrayMins(new int[]{3, 1, 2, 4}), 17);
        check(new Solution().sumSubarrayMins(new int[]{11, 81, 94, 43, 3}), 444);
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
    public int sumSubarrayMins(int[] arr) {
        int n = arr.length;
        long MOD = 1_000_000_007L, res = 0;
        int[] left = new int[n], right = new int[n];
        Deque<Integer> st = new ArrayDeque<>();
        for (int i = 0; i < n; i++) {
            while (!st.isEmpty() && arr[st.peek()] > arr[i]) st.pop();
            left[i] = i - (st.isEmpty() ? -1 : st.peek());
            st.push(i);
        }
        st.clear();
        for (int i = n - 1; i >= 0; i--) {
            while (!st.isEmpty() && arr[st.peek()] >= arr[i]) st.pop();
            right[i] = (st.isEmpty() ? n : st.peek()) - i;
            st.push(i);
        }
        for (int i = 0; i < n; i++) res = (res + (long) arr[i] * left[i] % MOD * right[i]) % MOD;
        return (int) res;
    }
}
