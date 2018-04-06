package main

import (
    "fmt"
    "reflect"
    "time"
    "strconv"
    "math/rand"
)

type Category struct {
    id int
    par *Category
    name string
}

var categories = make([]Category, 0, 1)

func get_categories() []Category {
    return categories
}

func (c *Category) get_name() string {
    return c.name
}

func (c *Category) is_ancestor(a *Category) bool {
    var cur *Category = c
    for {
        if ((cur.id == a.id) || (cur == cur.par)) {
            break
        }
        cur = cur.par
    }
    return cur.id == a.id
}

func (c *Category) is_descendant(a *Category) bool {
    return (*a).is_ancestor(c)
}

func get_root(c Category) Category {
    var cur Category = c
    for {
        if (cur == *cur.par) {
            break
        }
        cur = *cur.par
    }
    return cur
}

func get_ancestors(c Category) []Category {
    var all = get_categories()
    var ancestors = make([]Category, 0, 1)
    for _, cat := range all {
        if (c.is_ancestor(&cat)) {
            ancestors = append(ancestors, cat)
        }
    }
    return ancestors
}

func get_descendants(c Category) []Category {
    var all = get_categories()
    var descendants = make([]Category, 0, 1)
    for _, cat := range all {
        if (c.is_descendant(&cat)) {
            descendants = append(descendants, cat)
        }
    }
    return descendants
}

func unittest1() {
    var c Category
    c.id, c.par, c.name = 3, &c, "C"
    var b = Category{id : 2, par : &c, name : "B"}
    var a = Category{id : 1, par : &b, name : "A"}
    var d = Category{id : 4, par : &a, name : "D"}
    var e = Category{id : 5, par : &a, name : "E"}
    var f = Category{id : 6, par : &b, name : "F"}
    var g = Category{id : 7, par : &b, name : "G"}
    var h Category
    h.id, h.par, h.name = 8, &h, "H"
    var i Category
    i.id, i.par, i.name = 9, &i, "I"
    var j = Category{id : 10, par : &i, name : "J"}
    var k = Category{id : 11, par : &j, name : "K"}
    if ((c.is_ancestor(&c) == false) ||
        (a.is_ancestor(&c) == false) ||
        (d.is_ancestor(&b) == false) ||
        (e.is_ancestor(&e) == false) ||
        (f.is_ancestor(&c) == false) ||
        (b.is_ancestor(&a) == true) ||
        (d.is_ancestor(&e) == true) ||
        (f.is_ancestor(&g) == true) ||
        (g.is_ancestor(&e) == true) ||
        (c.is_ancestor(&g) == true)) {
        fmt.Println("is_ancestor test failed!")
    }
    if (h.is_ancestor(&h) == false) {
        fmt.Println("is_ancestor test failed!")
    }
    if ((j.is_ancestor(&i) == false) ||
        (k.is_ancestor(&k) == false) ||
        (i.is_ancestor(&j) == true)) {
        fmt.Println("is_ancestor test failed!")
    }
    fmt.Println("is_ancestor OK")
}

func unittest2() {
    var c Category
    c.id, c.par, c.name = 3, &c, "C"
    categories = append(categories, c)
    var b = Category{id : 2, par : &c, name : "B"}
    categories = append(categories, b)
    var a = Category{id : 1, par : &b, name : "A"}
    categories = append(categories, a)
    var d = Category{id : 4, par : &a, name : "D"}
    categories = append(categories, d)
    var e = Category{id : 5, par : &a, name : "E"}
    categories = append(categories, e)
    var f = Category{id : 6, par : &b, name : "F"}
    categories = append(categories, f)
    var g = Category{id : 7, par : &b, name : "G"}
    categories = append(categories, g)
    var h Category
    h.id, h.par, h.name = 8, &h, "H"
    categories = append(categories, h)
    var i Category
    i.id, i.par, i.name = 9, &i, "I"
    categories = append(categories, i)
    var j = Category{id : 10, par : &i, name : "J"}
    categories = append(categories, j)
    var k = Category{id : 11, par : &j, name : "K"}
    categories = append(categories, k)
    var aa = []Category {a, d, e}
    if (reflect.DeepEqual(get_descendants(a), aa) == false) {
        fmt.Println("get_descendants test failed!")
    }
    var bb = []Category {b, a, d, e, f, g}
    if (reflect.DeepEqual(get_descendants(b), bb) == false) {
        fmt.Println("get_descendants test failed!")
    }
    var cc = []Category {c, b, a, d, e, f, g}
    if (reflect.DeepEqual(get_descendants(c), cc) == false) {
        fmt.Println("get_descendants test failed!")
    }
    var dd = []Category {d}
    if (reflect.DeepEqual(get_descendants(d), dd) == false) {
        fmt.Println("get_descendants test failed!")
    }
    fmt.Println("get_descendants OK")
}

const n = 1000
const query = 10000

func prep() {
    var root Category
    root.id, root.par, root.name = 0, &root, "0"
    var p *Category
    p = &root
    for i := 0; i < n; i ++ {
        var c Category
        c.id = i
        c.name = strconv.Itoa(i)
        c.par = p
        categories = append(categories, c)
        p = &c
    }
}

func benchmark1() {
    start := time.Now()
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < query; i ++ {
        a := rand.Intn(n - 2)
        b := rand.Intn(n - 2)
        categories[a].is_ancestor(&categories[b])
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
        var _ []Category = get_descendants(categories[a])
    }
    end := time.Now()
    diff := (end.Sub(start)).Seconds()
    fmt.Println(diff)
}

func main() {
    /*
    prep()
    unittest1()
    unittest2()
    benchmark1()
    benchmark2()
    */
}
