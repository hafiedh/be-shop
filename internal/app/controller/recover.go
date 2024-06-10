package controller

import "log/slog"

func Recover() {
	defer func() {
		if r := recover(); r != nil {
			slog.Error("CreateGroup - something went wrong", r)
		}
	}()
}
