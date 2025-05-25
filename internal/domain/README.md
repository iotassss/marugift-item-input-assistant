# domain

- エラーはerrors.goに記載する。
- repositoryの実装等で発生した特定のエラーをusecaseに伝える必要がある場合はerrors.goにエラーを定義し、usecaseでerrors.Is()で判定する。
- loaderは、マスタデータのseed用途など「一方向の読み取り」だけを行う処理に使用する。
