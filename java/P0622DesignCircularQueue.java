// 622. Design Circular Queue (Medium)
// https://leetcode.com/problems/design-circular-queue/
// Run (Java 11+): java P0622DesignCircularQueue.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0622DesignCircularQueue {
    public static void main(String[] args) {
        MyCircularQueue cq = new MyCircularQueue(3);
        check(cq.enQueue(1) && cq.enQueue(2) && cq.enQueue(3) && !cq.enQueue(4), true);
        check(cq.Rear(), 3);
        check(cq.deQueue() && cq.enQueue(4), true);
        check(cq.Rear(), 4);
        check(cq.Front(), 2);
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

class MyCircularQueue {
    private final int[] a;
    private int head = 0, n = 0;

    public MyCircularQueue(int k) { a = new int[k]; }

    public boolean enQueue(int value) {
        if (isFull()) return false;
        a[(head + n++) % a.length] = value;
        return true;
    }

    public boolean deQueue() {
        if (isEmpty()) return false;
        head = (head + 1) % a.length;
        n--;
        return true;
    }

    public int Front() { return isEmpty() ? -1 : a[head]; }
    public int Rear() { return isEmpty() ? -1 : a[(head + n - 1) % a.length]; }
    public boolean isEmpty() { return n == 0; }
    public boolean isFull() { return n == a.length; }
}
