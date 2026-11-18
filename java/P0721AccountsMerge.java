// 721. Accounts Merge (Medium)
// https://leetcode.com/problems/accounts-merge/
// Run (Java 11+): java P0721AccountsMerge.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0721AccountsMerge {
    public static void main(String[] args) {
        List<List<String>> acc = List.of(List.of("John", "a@x", "b@x"), List.of("John", "a@x", "c@x"), List.of("Mary", "m@x"));
        List<List<String>> r = new Solution().accountsMerge(acc);
        r.sort((p, q) -> p.get(1).compareTo(q.get(1)));
        check(r, "[[John, a@x, b@x, c@x], [Mary, m@x]]");
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
    private final Map<String, String> parent = new HashMap<>();

    private String find(String x) {
        parent.putIfAbsent(x, x);
        while (!parent.get(x).equals(x)) {
            parent.put(x, parent.get(parent.get(x)));
            x = parent.get(x);
        }
        return x;
    }

    public List<List<String>> accountsMerge(List<List<String>> accounts) {
        Map<String, String> owner = new HashMap<>();
        for (List<String> acc : accounts)
            for (int i = 1; i < acc.size(); i++) {
                owner.put(acc.get(i), acc.get(0));
                parent.put(find(acc.get(i)), find(acc.get(1)));
            }
        Map<String, TreeSet<String>> groups = new HashMap<>();
        for (String e : owner.keySet()) groups.computeIfAbsent(find(e), k -> new TreeSet<>()).add(e);
        List<List<String>> res = new ArrayList<>();
        for (Map.Entry<String, TreeSet<String>> g : groups.entrySet()) {
            List<String> row = new ArrayList<>();
            row.add(owner.get(g.getKey()));
            row.addAll(g.getValue());
            res.add(row);
        }
        return res;
    }
}
