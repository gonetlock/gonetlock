GoNetLock
===============

GoNetLock is a distributed lock server written in Go.

# Features
* Protobuf based network protocol
* Both waiting and breaking lock requests
* Read locks
* Multiple clients per connection
* Lock wait and lock hold timeouts
* Automatic lock release when a client's last connection drops
* Option to ignore connection drop

# Covered use cases
* Wait (until an optional timeout) in a queue to get the lock
* Get the lock if available or give up if not
* Wait for the lock and then ignore when the connection drops. This allows us to get the lock with a shell command. A held lock can be refreshed or released before it times out.
* Each lock is assigned to a client and each connection can have multiple clients. This allows us to separate locks for each request that our app is handling at the same time.
* Same client can use more than one connection. This allows us to create a new connection to the lock server before closing the old one without losing any locks.

# Alternatives
* If your app is running on a single server, you can use filesystem locks instead.
* Symfony's lock component supports a bunch of stores and has a good overview of what they can or can't do: https://symfony.com/doc/current/components/lock.html#available-stores

# Usage examples

## Avoiding errors caused by race conditions
When a new user is registering, we can make sure that if the user managed to send multiple requests at the same time, our app won't break by trying to insert a duplicate record into our DB.
```
lock = getLock('registration-user@email.com')

// make sure that the user doesn't already exist

// register the user

lock.release()
```

## Limiting the number of parallel processes
If we're running processes with cron, we can make sure that only one or some other fixed number of processes are running at the same time
```
maxProcesses = 3

lock = null
for (i = 0; i < maxProcesses; i++) {
  lock = getBreakingLock('cronjob' ~ i)

  if (lock) {
    break
  }
}

if (!lock) {
  quit('Could not get a lock, quitting.')
}
```

## Caching data for heavy requests
To avoid running heavy database queries multiple times, we can make sure that we run them only once and then cache them. This is especially helpful when the cache is warming up under high load.
```
data = getFromCache('big-data')

if (!data) {
  lock = getLock('cache-big-data')

  // check again in case another request cached the data before we got the lock
  data = getFromCache('big-data')

  if (!data) {
    data = getBigDataFromDb()
    storeToCache('big-data', data)
  }

  lock.release()
}
```
