package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/shanepeckham/hackfulfillorder/controllers:OrderController"] = append(beego.GlobalControllerRouter["github.com/shanepeckham/hackfulfillorder/controllers:OrderController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

}