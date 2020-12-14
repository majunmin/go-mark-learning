# Reflect.ValueOf 与 Reflect.TypeOf

## 反射类型 Type 和反射种类 Kind

![反射类型](https://raw.githubusercontent.com/majunmin/image/master/img%E5%8F%8D%E5%B0%84%E7%B1%BB%E5%9E%8B.jpg)

反射种类
```go
// reflect.go
type Kind uint

const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	String
	Struct
	UnsafePointer
)
```

```go
// reflect.go
	// IsVariadic reports whether a function type's final input parameter
	// is a "..." parameter. If so, t.In(t.NumIn() - 1) returns the parameter's
	// implicit actual type []T.
	//
	// For concreteness, if t represents func(x int, y ... float64), then
	//
	//	t.NumIn() == 2
	//	t.In(0) is the reflect.Type for "int"
	//	t.In(1) is the reflect.Type for "[]float64"
	//	t.IsVariadic() == true
	//
	// IsVariadic panics if the type's Kind is not Func.
	IsVariadic() bool

	// Elem returns a type's element type.
	// It panics if the type's Kind is not Array, Chan, Map, Ptr, or Slice.
	Elem() Type
```

> 在使用反射时，需要首先理解类型（Type）和种类（Kind）的区别。
> 编程中，使用最多的是`类型`，但在反射中，当需要区分一个大品种的类型时，就会用到`种类(Kind)`.
> 例如，需要统一判断类型中的指针时，使用种类(Kind)信息就较为方便。

Kind

Go 程序中的类型(Type)指的是系统原生数据类型，如 int、string、bool、float32 等类型，以及使用 type 关键字定义的类型，这些类型的名称就是其类型本身的名称。
 例如使用 type A struct{} 定义结构体时，A 就是 struct{} 的类型。
 
`Map`、`Slice`、`Chan` 属于引用类型，使用起来类似于`指针`，但是在种类常量定义中仍然属于独立的种类,不属于 `Ptr`
`type A struct{}` 定义的结构体属于 Struct 种类，`*A 属于 Ptr`.


 

## Reflect.TypeOf




----------------
[Go 系列教程 —— 34. 反射](https://studygolang.com/articles/13178)