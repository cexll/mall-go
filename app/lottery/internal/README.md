# business structure

The directory structure within the project can take multiple forms, such as DDD or generated structures.


## DDD
The structure of DDD also varies in different versions, both domestic and foreign.

`standard version`
```bash
domain: used to store domain layer code and models.
application: used to store application layer code and services.
interfaces: used to store external interface and adapter code.
```

`kratos`
```
biz: used to store code and models related to business logic.
data: used to store code and models related to data access.
service: used to store code and models related to service implementation.
```


## Generated
```
gen: used to store generated files such as gorm-gen and swagger-gen.
model: used to store database models.
logic: used to write business logic.
service: used to store code and models related to service implementation.
```
