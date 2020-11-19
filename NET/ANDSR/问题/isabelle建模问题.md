# isabelle问题

## 1

```isabelle
record SGRule =
    id :: str 
    remote_group_id :: str
    sg_direction :: Direction
    remote_ip_prefix :: "raw_ipv4addr option"
    proto :: primitive_protocol
    port_range_min :: "nat option"
    port_range_max :: "nat option"
    security_group_id :: str 
```

其中`remote_ip_prefix`为什么要使用`raw_ipv4addr ( "(32 word × 32 word) list")`

## 2

