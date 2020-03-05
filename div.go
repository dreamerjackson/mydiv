package mydiv
import "github.com/pkg/errors"

func Div(a int,b int) (int,error){
	if b==0{
		return 0,errors.Errorf("v1.0.1 b can't = 0")
	}
	return a/b,nil
}