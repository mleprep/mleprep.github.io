// 225. Implement Stack using Queues (Easy)
// https://leetcode.com/problems/implement-stack-using-queues/
// Run (Java 11+): java P0225ImplementStackUsingQueues.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0225ImplementStackUsingQueues {
    public static void main(String[] args) {
        MyStack s = new MyStack();
        s.push(1); s.push(2);
        check(s.top(), 2);
        check(s.pop(), 2);
        check(s.empty(), false);
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

class MyStack {
    private final Deque<Integer> q = new ArrayDeque<>();

    public void push(int x) {
        q.addLast(x);
        for (int i = 1; i < q.size(); i++) q.addLast(q.pollFirst());
    }

    public int pop() { return q.pollFirst(); }
    public int top() { return q.peekFirst(); }
    public boolean empty() { return q.isEmpty(); }
}
