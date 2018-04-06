package main

import (
    "fmt"
    "time"
    "math/rand"
)

var n int = 10000
var tree = make([][]int, n)
var LCATree = make([][]int, n + 1)
var par = make([]int, n)
var name = make([]string, n)
var fastpar = make([]int, n)

func add_edge(u, p int) { //(ID, parent ID)
    if (p == n) {
        LCATree[p] = append(LCATree[p], u)
        return
    }
    if (u != p) {
        tree[u] = append(tree[u], p)
        tree[p] = append(tree[p], u)
        LCATree[p] = append(LCATree[p], u)
    }
    par[u] = p
    fastpar[u] = p
}

func get_name(id int) string {
    return name[id]
}

func get_root(id int) int {
    if (fastpar[id] == id) {
        return id
    } else {
        fastpar[id] = get_root(fastpar[id])
        return fastpar[id]
    }
}

func get_ancestors(id int) []int {
    var ancestors = make([]int, 0, 1)
    cur := id
    for {
        ancestors = append(ancestors, cur)
        if (par[cur] == cur) {
            break
        }
        cur = par[cur]
    }
    return ancestors
}

func dfs(u int, descendants *[]int) {
    *descendants = append(*descendants, u)
    for i := 0; i < len(tree[u]); i ++ {
        v := tree[u][i]
        if (v != par[u]) {
            dfs(v, descendants)
        }
    }
    return
}

func get_descendants(id int) []int {
    var descendants = make([]int, 0, 1)
    dfs(id, &descendants)
    return descendants
}

//LCA
const LOGM = 30
var N = n + 1
var root = n
var depth = make([]int, N)
var parent = make([][]int, LOGM)
func dfs_lca(u, prev, d int) {
    parent[0][u] = prev
    depth[u] = d
    for i := 0; i < len(LCATree[u]); i ++ {
        dfs_lca(LCATree[u][i], u, d + 1)
    }
}
func build() {
    for i := 0; i < LOGM; i ++ {
        parent[i] = make([]int, N)
    }
    var cnt = make([]int, n)
    for i := 0; i < n; i ++ {
        cnt[i] = 0
    }
    for i := 0; i < n; i ++ {
        cnt[get_root(i)] ++
    }
    for i := 0; i < n; i ++ {
        if (cnt[i] > 0) {
            add_edge(i, root)
        }
    }
    dfs_lca(root, -1, 0)
    for k := 0; k < LOGM - 1; k ++ {
        for i := 0; i < N; i ++ {
            if (parent[k][i] < 0) {
                parent[k + 1][i] = -1
            } else {
                parent[k + 1][i] = parent[k][parent[k][i]]
            }
        }
    }
}
func lca(u, v int) int {
    if (depth[u] > depth[v]) {
        u, v = v, u
    }
    for k := 0; k < LOGM; k ++ {
        var d = depth[v] - depth[u]
        var ku uint32 = uint32(k)
        if ((d >> ku) & 1 == 1) {
            v = parent[k][v]
        }
    }
    if (u == v) {
        return u
    }
    for k := LOGM - 1; k >= 0; k -- {
        if (parent[k][u] != parent[k][v]) {
            u = parent[k][u]
            v = parent[k][v]
        }
    }
    return parent[0][u];
}
func LCA_unittest() {
    if ((lca(1, 4) != 0) ||
        (lca(5, 10) != 12) ||
        (lca(5, 7) != 3) ||
        (lca(7, 7) != 7) ||
        (lca(2, 5) != 1) ||
        (lca(10, 11) != 10)) {
        fmt.Println("LCA unittest failed!")
    }
}
//end of LCA

func is_ancestor(u, v int) bool { //whether u is v's ancestor
    l := lca(u, v)
    return l == u
}

func is_descendant(u, v int) bool { //whether u is v's descendant
    l := lca(u, v)
    return l == v
}

const query = 10000

func prep() {
    for i := 0; i < n - 1; i ++ {
        add_edge(i + 1, i)
    }
    build()
}

func benchmark1() {
    start := time.Now()
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < query; i ++ {
        a := rand.Intn(n - 2)
        b := rand.Intn(n - 2)
        is_ancestor(a, b)
    }
    end := time.Now()
    diff := (end.Sub(start)).Seconds()
    fmt.Println(diff)
}

func benchmark2() {
    start := time.Now()
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < query; i ++ {
        a := rand.Intn(n - 2)
        var _ []int = get_ancestors(a)
    }
    end := time.Now()
    diff := (end.Sub(start)).Seconds()
    fmt.Println(diff)
}

func main() {
    /*
    prep()
    benchmark1()
    benchmark2()
    */
}
