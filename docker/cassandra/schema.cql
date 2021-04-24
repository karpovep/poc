create keyspace if not exists cloud
    with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
