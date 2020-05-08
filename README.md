This example shows additional improvements by replacing  [gorilla/websocket]
(https://github.com/gorilla/websocket/) library with [gobwas/ws](https://github.com/gobwas/ws)


This allows greater performance and lower memory footprint, 
mostly due to the performant design in gobwas/ws library 
that allows to reuse the allocated buffers between connections

**How to run:**

set path to keys in environment variables, e.g.:


then build:
_make_
