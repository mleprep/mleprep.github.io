// 841. Keys and Rooms (Medium)
// https://leetcode.com/problems/keys-and-rooms/
// Run (Java 11+): java P0841KeysAndRooms.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0841KeysAndRooms {
    public static void main(String[] args) {
        check(new Solution().canVisitAllRooms(List.of(List.of(1), List.of(2), List.of(3), List.of())), true);
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
    public boolean canVisitAllRooms(List<List<Integer>> rooms) {
        boolean[] seen = new boolean[rooms.size()];
        seen[0] = true;
        Deque<Integer> st = new ArrayDeque<>(List.of(0));
        int count = 1;
        while (!st.isEmpty())
            for (int k : rooms.get(st.pop()))
                if (!seen[k]) { seen[k] = true; count++; st.push(k); }
        return count == rooms.size();
    }
}
