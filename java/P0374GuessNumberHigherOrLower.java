// 374. Guess Number Higher or Lower (Easy)
// https://leetcode.com/problems/guess-number-higher-or-lower/
// Run (Java 11+): java P0374GuessNumberHigherOrLower.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0374GuessNumberHigherOrLower {
    public static void main(String[] args) {
        GuessGame.pick = 1702766719;
        check(new Solution().guessNumber(2126753390), 1702766719);
        GuessGame.pick = 6;
        check(new Solution().guessNumber(10), 6);
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

class GuessGame { // provided by LeetCode
    static int pick;
    int guess(int num) { return Integer.compare(pick, num); }
}

class Solution extends GuessGame {
    public int guessNumber(int n) {
        int lo = 1, hi = n;
        while (true) {
            int mid = lo + (hi - lo) / 2;          // avoids int overflow
            int g = guess(mid);
            if (g == 0) return mid;
            if (g < 0) hi = mid - 1; else lo = mid + 1;
        }
    }
}
