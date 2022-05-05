# （WIP）repos-researcher
参考になるGitHubのアカウントやリポジトリを管理するWebアプリ

## Usage
### 【検索】
* アカウント
    * ソート, フィルタリング
        * 保有リポジトリ数, コミット数, Star数, 被スター数, Fork数, 利用言語
* リポジトリ
    * ソート, フィルタリング
        * 対象アカウント, コミット数, 言語, 修正期間, 作成日時
### 【保存】
*  アカウント
    * read later（後で読む）
    * like（保存）
*  リポジトリ
    * read later（後で読む）
    * like（保存）

## Architecture
* FE
    * Amplify
        * React
* BE
    * Lambda
        * Go
* DB
    * DynamoDB
* CI/CD
    * GitHub Actions
    * Amazon ECR
    * ※ mainブランチへのpushをトリガーにコンテナイメージをビルドしてECRにプッシュし、Lambdaにデプロイ