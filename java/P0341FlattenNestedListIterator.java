// 341. Flatten Nested List Iterator (Medium)
// https://leetcode.com/problems/flatten-nested-list-iterator/
// Run (Java 11+): java P0341FlattenNestedListIterator.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0341FlattenNestedListIterator {
    public static void main(String[] args) {
        NestedIterator it = new NestedIterator(List.of(new NI(new NI(1), new NI(1)), new NI(2), new NI(new NI(1), new NI(1))));
        List<Integer> out = new ArrayList<>();
        while (it.hasNext()) out.add(it.next());
        check(out, "[1, 1, 2, 1, 1]");
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

interface NestedInteger { // provided by LeetCode
    boolean isInteger();
    Integer getInteger();
    List<NestedInteger> getList();
}

class NI implements NestedInteger { // local stand-in for tests
    private final Integer v; private final List<NestedInteger> l;
    NI(int v) { this.v = v; this.l = null; }
    NI(NestedInteger... xs) { this.v = null; this.l = List.of(xs); }
    public boolean isInteger() { return v != null; }
    public Integer getInteger() { return v; }
    public List<NestedInteger> getList() { return l; }
}

class NestedIterator implements Iterator<Integer> {
    private final Deque<NestedInteger> st = new ArrayDeque<>();

    public NestedIterator(List<NestedInteger> nestedList) {
        for (int i = nestedList.size() - 1; i >= 0; i--) st.push(nestedList.get(i));
    }

    @Override
    public Integer next() { hasNext(); return st.pop().getInteger(); }

    @Override
    public boolean hasNext() {
        while (!st.isEmpty() && !st.peek().isInteger()) {
            List<NestedInteger> l = st.pop().getList();
            for (int i = l.size() - 1; i >= 0; i--) st.push(l.get(i));
        }
        return !st.isEmpty();
    }
}
