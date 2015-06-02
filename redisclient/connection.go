package redisclient

import (
    "github.com/fzzy/radix/redis"
)

func Connection(host string) (*redis.Client, error) {
    client, err := redis.Dial("tcp", host)
    if err != nil {
        return client, err
    }
    return client, nil
}
