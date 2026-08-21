create a system where user can post there blogs and can fetch 'em. The role of msg broker will be to update the count of the blogs per user and add perform a write on every blog, then save it.

So, here will have 3 tasks to be done, So now, since we have three things to do, that is:
- write
- save
- count.

There will be three separate consumers that will be working for each, respectively. That means we should be having three topics, and within each topic, three partitions and consumer.

We will not be creating a different topic for each because they all have the same event. It's an event domain, for example. It's not like the execution is different. It is on the same block that we are performing different execution, not on different blocks. One topic, many consumers.