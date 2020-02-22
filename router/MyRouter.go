package router

func MyRouter(base string) {
	r := Router.Group("/" + base)

}
