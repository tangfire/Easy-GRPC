# 服务端
1. 编写一个结构体，名字叫什么不重要
2. 重要的是得实现protobuf中的所有方法
3. 监听端口
4. 注册服务


# 客户端

1. 建立连接
2. 调用方法


# 编写风格

1. 文件名建议下划线，例如：my_student.proto
2. 包名和目录名对应
3. 服务名、方法名、消息名均为大驼峰
4. 字段名为下划线


# 多个proto文件

当项目大起来之后，会有很多个service，rpc，message

我们会将不同服务放在不同的proto文件中

还可以放一些公共的proto文件

对于这方面的资料，可以说全网的正确资料真的相当少

其实本质就是生成go文件，需要在一个包内