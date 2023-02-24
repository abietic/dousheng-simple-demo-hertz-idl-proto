init:
	# hz new -I=idl --idl=idl/hello/hello.proto --protoc-plugins=openapi::./docs --mod=middleware/hertz
	hz new -I idl -idl idl/douyin/core.proto -module dousheng/router
	hz update -I idl -idl idl/douyin/extra/first.proto
	hz update -I idl -idl idl/douyin/extra/second.proto
update:
	# hz update -I=idl --idl=idl/hello/hello.proto --protoc-plugins=openapi::./docs --mod=middleware/hertz
	hz new -I idl -idl idl/douyin/core.proto -module dousheng/router
	hz update -I idl -idl idl/douyin/extra/first.proto
	hz update -I idl -idl idl/douyin/extra/second.proto
