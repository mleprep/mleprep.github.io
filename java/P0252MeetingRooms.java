// 252. Meeting Rooms (Easy)
// https://leetcode.com/problems/meeting-rooms/
// Run (Java 11+): java P0252MeetingRooms.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0252MeetingRooms {
    public static void main(String[] args) {
        check(new Solution().canAttendMeetings(new int[][]{{0, 30}, {5, 10}, {15, 20}}), false);
        check(new Solution().canAttendMeetings(new int[][]{{7, 10}, {2, 4}}), true);
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
    public boolean canAttendMeetings(int[][] intervals) {
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[0], b[0]));
        for (int i = 1; i < intervals.length; i++) if (intervals[i][0] < intervals[i - 1][1]) return false;
        return true;
    }
}
