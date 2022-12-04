基于schema的定义（字段，索引，边，配置）来生成对应的数据表操作代码
ent generate --target <target dirpath> <template dirpath>
template dirpath目的是指定模板所在路径，根据所在目录下的模板来生成对应的代码
--target <target dirpath?目的是指定生成操作代码的路径
使用ent generate --target gen/entschema spec/schema将我们的模板生成实际的操作代码到gen/entschema路径
