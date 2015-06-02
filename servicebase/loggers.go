package servicebase

import (
    "log"
    "os"
)

type Loggers struct {
    I *log.Logger
    E *log.Logger
}

func InitLoggers(title string) *Loggers {

    info := log.New(os.Stdout, title + " info ", log.LstdFlags)
    err := log.New(os.Stdout, title + " error ", log.LstdFlags)

    return &Loggers{info, err}
}
