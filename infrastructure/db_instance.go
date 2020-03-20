package infrastructure

package (
	_ "github.com/go-sql-driver/mysql"
)

func NewInstance() (*sql.DB, err) {
	return sql.Open("mysql", "root@tcp(127.0.0.1:3306)/cocomeshi")
} 