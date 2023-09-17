# セットアップスクリプトの実行
# DB自身とDB内のスキーマの作成
mysql -u admin -padmin mdet < "/docker-entrypoint-initdb.d/001-create-table