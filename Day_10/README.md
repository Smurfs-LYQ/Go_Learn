#### <center>Day_10</center>

1. 数据类型

#### <center>笔记</center>

1. > 为什么需要依赖管理
  
    - 最早的时候，Go所依赖的所有的第三方库都放在GOPATH这个目录下面。这就导致了同一个库只能保存一个版本的代码。如果不同的项目依赖同一个第三方库的不同版本，应该怎么解决？

    

2. > dodep

    - Go语言从v1.5开始引入 `vendor` 模式，如果项目目录下有vendor目录，那么go工具链就会优先使用 `vendor` 内的包进行编译、测试等。

    - `godep` 是一个通过vender模式实现的Go语言的第三方依赖管理工具，类似的还有由社区维护准官方包管理工具 `dep`。

    - **安装**

        ```go
        go get github.com/tools/godep
        ```

    - **基本命令**

        安装好godep之后，在终端输入 `godep` 查看支持的所有命令。

        ```go
        godep save     将依赖项输出并复制到Godeps.json文件中
        godep go       使用保存的依赖项运行go工具
        godep get      下载并安装具有指定依赖项的包
        godep path     打印依赖的GOPATH路径
        godep restore  在GOPATH中拉取依赖的版本
        godep update   更新选定的包或go版本
        godep diff     显示当前和以前保存的依赖项集之间的差异
        godep version  查看版本信息
        ```

        使用 `godep help [command]` 可以看看具体命令的帮助信息。

    - **使用godep**

        在项目目录下执行 `godep save` 命令，会在当前项目中创建 `Godeps` 和 `vender` 两个文件夹。
        其中 `Godeps` 文件夹下有一个 `Godeps.json` 的文件，里面记录了项目所依赖的包信息。`vendor` 文件夹下是项目依赖的包的源代码文件。

    - **vender机制**

        Go1.5版本之后开始支持，能够控制Go语言程序编译时依赖包搜索路径的优先级。

        例如查找项目的某个依赖包，首先会在项目根目录下的 `vender` 文件夹中查找，如果没有找到就会去 `$GOPATH/src` 目录下查找。

    - **godep开发流程**

      - 保证程序能够正常编译
      - 执行 `godep save` 保存当前项目的所有第三方依赖的版本信息和代码
      - 提交Godeps目录和vender目录到代码库
      - 如果要更新依赖的版本，可以直接修改 `Godeps.json` 文件中的对应项

3. > go module

    - `go module` 是Go1.11版本之后官方推出的版本管理工具

4. >