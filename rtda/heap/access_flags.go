package heap

const(
	ACC_PUBLIC = 0x0001 // class fields method
	ACC_PRIVATE = 0x0002 // fields method
	ACC_PROTECTED = 0x0004 // fields method
	ACC_STATIC = 0x0008 // fields method
	ACC_FINAL = 0x0010 //class fields method
	ACC_SUPER = 0x0020 // class
	ACC_SYNCHRONIZED = 0x0020 // method
	ACC_VOLATILE = 0x0040 // fields
	ACC_BRIDGE = 0x0040 // method
	ACC_TRANSIENT = 0x0080 // fields
	ACC_VARARGS = 0x0080 //method
	ACC_NATIVE = 0x0100 //method
	ACC_INTERFACE = 0x0200 // class
	ACC_ABSTRACT = 0x0400 //class method
	ACC_STRICT = 0x0800 // method
	ACC_SYNTHETIC = 0x1000 // class fields method
	ACC_ANNOTATION = 0x2000 // calss
	ACC_ENUM = 0x4000 // class fields
)