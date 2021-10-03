https://redis.com/solutions/use-cases/messaging/

https://github.com/hibiken/asynq

https://www.cloudcomputingpatterns.org/at_least_once_delivery/

Communicating through a stream is different than using a message queue. As mentioned previously, message queues are “push,” while streams are “pull.” In practice, this means that each service writes to its own stream and other services will optionally observe (i.e. “pull” from) it. This makes one-to-many communication much more efficient than with message queues.

Space efficiency is a welcome property for all communication channels that persist messages. For event streams, though, it’s fundamental, as they are often used for long-term information storage. (We mentioned above that immutable logs are fast at appending new entries and at seeking through history.)

Redis Streams is an implementation of the immutable log that uses radix trees as the underlying data structure. Each stream entry is identified by a timestamp and can contain an arbitrary set of field-value pairs. Entries of the same stream can have different fields, but Redis is able to compress multiple events in a row that share the same schema. This means that if your events have stable set of fields you won’t pay a storage price for each field name, letting you use longer and more descriptive key names without any downside.