// 253. Meeting Rooms II (Medium)
// https://leetcode.com/problems/meeting-rooms-ii/
// Run (Java 11+): java P0253MeetingRoomsIi.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0253MeetingRoomsIi {
    public static void main(String[] args) {
        check(new Solution().minMeetingRooms(new int[][]{{0, 30}, {5, 10}, {15, 20}}), 2);
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
    public int minMeetingRooms(int[][] intervals) {
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[0], b[0]));
        PriorityQueue<Integer> ends = new PriorityQueue<>();   // end times of rooms in use
        for (int[] iv : intervals) {
            if (!ends.isEmpty() && ends.peek() <= iv[0]) ends.poll();   // reuse the freed room
            ends.add(iv[1]);
        }
        return ends.size();
    }
}
