## Conner

> Contextual, structured errors for Go 1.13+

This is not an a logging library - it's a supplemental error library.

If an error happens somewhere in the call stack, then each caller can append some metadata to the error and pass it up so that the final root caller will have all the information about each of the contexts that the error happened in and can use a logger like [github.com/rs/zerolog](https://github.com/rs/zerolog) or [github.com/sirupsen/logrus](https://github.com/sirupsen/logrus).


## Background

Go 1.13 finally brought us [errors.As](https://golang.org/pkg/errors/#As) and [fmt.Errorf("Foo: %w", err)](https://golang.org/pkg/fmt/#Errorf) for [wrapping errors](https://medium.com/@felipedutratine/golang-how-to-handle-errors-in-v1-13-fda7f035d027) by providing more (string) context.

However, production applications often have to send structured logs (JSON) to collectors so that large volumes of information can be processed while still providing clear traces of individual clients or spans.

Now that we know how to add to error message strings - how do we also add context values to these error objects like `request-id` or `user-id`?

In other words, I want the application to both be able to provide detailed text logs as well as detailed structured logs. So the "EOF" error encountered in `FooProcessor() -> Request() -> JSON.Decode() -> EOF`

Should be accessible as either a terminal output of

    FooProcessor: 114153: Request: JSON Decode: EOF

A plain text log of

    "2019/10/07 11:55:55 FooProcessor: 114153: Request: JSON Decode: EOF"

Or a structured log of:

     {"foo_id":"114153","level":"info","msg":"JSON Decode: EOF","time":"2019-10-07T11:55:55-00:00"}

This is the problem Conner solves: structured errors with additional parameters at each wrapped level to help when logging errors.
