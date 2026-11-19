// 1396. Design Underground System (Medium)
// https://leetcode.com/problems/design-underground-system/
// Run (Java 11+): java P1396DesignUndergroundSystem.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1396DesignUndergroundSystem {
    public static void main(String[] args) {
        UndergroundSystem u = new UndergroundSystem();
        u.checkIn(45, "Leyton", 3); u.checkIn(27, "Leyton", 10);
        u.checkOut(45, "Waterloo", 15); u.checkOut(27, "Waterloo", 20);
        check(u.getAverageTime("Leyton", "Waterloo"), 11.0);
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

class UndergroundSystem {
    private final Map<Integer, Object[]> active = new HashMap<>();   // id -> {station, time}
    private final Map<String, long[]> stats = new HashMap<>();       // "a->b" -> {total, trips}

    public void checkIn(int id, String stationName, int t) { active.put(id, new Object[]{stationName, t}); }

    public void checkOut(int id, String stationName, int t) {
        Object[] in = active.remove(id);
        long[] s = stats.computeIfAbsent(in[0] + "->" + stationName, k -> new long[2]);
        s[0] += t - (int) in[1];
        s[1]++;
    }

    public double getAverageTime(String startStation, String endStation) {
        long[] s = stats.get(startStation + "->" + endStation);
        return (double) s[0] / s[1];
    }
}
