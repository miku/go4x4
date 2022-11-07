# Notes

* nested goroutines

```
run([s1, s2, s3, ...])

for s in sources:                     # 5
   for resource in source:            # 20
        for processor in processors:
            if suitable:      
                process(resource)   
            write(resource)
```

* buffer, losing messages
* task queues
* settings.py like thing