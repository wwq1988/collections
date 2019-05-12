
package commons

import (
	"sort"
	"math/rand"
)

type Int64Stream struct{
	value	[]int64
	defaultReturn int64
}

func StreamOfInt64(value []int64) *Int64Stream {
	return &Int64Stream{value:value,defaultReturn:0}
}

func(c *Int64Stream) OrElase(defaultReturn int64)  *Int64Stream {
	c.defaultReturn = defaultReturn
	return c
}


func(c *Int64Stream) Concate(given []int64)  *Int64Stream {
	value := make([]int64, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Int64Stream) Drop(n int)  *Int64Stream {
	l := len(c.value) - n
	if l < 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Int64Stream) Filter(fn func(int, int64)bool)  *Int64Stream {
	value := make([]int64, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Int64Stream) First() int64 {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[0]
}

func(c *Int64Stream) Last() int64 {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[len(c.value)-1]
}

func(c *Int64Stream) Map(fn func(int, int64)) *Int64Stream {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Int64Stream) Reduce(fn func(int64, int64, int) int64,initial int64) int64   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Int64Stream) Reverse()  *Int64Stream {
	value := make([]int64, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Int64Stream) Unique()  *Int64Stream{
	value := make([]int64, 0, len(c.value))
	seen:=make(map[int64]struct{})
	for _, each := range c.value {
		if _,exist:=seen[each];exist{
			continue
		}		
		seen[each]=struct{}{}
		value=append(value,each)			
	}
	c.value = value
	return c
}

func(c *Int64Stream) Append(given int64) *Int64Stream {
	c.value=append(c.value,given)
	return c
}

func(c *Int64Stream) Len() int {
	return len(c.value)
}

func(c *Int64Stream) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *Int64Stream) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *Int64Stream)  Sort()  *Int64Stream {
	sort.Slice(c.value, func(i,j int)bool{
		return c.value[i] < c.value[j]
	})
	return c 
}

func(c *Int64Stream) All(fn func(int, int64)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *Int64Stream) Any(fn func(int, int64)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *Int64Stream) Paginate(size int)  [][]int64 {
	var pages  [][]int64
	prev := -1
	for i := range c.value {
		if (i-prev) < size-1 && i != (len(c.value)-1) {
			continue
		}
		pages=append(pages,c.value[prev+1:i+1])
		prev=i
	}
	return pages
}

func(c *Int64Stream) Pop() int64{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *Int64Stream) Prepend(given int64) *Int64Stream {
	c.value = append([]int64{given},c.value...)
	return c
}

func(c *Int64Stream) Max() int64{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var max int64 = c.value[0]
	for _,each := range c.value {
		if max < each {
			max = each
		}
	}
	return max
}


func(c *Int64Stream) Min() int64{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var min int64 = c.value[0]
	for _,each := range c.value {
		if each  < min {
			min = each
		}
	}
	return min
}

func(c *Int64Stream) Random() int64{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *Int64Stream) Shuffle() *Int64Stream {
	if len(c.value) <= 0 {
		return c
	}
	indexes := make([]int, len(c.value))
	for i := range c.value {
		indexes[i] = i
	}
	
	rand.Shuffle(len(c.value), func(i, j int) {
		c.value[i], c.value[j] = 	c.value[j], c.value[i] 
	})
	
	return c
}

func(c *Int64Stream) Collect() []int64{
	return c.value
}


type Int64PStream struct{
	value	[]*int64
	defaultReturn *int64
}

func PStreamOfInt64(value []*int64) *Int64PStream {
	return &Int64PStream{value:value,defaultReturn:nil}
}

func(c *Int64PStream) OrElse(defaultReturn *int64)  *Int64PStream {
	c.defaultReturn = defaultReturn
	return c
}

func(c *Int64PStream) Concate(given []*int64)  *Int64PStream {
	value := make([]*int64, len(c.value)+len(given))
	copy(value,c.value)
	copy(value[len(c.value):],given)
	c.value = value
	return c
}

func(c *Int64PStream) Drop(n int)  *Int64PStream {
	l := len(c.value) - n
	if l < 0 {
		l = 0
	}
	c.value = c.value[len(c.value)-l:]
	return c
}

