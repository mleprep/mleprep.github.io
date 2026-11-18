// 232. Implement Queue using Stacks (Easy)
// https://leetcode.com/problems/implement-queue-using-stacks/
// Run (Java 11+): java P0232ImplementQueueUsingStacks.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0232ImplementQueueUsingStacks {
    public static void main(String[] args) {
        MyQueue q = new MyQueue();
        q.push(1); q.push(2);
        check(q.peek(), 1);
        check(q.pop(), 1);
        check(q.empty(), false);
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

class MyQueue {
    private final Deque<Integer> in = new ArrayDeque<>(), out = new ArrayDeque<>();

    public void push(int x) { in.push(x); }

    private void move() { if (out.isEmpty()) while (!in.isEmpty()) out.push(in.pop()); }

    public int pop() { move(); return out.pop(); }
    public int peek() { move(); return out.peek(); }
    public boolean empty() { return in.isEmpty() && out.isEmpty(); }
}
