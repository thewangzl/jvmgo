package heap

import "jvmgo/ch08/classfile"
import "strings"

type Class struct {
	accessFlags uint16
	name 		string	//thisClassName
	superClassName 	string
	interfaceNames 	[]string
	constantPool 	*ConstantPool
	fields 			[]*Field
	methods 		[]*Method
	loader 			*ClassLoader
	superClass 		*Class
	interfaces 		[]*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars 		  Slots
	initStarted      bool
}

func newClass(cf *classfile.ClassFile ) *Class{
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	
	//
	return class
}


func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags & ACC_PUBLIC
}
func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}
func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}
func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}
func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

//getters
func (self *Class) Name() string{
	return self.name
}

func (self *Class) ConstantPool() *ConstantPool{
	return self.constantPool
}

func (self *Class) Fields() []*Field {
	return self.fields
}

func (self *Class) Methods() []*Method {
	return self.methods
}

func (self *Class) Loader() *ClassLoader {
	return self.loader
}

func (self *Class) SuperClass() *Class {
	return self.superClass
}

func (self *Class) StaticVars() Slots {
	return self.staticVars
}

func (self *Class) InitStarted() bool {
	return self.initStarted
}

func (self *Class) StartInit() {
	self.initStarted = true
}

func (self *Class) GetPackageName() string{
	if i:= strings.LastIndex(self.name, "/"); i >= 0{
		return self.name[: i]
	}
	return ""
}

func (self *Class) isAccessibleTo(other *Class) bool{
	return self.IsPublic() || self.GetPackageName() == other.GetPackageName()
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}

func (self *Class) GetMainMethod() *Method{
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) GetClinitMethod() *Method{
	return self.getStaticMethod("<clinit>", "()V")
}

func (self *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() && method.name == name && method.descriptor == descriptor{
			return method
		}
	}
	return nil
}

func (self *Class) isJlObject() bool {
	return self.name == "java/lang/Object"
}

func (self *Class) isJlCloneable() bool {
	return self.name == "java/lang/Cloneable"
}

func (self *Class) isJioSerializable() bool {
	return self.name == "java/io/Serializable"
}

func (self *Class) ArrayClass() *Class{
	arrayClassName := getArrayClassName(self.name)
	return self.loader.LoadClass(arrayClassName)
}

func (self *Class) getField(name, descriptor string, isStatic bool) *Field {
	for c := self; c != nil; c = c.superClass {
		for _, field := range c.fields {
			if field.IsStatic() == isStatic && field.name == name && field.descriptor == descriptor {
				return field
			}
		}
	}
	return nil
}