|   项目名    |    功能    |
| :---------: | :--------: |
| staruml-cpp | 从UML到CPP |

插件启动流程
1. 从jars文件夹里面读取，可以动态加载
2. 加载到PluginJAR类里面
3. 查找.props后缀的文件
4. 读取action.xml里面的动作
5. 