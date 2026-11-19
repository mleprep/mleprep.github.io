// 759. Employee Free Time (Hard)
// https://leetcode.com/problems/employee-free-time/
// Run (Java 11+): java P0759EmployeeFreeTime.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0759EmployeeFreeTime {
    public static void main(String[] args) {
        List<Interval> fr = new Solution().employeeFreeTime(List.of(
            List.of(new Interval(1, 3), new Interval(6, 7)), List.of(new Interval(2, 4)), List.of(new Interval(2, 5), new Interval(9, 12))));
        check(fr.stream().map(iv -> List.of(iv.start, iv.end)).collect(Collectors.toList()), List.of(List.of(5, 6), List.of(7, 9)));
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

class Interval { // provided by LeetCode
    public int start, end;
    public Interval() {}
    public Interval(int _start, int _end) { start = _start; end = _end; }
}

class Solution {
    public List<Interval> employeeFreeTime(List<List<Interval>> schedule) {
        List<Interval> all = new ArrayList<>();
        for (List<Interval> emp : schedule) all.addAll(emp);
        all.sort((a, b) -> Integer.compare(a.start, b.start));
        List<Interval> res = new ArrayList<>();
        int end = all.get(0).end;
        for (Interval iv : all) {
            if (iv.start > end) res.add(new Interval(end, iv.start));
            end = Math.max(end, iv.end);
        }
        return res;
    }
}
