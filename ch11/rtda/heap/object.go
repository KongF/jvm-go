package heap

type Object struct {
	class *Class
	data  interface{}
	extra interface{}
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}
func (self *Object) Class() *Class {
	return self.class
}
func (self *Object) Fields() Slots {
	return self.data.(Slots)
}
func (self *Object) Extra() interface{} {
	return self.extra
}
func (self *Object) SetExtra(extra interface{}) {
	self.extra = extra
}
func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}

func (self *Object) SetRefVar(name string, descriptor string, ref *Object) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetRef(field.slotId, ref)
}

func (self *Object) GetRefVar(name string, descriptor string) *Object {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetRef(field.slotId)
}

func (self *Object) Clone() *Object {
	return &Object{
		class: self.class,
		data:  self.cloneData(),
	}
}

func (self *Object) cloneData() interface{} {
	switch self.data.(type) {
	case []int8:
		elements := self.data.([]int8)
		elements2 := make([]int8, len(elements))
		copy(elements2, elements)
		return elements2
	case []int16:
		elements := self.data.([]int16)
		elements2 := make([]int16, len(elements))
		copy(elements2, elements)
		return elements2
	case []uint16:
		elements := self.data.([]uint16)
		elements2 := make([]uint16, len(elements))
		copy(elements2, elements)
		return elements2
	case []int32:
		elements := self.data.([]int32)
		elements2 := make([]int32, len(elements))
		copy(elements2, elements)
		return elements2
	case []int64:
		elements := self.data.([]int64)
		elements2 := make([]int64, len(elements))
		copy(elements2, elements)
		return elements2
	case []float32:
		elements := self.data.([]float32)
		elements2 := make([]float32, len(elements))
		copy(elements2, elements)
		return elements2
	case []float64:
		elements := self.data.([]float64)
		elements2 := make([]float64, len(elements))
		copy(elements2, elements)
		return elements2
	case []*Object:
		elements := self.data.([]*Object)
		elements2 := make([]*Object, len(elements))
		copy(elements2, elements)
		return elements2
	default: // []Slot
		slots := self.data.(Slots)
		slots2 := newSlots(uint(len(slots)))
		copy(slots2, slots)
		return slots2
	}
}
