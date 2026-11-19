// 284. Peeking Iterator (Medium)
// https://leetcode.com/problems/peeking-iterator/
// Run (Java 11+): java P0284PeekingIterator.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0284PeekingIterator {
    public static void main(String[] args) {
        PeekingIterator pi = new PeekingIterator(List.of(1, 2, 3).iterator());
        check(pi.next(), 1);
        check(pi.peek(), 2);
        check(pi.next(), 2);
        check(pi.next(), 3);
        check(pi.hasNext(), false);
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

class PeekingIterator implements Iterator<Integer> {
    private final Iterator<Integer> it;
    private Integer nxt;

    public PeekingIterator(Iterator<Integer> iterator) {
        it = iterator;
        nxt = it.hasNext() ? it.next() : null;
    }

    public Integer peek() { return nxt; }

    @Override
    public Integer next() {
        Integer val = nxt;
        nxt = it.hasNext() ? it.next() : null;
        return val;
    }

    @Override
    public boolean hasNext() { return nxt != null; }
}
