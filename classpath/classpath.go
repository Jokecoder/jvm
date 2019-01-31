package classpath

import (
	"fmt"
	"os"
	"path/filepath"
)

// 定义三种类加载器
type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

// 解析参数 -jre -cp
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

// 启动类加载器与扩展类加载器
func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

// 用户类加载器
func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newWildcardEntry(cpOption)
}

// 加载class文件的字节码 顺序 boot > ext > user
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"

	if data, Entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, Entry, nil
	}

	if data, Entry, err := self.extClasspath.readClass(className); err == nil {
		fmt.Println("ext")
		return data, Entry, nil
	}

	return self.userClasspath.readClass(className)
}

// 获取jre路径优先级为 -jre > . > JAVA_HOME
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists(". /jre") {
		return ". /jre"
	}

	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	} else {
		return "/Users/ego/go/src/JVM/jre"
	}
	panic(" Can not find jre folder!")
}

// 判断路径是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}
