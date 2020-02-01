package fib

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

/*
反射类型 reflect.TypeOf
反射值   reflect.ValueOf

- reflect.TypeOf 返回类型（reflect.Type)
- reflect.ValueOf 返回值  (reflect.Value)
- 可以从reflect.Value获得类型
- 通过kind来判断类型（枚举）
*/

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	case reflect.Ptr:
		fmt.Println("Ptr")
	default:
		fmt.Println("other type")
	}
}

func TestBaseType(t *testing.T) {
	var f float64 = 1
	CheckType(f)  // Float
	CheckType(&f) // Ptr
}

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f)) // int64 10
	t.Log(reflect.ValueOf(f).Type())             //int64
}

// 通过字符串的形式调用类型的某个方法
/*
按名字访问结构的成员
reflect.ValueOf(*e).FieldByName("Name")

按名字访问结构的方法
reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})
*/

type EmployeeT struct {
	EmployeeId string
	Name       string `format:"normal"` //struct tag ,类似java的注解
	Age        int
}

func (e *EmployeeT) UpdateAge(newVal int) {
	e.Age = newVal
}

func TestInvokeByName(t *testing.T) {
	e := &EmployeeT{"1", "Mike", 30}
	// 按名字获取成员
	t.Logf("Name:value(%[1]v)， Type(%[1]T)",
		reflect.ValueOf(*e).FieldByName("EmployeeId")) //  Name:value(1)， Type(reflect.Value)

	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'name' field")
	} else {
		// 通过反射获取struct tag
		t.Log("Tag: format", nameField.Tag.Get("format")) // Tag: format normal
	}

	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(99)})
	t.Log("update age:", e) // update age: &{1 Mike 99}   age更新为99
}

/*
反射万能程序

DeepEqual 比较切片和map
*/

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "two", 2: "two", 4: "three"}
	//fmt.Println(a == b) // 编译不通过
	fmt.Println(reflect.DeepEqual(a, b)) //false

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{1, 2, 4}
	t.Log("s1 == s2 ? ", reflect.DeepEqual(s1, s2)) // true
	t.Log("s1 == s3 ? ", reflect.DeepEqual(s1, s3)) // false

}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

/*
 结构体填充
*/
func fillBySettings(st interface{}, settings map[string]interface{}) error {

	// func (v Value) Elem() Value
	// Elem returns the value that the interface v contains or that the pointer v points to.
	// It panics if v's Kind is not Interface or Ptr.
	// It returns the zero Value if v is nil.

	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		return errors.New("the first param should be a pointer to the struct type.")
	}
	// Elem() 获取指针指向的值
	if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
		return errors.New("the first param should be a pointer to the struct type.")
	}

	if settings == nil {
		return errors.New("settings is nil.")
	}

	var (
		field reflect.StructField
		ok    bool
	)

	for k, v := range settings {
		// 检查类型里有没有field   （因为是指针类型，需要用Elem(),获取指针类型的）
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}

	}
	return nil
}

func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Mike", "Age": 30}
	e := Employee{}
	if err := fillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(e)
	c := new(Customer)
	if err := fillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}

/*
不安全的编程 unsafe

## 不安全编程的危险性

i := 10
// 指针可以转换成任意类型的指针，利用它来实现类型转换. 非常危险
f := *(float64)(unsafe.Pointer(&i))

*/

func TestUnsafe(t *testing.T) {
	i := 10
	f := *(*float64)(unsafe.Pointer(&i))
	t.Log(unsafe.Pointer(&i)) //  0xc000016298
	t.Log(f)                  // 5e-323
}

// atomic 进行指针的读写操作
//原子类型操作
func TestAtomic(t *testing.T) {
	var shareBufPtr unsafe.Pointer
	writeDataFn := func() {
		data := []int{}
		for i := 0; i < 100; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(&shareBufPtr, unsafe.Pointer(&data))
	}
	readDataFn := func() {
		data := atomic.LoadPointer(&shareBufPtr)
		fmt.Println(data, *(*[]int)(data))
	}
	var wg sync.WaitGroup
	writeDataFn()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				writeDataFn()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				readDataFn()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
