# sgn-layout
Convenient and quick service startup, unified project template, three-key generation of microservices

# 1. Create a project-specific template
Assume the project directory is projectName

```
kratos new $(projectName) -r git@github.com:itsgn/sgn-layout.git
```
# 2. Initialize the project package
The value of projectName here needs to be consistent with the first step
```
make initProject PROJECT=$(projectName)
```
# 3.Initialize project services
```
make initNewService ServiceUpperName=Store ServiceLowerName=store
```

# Summary
The above operation will generate a directory named projectName, containing a crud service named StoreService.
* db: postgresql, orm: gorm
* cache:redis
* trace:jaeger
* registry && discovery: etcd
* log: kratos/v2/log json format
* validate: protoc-gen-validate


pkg/utils: Commonly used function packages
# Other notes

0. If the environment is set to dev, it is a development environment. You can set rules based on this variable, such as debug mode, print sql

1. Integer values ​​are all int64. When converting proto to json, the int64 type will be converted to string type due to precision issues

2. The interface document is in openapi.yaml
3. The local development and docker deployment files are configs/config-dev.yaml, which can be changed in Dockerfile



