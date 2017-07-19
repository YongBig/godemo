//当前程序的包名，注意必须在非注释代码的首行申明
package main

//导入其它包,导入其他包，必须要引用，不然会报错. stu为fmt的别名可自定义，
// 如果为 . 则为省略在调用包中的方法时可直接写方法，变量使用
import "fmt"
//import stu "fmt"
//import . "fmt"
//import "os"
//import "time"
//import "strings"

//多个导入第三方包简写
/*import (
	"fmt"
	"os"
	"time"
	"strings"
)*/

//常量定义
const pi  = "我是很帅，对不对？对！"

//全局变量申明
var name = "呸！我才是最帅的！"

//一般类型申明
type newtype int

//结构声明
type du struct {

}

//接口申明
type yong interface {

}

//由main函数作为程序入口点启动
func main()  {
	fmt.Println("我说：" + pi + "刘德华说：" + name)
}

//此外注意函数 变量命名要注意的点
/*名称区分大小写，首字母为大写则其他包中可以调用,为小写则其他包中不能使用，只作用当前的包内使用*/

