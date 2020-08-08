package heap



func (self *Class) IsAssignableFrom(other *Class) bool {
	s, t := other, self
	if s == t{
		return true
	}
	if !s.IsArray(){
		if !s.IsInterface(){
			// s is class
			if !t.IsInterface(){
				// t is not interface
				return s.IsSubClassOf(t)
			}else {
				// t is interface
				return s.IsImplements(t)
			}
		} else {
			// s is interface
			if !t.IsInterface() {
				// t is not interface
				return t.isJlObject()
			}else {
				// t is interface
				return t.isSuperInterfaceOf(s)
			}
		}
	} else {
		// s is arrray
		if !t.IsArray() {
			if ! t.IsInterface() {
				// t is class 
				return t.isJlObject()
			}else {
				// t is interface
				return t.isJlCloneable() || t.isJioSerializable()
			}
		}else {
			// t is array
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.IsAssignableFrom(sc)
		}
	}
	return false
}

func (self *Class) IsSubClassOf(other *Class) bool{
	for c := self.superClass; c != nil ; c = c.superClass{
		if c == other {
			return true
		}
	}
	return false
}

func (self *Class) IsImplements(iface *Class) bool{
	for c:= self; c != nil; c = c.superClass{
		for _, i := range c.interfaces {
			if i == iface || i.IsSubInterfaceOf(iface){
				return true
			}
		}
	}
	return false
}

func (self *Class) IsSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == iface || superInterface.IsSubInterfaceOf(iface){
			return true
		}
	}
	return false
}

func (self *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(self)
}

// iface extends self
func (self *Class) isSuperInterfaceOf(iface *Class) bool {
	return iface.IsSubInterfaceOf(self)
}