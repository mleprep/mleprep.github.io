// 588. Design In-Memory File System (Hard)
// https://leetcode.com/problems/design-in-memory-file-system/
// Run (Java 11+): java P0588DesignInMemoryFileSystem.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0588DesignInMemoryFileSystem {
    public static void main(String[] args) {
        FileSystem fs = new FileSystem();
        check(fs.ls("/"), List.of());
        fs.mkdir("/a/b/c");
        fs.addContentToFile("/a/b/c/d", "hello");
        check(fs.ls("/"), List.of("a"));
        check(fs.ls("/a/b/c/d"), List.of("d"));
        check(fs.readContentFromFile("/a/b/c/d"), "hello");
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

class FileSystem {
    private static class Node { TreeMap<String, Node> kids = new TreeMap<>(); StringBuilder file; }
    private final Node root = new Node();

    private Node walk(String path) {
        Node node = root;
        for (String part : path.split("/")) if (!part.isEmpty()) node = node.kids.computeIfAbsent(part, k -> new Node());
        return node;
    }

    public List<String> ls(String path) {
        Node node = walk(path);
        if (node.file != null) return List.of(path.substring(path.lastIndexOf('/') + 1));
        return new ArrayList<>(node.kids.keySet());
    }

    public void mkdir(String path) { walk(path); }

    public void addContentToFile(String filePath, String content) {
        Node node = walk(filePath);
        if (node.file == null) node.file = new StringBuilder();
        node.file.append(content);
    }

    public String readContentFromFile(String filePath) { return walk(filePath).file.toString(); }
}
