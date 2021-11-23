package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent --idtype=int64 --template glob=./tmpl/*.tmpl --target=../model generate ./schema
