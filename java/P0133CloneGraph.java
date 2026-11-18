// 133. Clone Graph (Medium)
// https://leetcode.com/problems/clone-graph/
// Run (Java 11+): java P0133CloneGraph.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0133CloneGraph {
    public static void main(String[] args) {
        Node a = new Node(1), b = new Node(2);
        a.neighbors.add(b); b.neighbors.add(a);
        Node c = new Solution().cloneGraph(a);
        check(c != a && c.neighbors.get(0).val == 2 && c.neighbors.get(0).neighbors.get(0) == c, true);
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

class Node { // provided by LeetCode
    public int val;
    public List<Node> neighbors = new ArrayList<>();
    public Node(int val) { this.val = val; }
}

class Solution {
    private final Map<Node, Node> clones = new HashMap<>();

    public Node cloneGraph(Node node) {
        if (node == null) return null;
        if (clones.containsKey(node)) return clones.get(node);
        Node c = new Node(node.val);
        clones.put(node, c);
        for (Node nb : node.neighbors) c.neighbors.add(cloneGraph(nb));
        return c;
    }
}
