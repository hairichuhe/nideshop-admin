
<div align=center>
<font style="font-size:50px;"> nideshop-admin</font>
</div>
<div align=center>
<img src="https://img.shields.io/badge/vue-2.6.10-brightgreen"/>
<img src="https://img.shields.io/badge/element--ui-2.12.0-green"/>
<img src="https://img.shields.io/badge/golang-1.12-blue"/>
<img src="https://img.shields.io/badge/gin-1.4.0-lightBlue"/>
<img src="https://img.shields.io/badge/gorm-1.9.10-red"/>
</div>


# nideshop-admin gin+vue开源快速项目模板
本模板使用前端ui框架为 element-ui https://element.eleme.cn/#/zh-CN 前端组件可查看elementUi文档使用
## 写在前面
非常感谢 [@piexlmax](https://github.com/piexlmax)提供项目脚手架，脚手架地址 [gin-vue-admin](https://github.com/piexlmax/gin-vue-admin)
    

## 环境搭建教学视频

腾讯视频：https://v.qq.com/x/page/e3008xjxqtu.html    (等待最新视频录制)
    
## 模板使用教学及展示视频

腾讯视频：https://v.qq.com/x/page/c3008y2ukba.html    (等待最新视频录制)

## 技术选型
    1.后端采用golang框架gin，快速搭建基础restful风格API
    2.前端项目采用VUE框架，构建基础页面
    3.数据库采用Mysql(5.6.44)版本不同可能会导致SQL导入失败
    4.使用redis实现记录当前活跃用户的jwt令牌并实现多点登录限制
    5.使用swagger构建自动化文档
    6.使用fsnotify和viper实现json格式配置文件
    7.使用logrus实现日志记录
    8.使用gorm实现对数据库的基本操作

## 项目说明
    static/config存放mysql相关配置。可以根据自己的mysql数据库名 用户名 密码修改对应配置
    vue项目存放于QMPlusVuePage文件夹下
    开源不易，感谢各位支持，错误指出即刻改正，改写纠错，感谢star支持
## TODO
###脚手架
    1.基本用户注册登录 √
    2.用户等基础数据CURD √
    3.调用des实现数据加密 √
    4.实现基于jwt的权限管理 √
    6.封装了分页方法，实现分页接口并且复制粘贴就可使用分页 √
    7.前端分页mixin封装 分页方法调用mixins即可 √
    8.图片上传前端下载功能 √
    9.富文本编辑器，MarkDown编辑器功能嵌入 √
    10.增加条件搜索示例 前端文件参考src\view\superAdmin\api\api.vue 后台文件参考 model\dnModel\api.go √
    11.增加了多点登录限制 体验需要再 static\config中 把 system中的useMultipoint 修改为 true(需要自行配置redis和config中的redis参数)(测试阶段，有bug请及时反馈)√
    12.增加了配置文件管理功能 √
###nideshop-admin
    1.商品管理
    2.商品分类管理
## 使用说明
    1.golang api server 基于go.mod 如果golang版本低于1.11 请自行升级golang版本
    2.支持go.mod的golang版本在运行go list 和 编译之前都会自动下载所需要的依赖包
    3.前端项目node建议高于V8.6.0
    4.到前端项目目录下运行 npm i 安装所需依赖
    5.依赖安装完成直接运行 npm run serve即可启动项目
    6.如果要使用swagger自动化文档 首先需要安装 swagger
````
go get -u github.com/swaggo/swag/cmd/swag
````
由于国内没法安装到X包下面的东西 如果可以翻墙 上面的命令就可以让你安心使用swagger了
如果没有翻墙的办法那就先装一下 gopm

````
go get -v -u github.com/gpmgo/gopm
````
此时你就可以使用 gopm了
这时候执行
````
gopm get -g -v github.com/swaggo/swag/cmd/swag
````
等待安装完成以后
到我们GOPATH下面的/src/github.com/swaggo/swag/cmd/swag路径
执行
````
go install
````
安装完成过后在项目目录下运行
````
swag init
````
项目文件夹下面会有 doc文件夹出现
这时候登录 localhost:8888/swagger/index.html
就可以看到 swagger文档啦

## docker镜像
   感谢 [@chenlinzhong](https://github.com/chenlinzhong)提供docker镜像
   
      #启动容器
      docker run -itd --net=host --name=go_container shareclz/go_node /bin/bash;
      
      #进入容器
      docker exec -it go_container /bin/bash;
      git clone https://github.com/piexlmax/nideshop-admin.git /data1/www/htdocs/go/admin;
      
      #启动前端
      cd /data1/www/htdocs/go/admin/QMPlusVuePage;
      cnpm i ;
      npm run serve;
      
      #修改数据库配置
      vi /data1/www/htdocs/go/admin/QMPlusServer/static/dbconfig/config.json;
      
      #启动后端
      cd /data1/www/htdocs/go/admin/QMPlusServer;
      go run main.go;


## 联系方式

<div align=center>
<h3>qq交流群:234706951</h3>
<h3>微信交流群加微信：14320794 备注"加入nideshop-admin交流群"</h3>
</div>


## 更新日志
