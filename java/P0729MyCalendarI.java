// 729. My Calendar I (Medium)
// https://leetcode.com/problems/my-calendar-i/
// Run (Java 11+): java P0729MyCalendarI.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0729MyCalendarI {
    public static void main(String[] args) {
        MyCalendar c = new MyCalendar();
        check(c.book(10, 20), true);
        check(c.book(15, 25), false);
        check(c.book(20, 30), true);
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

class MyCalendar {
    private final TreeMap<Integer, Integer> cal = new TreeMap<>();   // start -> end

    public boolean book(int start, int end) {
        Integer prev = cal.floorKey(start), next = cal.ceilingKey(start);
        if (prev != null && cal.get(prev) > start) return false;
        if (next != null && next < end) return false;
        cal.put(start, end);
        return true;
    }
}
