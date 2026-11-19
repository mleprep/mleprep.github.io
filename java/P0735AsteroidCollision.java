// 735. Asteroid Collision (Medium)
// https://leetcode.com/problems/asteroid-collision/
// Run (Java 11+): java P0735AsteroidCollision.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0735AsteroidCollision {
    public static void main(String[] args) {
        check(new Solution().asteroidCollision(new int[]{10, 2, -5}), new int[]{10});
        check(new Solution().asteroidCollision(new int[]{-2, -1, 1, 2}), new int[]{-2, -1, 1, 2});
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
    public int[] asteroidCollision(int[] asteroids) {
        Deque<Integer> st = new ArrayDeque<>();
        for (int x : asteroids) {
            boolean alive = true;
            while (alive && x < 0 && !st.isEmpty() && st.peekLast() > 0) {
                int top = st.peekLast();
                if (top < -x) { st.pollLast(); continue; }
                if (top == -x) st.pollLast();
                alive = false;
            }
            if (alive) st.addLast(x);
        }
        return st.stream().mapToInt(Integer::intValue).toArray();
    }
}
