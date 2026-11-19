// 427. Construct Quad Tree (Medium)
// https://leetcode.com/problems/construct-quad-tree/
// Run (Java 11+): java P0427ConstructQuadTree.java
import java.util.*;
import java.util.function.*;
import java.util.stream.*;

public class P0427ConstructQuadTree {
    public static void main(String[] args) {
        StringBuilder sb = new StringBuilder();
        Deque<Node> q = new ArrayDeque<>(List.of(new Solution().construct(new int[][]{{1, 1, 0, 0}, {1, 1, 0, 0}, {1, 1, 1, 1}, {1, 1, 1, 1}})));
        while (!q.isEmpty()) {
            Node n = q.poll();
            sb.append(n.isLeaf ? (n.val ? "1" : "0") : "*");
            if (!n.isLeaf) q.addAll(List.of(n.topLeft, n.topRight, n.bottomLeft, n.bottomRight));
        }
        check(sb.toString(), "*1011");
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
    public boolean val, isLeaf;
    public Node topLeft, topRight, bottomLeft, bottomRight;
    public Node() {}
    public Node(boolean val, boolean isLeaf) { this.val = val; this.isLeaf = isLeaf; }
    public Node(boolean val, boolean isLeaf, Node topLeft, Node topRight, Node bottomLeft, Node bottomRight) {
        this.val = val; this.isLeaf = isLeaf;
        this.topLeft = topLeft; this.topRight = topRight; this.bottomLeft = bottomLeft; this.bottomRight = bottomRight;
    }
}

class Solution {
    public Node construct(int[][] grid) { return build(grid, 0, 0, grid.length); }

    private Node build(int[][] g, int r, int c, int n) {
        if (n == 1) return new Node(g[r][c] == 1, true);
        int h = n / 2;
        Node tl = build(g, r, c, h), tr = build(g, r, c + h, h), bl = build(g, r + h, c, h), br = build(g, r + h, c + h, h);
        if (tl.isLeaf && tr.isLeaf && bl.isLeaf && br.isLeaf && tl.val == tr.val && tr.val == bl.val && bl.val == br.val)
            return new Node(tl.val, true);
        return new Node(true, false, tl, tr, bl, br);
    }
}
