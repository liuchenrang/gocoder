#!/bin/env bash
go build -o gen main.go

sh -x ./make.sh test /Users/chen/IdeaProjects/lanlan-micro/userservice/user-service-srv coder.yml

# 废弃 sh -x ./make.sh sqb_user_red_config /Users/chen/IdeaProjects/lanlan-micro/userservice/user-service-srv config/user-coder.yml
