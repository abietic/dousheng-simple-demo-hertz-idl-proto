#!/bin/bash


hz new -I idl -idl idl/douyin/core.proto -module dousheng/router
hz update -I idl -idl idl/douyin/extra/first.proto
hz update -I idl -idl idl/douyin/extra/second.proto

