// 2402. Meeting Rooms III (Hard)
// https://leetcode.com/problems/meeting-rooms-iii/
// Run (Java 11+): java P2402MeetingRoomsIii.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P2402MeetingRoomsIii {
    public static void main(String[] args) {
        check(new Solution().mostBooked(2, new int[][]{{0, 10}, {1, 5}, {2, 7}, {3, 4}}), 0);
        check(new Solution().mostBooked(3, new int[][]{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}}), 1);
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
    public int mostBooked(int n, int[][] meetings) {
        Arrays.sort(meetings, (a, b) -> Integer.compare(a[0], b[0]));
        PriorityQueue<Integer> free = new PriorityQueue<>();
        for (int i = 0; i < n; i++) free.add(i);
        PriorityQueue<long[]> busy = new PriorityQueue<>((a, b) -> a[0] != b[0] ? Long.compare(a[0], b[0]) : Long.compare(a[1], b[1]));
        int[] count = new int[n];
        for (int[] m : meetings) {
            while (!busy.isEmpty() && busy.peek()[0] <= m[0]) free.add((int) busy.poll()[1]);
            int room;
            if (!free.isEmpty()) { room = free.poll(); busy.add(new long[]{m[1], room}); }
            else { long[] b = busy.poll(); room = (int) b[1]; busy.add(new long[]{b[0] + m[1] - m[0], room}); }
            count[room]++;
        }
        int best = 0;
        for (int i = 1; i < n; i++) if (count[i] > count[best]) best = i;
        return best;
    }
}
