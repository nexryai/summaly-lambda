## Misskey's Summaly Proxy on AWS Lambda

### これは何？
MisskeyのSummalyをAWS Lambdaで動かすためのリポジトリです。  
本家の公式実装ではなくGoで書かれたSummergoを使用するため軽量です。

### How to use

1. AWSのコンソールで適当な名前のLambda関数を作成します。ランタイムは"Amazon Linux 2"を選択してください。arm64環境がおすすめです。
2. ハンドラを`bootstrap`に設定します。
3. リポジトリをcloneして`GOOS=linux GOARCH=[amd64 or arm64（Lambdaのアーキテクチャに合わせて選択）] go build -o bootstrap main.go`を実行します。
4. `zip lambda-handler.zip bootstrap`でzipファイルを作成し、Lambda関数にアップロードします。
  * Webコンソールからアップロードする場合は、"コード" → "アップロード元" → ".zipファイル" でアップロードできます。