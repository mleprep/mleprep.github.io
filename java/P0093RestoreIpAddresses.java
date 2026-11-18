// 93. Restore IP Addresses (Medium)
// https://leetcode.com/problems/restore-ip-addresses/
// Run (Java 11+): java P0093RestoreIpAddresses.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0093RestoreIpAddresses {
    public static void main(String[] args) {
        check(new Solution().restoreIpAddresses("25525511135"), "[255.255.11.135, 255.255.111.35]");
        check(new Solution().restoreIpAddresses("0000"), "[0.0.0.0]");
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
    public List<String> restoreIpAddresses(String s) {
        List<String> res = new ArrayList<>();
        bt(s, 0, new ArrayList<>(), res);
        return res;
    }

    private void bt(String s, int i, List<String> parts, List<String> res) {
        int rem = 4 - parts.size(), left = s.length() - i;
        if (rem == 0) { if (left == 0) res.add(String.join(".", parts)); return; }
        if (left < rem || left > 3 * rem) return;
        for (int L = 1; L <= 3 && i + L <= s.length(); L++) {
            String p = s.substring(i, i + L);
            if ((p.charAt(0) == '0' && L > 1) || Integer.parseInt(p) > 255) break;
            parts.add(p);
            bt(s, i + L, parts, res);
            parts.remove(parts.size() - 1);
        }
    }
}
