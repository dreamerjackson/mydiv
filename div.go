package mydiv
import "github.com/pkg/errors"
func Div(a int,b int) (int,error){
	if b==0{
		return 0,errors.Errorf("b can't = 0")
	}
	return a/b,nil
}