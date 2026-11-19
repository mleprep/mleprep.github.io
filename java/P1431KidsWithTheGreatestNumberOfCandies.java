// 1431. Kids With the Greatest Number of Candies (Easy)
// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
// Run (Java 11+): java P1431KidsWithTheGreatestNumberOfCandies.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1431KidsWithTheGreatestNumberOfCandies {
    public static void main(String[] args) {
        check(new Solution().kidsWithCandies(new int[]{2, 3, 5, 1, 3}, 3), List.of(true, true, true, false, true));
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
    public List<Boolean> kidsWithCandies(int[] candies, int extraCandies) {
        int best = Arrays.stream(candies).max().getAsInt();
        List<Boolean> res = new ArrayList<>();
        for (int c : candies) res.add(c + extraCandies >= best);
        return res;
    }
}
