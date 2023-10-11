package connection

func (c *Connection) SetDatabase(name string) {
    c.Database = c.Client.Database(name, nil)
}
