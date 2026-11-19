// 1352. Product of the Last K Numbers (Medium)
// https://leetcode.com/problems/product-of-the-last-k-numbers/
// Run (Java 11+): java P1352ProductOfTheLastKNumbers.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P1352ProductOfTheLastKNumbers {
    public static void main(String[] args) {
        ProductOfNumbers p = new ProductOfNumbers();
        for (int x : new int[]{3, 0, 2, 5, 4}) p.add(x);
        check(p.getProduct(3), 40);
        check(p.getProduct(4), 0);
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

class ProductOfNumbers {
    private final List<Integer> pre = new ArrayList<>(List.of(1));

    public void add(int num) {
        if (num == 0) { pre.clear(); pre.add(1); }
        else pre.add(pre.get(pre.size() - 1) * num);
    }

    public int getProduct(int k) {
        int n = pre.size();
        return k >= n ? 0 : pre.get(n - 1) / pre.get(n - 1 - k);
    }
}
