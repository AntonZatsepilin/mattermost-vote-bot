box.cfg{
    listen = '0.0.0.0:3301',
    wal_mode = 'none',
    memtx_memory = 256 * 1024 * 1024
}

box.schema.space.create('polls', {
    if_not_exists = true,
    format = {
        {name = 'id', type = 'string'},
        {name = 'creator', type = 'string'},
        {name = 'question', type = 'string'},
        {name = 'options', type: 'map'},
        {name = 'status', type = 'string'},
        {name = 'created_at', type = 'unsigned'}
    }
})

box.space.polls:create_index('primary', {
    parts = {'id'},
    if_not_exists = true
})

box.schema.space.create('votes', {
    if_not_exists = true,
    format = {
        {name = 'poll_id', type = 'string'},
        {name = 'user_id', type = 'string'},
        {name = 'option', type = 'string'},
        {name = 'timestamp', type = 'unsigned'}
    }
})

box.space.votes:create_index('primary', {
    parts = {'poll_id', 'user_id'},
    if_not_exists = true
})