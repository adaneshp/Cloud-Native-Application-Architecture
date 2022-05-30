# Lab 2: Go Interfaces

- This code was developed in January, 2022.

## Instructions

Interfaces are used to separate behavior of a type from its concrete implementation. In C++ abstract classes serve this purpose. In Go, interfaces may be implemented by an arbitrary number of types without explicit inheritance (like in C++). And conversely, a type may implement an arbitrary number of interfaces. 
Every type implements the empty interface (interface{})

Example use of interface (From Tour of Go)

MyFloat and Vertex both implicitly implement the Abser interface since they both have the Abs() method
```
type Abser interface {
	Abs() float64
}

type MyFloat float64
// Note: Adding methods to predefined types such as float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}


func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// Error: In the following line, vertex (the value type) doesn't implement Abser because the Abs method is defined only on *Vertex (the pointer type).
	a = v
	fmt.Println(a.Abs())
}
```

Let’s take a look at the use of interfaces to define abstract behavior of a cache. Here Cacher is an interface that defines two methods Get() and Put(). These take in an empty interface as an argument allowing accepting arguments of any type.

lruCache is a concrete implementation of a cache. As long as lruCache implements the Get() and Put() methods, it satisfies the Cacher interface, and can be used anywhere where Cacher is used. Note that you could have another cache concrete type (perhaps, with different replacement policy) with the same methods that also satisfies the Cacher interface. The concrete implementation is thus separated from the behavior allowing the implementation to be changed without affecting the application that uses it (decoupling of implementation and use). 
```
package cache

import "errors"

type Cacher interface {
    Get(interface{}) (interface{}, error)
    Put(interface{}, interface{}) (error)
}

// Concrete LRU cache
type lruCache struct {
    size int  // Size of cache
    remaining int // Remaining capacity
    cache map[string]string // Actual storage of data
    queue []string // For keeping track of lru
}
// Constructor 
func NewCacher(size int) Cacher {
    return &lruCache{size:size, remaining:size, cache:make(map[string]string), queue:make([]string, size)}
}
```

## Extra
Implementing the Get() and Put() methods for lruCache.
