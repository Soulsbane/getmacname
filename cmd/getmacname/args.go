package main

type ProgramArgs struct {
	Address string `arg:"positional, required"`
}

func (args ProgramArgs) Description() string {
	return "Takes a MAC address and finds its manufacturers name."
}
