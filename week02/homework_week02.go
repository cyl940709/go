//需要往上抛弃。
//因为事务的边界在service层不在DAO层。
//假如一个业务逻辑调用了两个 DAO 的方法（A、B）组成一个事务体，若 A 正常执行了， 而 B 执行时却发生了异常，这时 A 的执行就需要回滚，若将 B 中的error处理掉了，那根本就不知道 B 是正常执行而是异常执行了，也无法处理事务了。
import (
"database/sql"
"github.com/pkg/errors"
)

type Customer struct {
CustomerId string
Name       string
}


var Db *sql.DB

func init() {
var err error
Db, err = sql.Open("mysql", "root:ttalbe@tcp(127.0.0.1:3306)/test?charset=utf8")
if err != nil {
panic(err)
}
}

func QueryCustomerById(id string) (Customer, error) {
var customer Customer
row := Db.QueryRow("select id ,name from customer where id = ?" ,id )
err := row.Scan(&customer.CustomerId,&customer.Name)
if err != nil{
return customer,errors.Wrap(err,"dao#QueryCustomerById err")
}
return customer,nil
}

func main(){
	defer dao.Db.Close()
	customer ,err :=dao.QueryCustomerById("123456")
	if err != nil{
		fmt.Printf("query customer err : %+v",err)
		return
	}
	fmt.Println("query customer : ",customer)
}