func(c *Int64PStream) Filter(fn func(int, *int64)bool)  *Int64PStream {
	value := make([]*int64, 0, len(c.value))
	for i, each := range c.value {
		if fn(i,each){
			value = append(value,each)
		}
	}
	c.value = value
	return c
}

func(c *Int64PStream) First() *int64 {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[0]
}

func(c *Int64PStream) Last() *int64 {
	if len(c.value) <= 0 {
		return c.defaultReturn
	} 
	return c.value[len(c.value)-1]
}

func(c *Int64PStream) Map(fn func(int, *int64)) *Int64PStream {
	for i, each := range c.value {
		fn(i,each)
	}
	return c
}

func(c *Int64PStream) Reduce(fn func(*int64, *int64, int) *int64,initial *int64) *int64   {
	final := initial
	for i, each := range c.value {
		final = fn(final,each,i)
	}
	return final
}

func(c *Int64PStream) Reverse()  *Int64PStream {
	value := make([]*int64, len(c.value))
	for i, each := range c.value {
		value[len(c.value)-1-i] = each
	}
	c.value = value
	return c
}

func(c *Int64PStream) Unique()  *Int64PStream{
	value := make([]*int64, 0, len(c.value))
	seen:=make(map[*int64]struct{})
	for _, each := range c.value {
		if _,exist:=seen[each];exist{
			continue
		}		
		seen[each]=struct{}{}
		value=append(value,each)			
	}
	c.value = value
	return c
}

func(c *Int64PStream) Append(given *int64) *Int64PStream {
	c.value=append(c.value,given)
	return c
}

func(c *Int64PStream) Len() int {
	return len(c.value)
}

func(c *Int64PStream) IsEmpty() bool {
	return len(c.value) == 0
}

func(c *Int64PStream) IsNotEmpty() bool {
	return len(c.value) != 0
}

func(c *Int64PStream)  Sort(less func(*int64,*int64) bool )  *Int64PStream {
	sort.Slice(c.value, func(i,j int)bool{
		return less(c.value[i],c.value[j])
	})
	return c 
}

func(c *Int64PStream) All(fn func(int, *int64)bool)  bool {
	for i, each := range c.value {
		if !fn(i,each){
			return false
		}
	}
	return true
}

func(c *Int64PStream) Any(fn func(int, *int64)bool)  bool {
	for i, each := range c.value {
		if fn(i,each){
			return true
		}
	}
	return false
}

func(c *Int64PStream) Paginate(size int)  [][]*int64 {
	var pages  [][]*int64
	prev := -1
	for i := range c.value {
		if (i-prev) < size-1 && i != (len(c.value)-1) {
			continue
		}
		pages=append(pages,c.value[prev+1:i+1])
		prev=i
	}
	return pages
}

func(c *Int64PStream) Pop() *int64{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	lastIdx := len(c.value)-1
	val:=c.value[lastIdx]
	c.value=c.value[:lastIdx]
	return val
}

func(c *Int64PStream) Prepend(given *int64) *Int64PStream {
	c.value = append([]*int64{given},c.value...)
	return c
}

func(c *Int64PStream) Max() *int64{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var max *int64 = c.value[0]
	for _,each := range c.value {
		if max == nil{
			max = each
			continue
		}
		if each != nil && *max <= *each {
			max = each
		}
	}
	return max
}


func(c *Int64PStream) Min() *int64{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	var min *int64 = c.value[0]
	for _,each := range c.value {
		if min == nil{
			min = each
			continue
		}
		if  each != nil && *each  <= *min {
			min = each
		}
	}
	return min
}

func(c *Int64PStream) Random() *int64{
	if len(c.value) <= 0 {
		return c.defaultReturn
	}
	n := rand.Intn(len(c.value))
	return c.value[n]
}

func(c *Int64PStream) Shuffle() *Int64PStream {
	if len(c.value) <= 0 {
		return c
	}
	indexes := make([]int, len(c.value))
	for i := range c.value {
		indexes[i] = i
	}
	
	rand.Shuffle(len(c.value), func(i, j int) {
		c.value[i], c.value[j] = 	c.value[j], c.value[i] 
	})
	
	return c
}

func(c *Int64PStream) Collect() []*int64{
	return c.value
}
