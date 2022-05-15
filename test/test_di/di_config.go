package main

import (
	"GoParser2/test/test_di/biz/impl"
	"GoParser2/test/test_di/biz/impl/gen"
	impl2 "GoParser2/test/test_di/dao/impl"
	"GoParser2/test/test_di/framework"
	"GoParser2/test/test_di/service"
)

var Container = framework.NewContainer()

func init() {
	Container.RegisterBean("xxxService", &service.XxxService{})
	Container.RegisterBean("xxxBizTx", &gen.XxxBizImplTx{})
	Container.RegisterBean("xxxBiz", &impl.XxxBizImpl{})
	Container.RegisterBean("xxxDao", &impl2.XxxDaoImpl{})
}
