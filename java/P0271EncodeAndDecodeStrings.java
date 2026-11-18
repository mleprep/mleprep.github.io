// 271. Encode and Decode Strings (Medium)
// https://leetcode.com/problems/encode-and-decode-strings/
// Run (Java 11+): java P0271EncodeAndDecodeStrings.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0271EncodeAndDecodeStrings {
    public static void main(String[] args) {
        Codec c = new Codec();
        List<String> data = Arrays.asList("we", "say", ":", "yes", "#4#", "");
        check(c.decode(c.encode(data)), data);
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

class Codec {
    public String encode(List<String> strs) {
        StringBuilder sb = new StringBuilder();
        for (String s : strs) sb.append(s.length()).append('#').append(s);
        return sb.toString();
    }

    public List<String> decode(String s) {
        List<String> res = new ArrayList<>();
        int i = 0;
        while (i < s.length()) {
            int j = s.indexOf('#', i);
            int n = Integer.parseInt(s.substring(i, j));
            res.add(s.substring(j + 1, j + 1 + n));
            i = j + 1 + n;
        }
        return res;
    }
}
