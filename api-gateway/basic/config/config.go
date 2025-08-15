package config

import (
	__ "api-gateway/basic/device-proto"
	pex "api-gateway/basic/pex-proto"
	zhc "api-gateway/basic/user-proto"
)

var DeviceSrv __.DeviceSrvClient
var PexSrv pex.EMQXSrvClient
var UserSrv zhc.UserClient
