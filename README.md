# jazzychai/extractsql

从庞大的数据库备份文件中提取单表的数据

使用方法
```bash
    # 从db.sql中提取user表的数据保存至result_user.sql
    ./extractsql -f db.sql -t user -r result_user.sql
```