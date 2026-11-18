// 380. Insert Delete GetRandom O(1) (Medium)
// https://leetcode.com/problems/insert-delete-getrandom-o1/
// Run (Java 11+): java P0380InsertDeleteGetrandomO1.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0380InsertDeleteGetrandomO1 {
    public static void main(String[] args) {
        RandomizedSet rs = new RandomizedSet();
        check(rs.insert(1) && !rs.remove(2) && rs.insert(2), true);
        check(rs.remove(1) && !rs.insert(2), true);
        check(rs.getRandom(), 2);
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

class RandomizedSet {
    private final List<Integer> vals = new ArrayList<>();
    private final Map<Integer, Integer> idx = new HashMap<>();
    private final Random rnd = new Random();

    public boolean insert(int val) {
        if (idx.containsKey(val)) return false;
        idx.put(val, vals.size());
        vals.add(val);
        return true;
    }

    public boolean remove(int val) {
        Integer i = idx.remove(val);
        if (i == null) return false;
        int last = vals.remove(vals.size() - 1);
        if (i < vals.size()) { vals.set(i, last); idx.put(last, i); }
        return true;
    }

    public int getRandom() { return vals.get(rnd.nextInt(vals.size())); }
}
