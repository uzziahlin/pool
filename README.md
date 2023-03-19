# 连接池

基于go语言的连接池实现，使用了泛型，可以更好的支持各种连接。

使用示例如下：
```
connPool := NewConnPool[net.Conn](30, 30 * time.Second, func() net.Conn {
    conn, err := net.Dial("tcp", "your server address")
    if err != nil {
        panic(err)
    }
    return conn
})

conn, err := connPool.Get(context.TODO())

defer func() {
    err = connPool.Put(context.TODO(), conn)
    
    if err != nil {
        // todo handle error
    }
}()

if err != nil {
    // todo handle error
}

// todo use conn

