package heap

type Object struct {
	class *Class
	data interface{}
}

func newObject(class *Class) *Object {
	return &Object {
		class : class,
		data : newSlots(class.instanceSlotCount),
	}
}

func (self *Object) Class() *Class{
	return self.class
}

func (self *Object) Fields() Slots{
	return self.data.(Slots)
}

func (self *Object) IsInstanceOf(class *Class) bool{
	return class.IsAssignableFrom(self.class)
}

func (self *Object) SetRefVar(name, descriptor string, ref *Object) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetRef(field.slotId, ref)
}

func (self *Object) GetRefVar(name, descriptor string) *Object {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetRef(field.slotId)
}