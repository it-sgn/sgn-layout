# sgn-layout
Convenient and quick service startup, unified project template, and three-click generation of microservices

# 1.Create project-specific templates
Assume that the project directory is projectName

```
kratos new $(projectName) -r github.com/it-sgn/sgn-layout.git
```
# 2. Initialize project package
The value of projectName here needs to be consistent with the first step
```
make initProject PROJECT=$(projectName)
```
# 3.Initialize project services
```
make initNewService ServiceUpperName=Store ServiceLowerName=store
```

# Summarize
The above operation will generate a directory called projectName, containing a crud service with a service name of StoreService.

* db: mysql, orm: gorm
* cache:redis
* trace:jaeger
* registry && discovery: etcd
* log: kratos/v2/log json格式
* validate: protoc-gen-validate


pkg/utils: Commonly used function packages

# Other things to note

0. If the environment is set to dev, it is a development environment. You can set rules according to this variable, such as debug mode and print sql.

1. All integer numerical types use int64. When converting proto to json, due to accuracy issues, the int64 type will be converted to string type.

2. The interface document is in openapi.yaml
3. The local development and docker deployment file is configs/config-dev.yaml, which can be changed in the Dockerfile.



