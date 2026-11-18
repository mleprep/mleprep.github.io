// 535. Encode and Decode TinyURL (Medium)
// https://leetcode.com/problems/encode-and-decode-tinyurl/
// Run (Java 11+): java P0535EncodeAndDecodeTinyurl.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0535EncodeAndDecodeTinyurl {
    public static void main(String[] args) {
        Codec c = new Codec();
        String u = "https://leetcode.com/problems/design-tinyurl";
        check(c.decode(c.encode(u)), u);
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
    private static final String ALPHA = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ";
    private final Map<String, String> longToShort = new HashMap<>(), shortToLong = new HashMap<>();

    public String encode(String longUrl) {
        String code = longToShort.get(longUrl);
        if (code == null) {
            long n = shortToLong.size() + 1;
            StringBuilder sb = new StringBuilder();
            do { sb.append(ALPHA.charAt((int) (n % 62))); n /= 62; } while (n > 0);
            code = sb.reverse().toString();
            longToShort.put(longUrl, code);
            shortToLong.put(code, longUrl);
        }
        return "http://tinyurl.com/" + code;
    }

    public String decode(String shortUrl) {
        return shortToLong.get(shortUrl.substring(shortUrl.lastIndexOf('/') + 1));
    }
}
